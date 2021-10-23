package meals

import (
	"database/sql"
	"errors"
	"fmt"
)

type MealsServiceInterface interface {
	ListMeals(category string) ([]Meal, error)
	CreateMeal(m Meal) (int, error)
	DeleteMeal(id int) error
	EditMeal(id int, m Meal) error
	SuggestMeals(mealCount int, category string) ([]Meal, error)
}

type SqliteMealsService struct {
	db *sql.DB
}

func NewSqliteMealsService(db *sql.DB) *SqliteMealsService {
	return &SqliteMealsService{
		db: db,
	}
}

func (s *SqliteMealsService) ListMeals(category string) ([]Meal, error) {
	meals := make([]Meal, 0)

	var rows *sql.Rows
	var err error

	if category != "" {
		rows, err = s.db.Query("SELECT * FROM meals WHERE category = ?", category)
	} else {
		rows, err = s.db.Query("SELECT * FROM meals")
	}

	if err != nil {
		return meals, errors.New(fmt.Sprintf("unable to prepare SELECT meal statement. Error %s", err.Error()))
	}

	var mealId int
	var mealName string
	var mealCategory string

	for rows.Next() {
		err = rows.Scan(&mealId, &mealName, &mealCategory)
		if err != nil {
			return meals, errors.New(fmt.Sprintf("unable to query meals table. Error %s", err.Error()))
		}

		meals = append(meals, Meal{
			ID:       mealId,
			Name:     mealName,
			Category: mealCategory,
		})
	}

	return meals, nil
}

func (s *SqliteMealsService) CreateMeal(m Meal) (int, error) {

	stmt, err := s.db.Prepare("INSERT INTO meals(name, category) values(?,?)")
	if err != nil {
		return 0, errors.New(fmt.Sprintf("unable to prepare INSERT meal statement. Error %s", err.Error()))
	}

	res, err := stmt.Exec(m.Name, m.Category)
	if err != nil {
		return 0, errors.New(fmt.Sprintf("unable to execute create meals table statement. Error %s", err.Error()))
	}

	insertedId, err := res.LastInsertId()
	if err != nil {
		return 0, errors.New(fmt.Sprintf("unable to retrieve last inserted id. Error %s", err.Error()))
	}

	return int(insertedId), nil
}

func (s *SqliteMealsService) DeleteMeal(id int) error {

	stmt, err := s.db.Prepare("DELETE FROM meals WHERE id = ?")
	if err != nil {
		return errors.New(fmt.Sprintf("unable to prepare DELETE meal statement. Error %s", err.Error()))
	}

	res, err := stmt.Exec(id)
	if err != nil {
		return errors.New(fmt.Sprintf("unable to execute delete meals table statement. Error %s", err.Error()))
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return errors.New(fmt.Sprintf("unable to retrieve affected rows. Error %s", err.Error()))
	}

	if rowsAffected == 0 {
		return errors.New(fmt.Sprintf("no meal deleted"))
	}

	return nil
}

func (s *SqliteMealsService) EditMeal(id int, m Meal) error {

	stmt, err := s.db.Prepare("UPDATE meals set name=?, category=? WHERE id=?")
	if err != nil {
		return errors.New(fmt.Sprintf("unable to prepare UPDATE meal statement. Error %s", err.Error()))
	}

	res, err := stmt.Exec(m.Name, m.Category, id)
	if err != nil {
		return errors.New(fmt.Sprintf("unable to execute update meals table statement. Error %s", err.Error()))
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return errors.New(fmt.Sprintf("unable to retrieve affected rows. Error %s", err.Error()))
	}

	if rowsAffected == 0 {
		return errors.New(fmt.Sprintf("no meal updated"))
	}

	return nil
}

func (s *SqliteMealsService) SuggestMeals(mealCount int, category string) ([]Meal, error) {
	meals := make([]Meal, 0)

	var rows *sql.Rows
	var err error

	if category != "" {
		rows, err = s.db.Query("SELECT * FROM meals WHERE category = ? ORDER BY random() LIMIT ?", category, mealCount)
	} else {
		rows, err = s.db.Query("SELECT * FROM meals ORDER BY random() LIMIT ?", mealCount)
	}

	if err != nil {
		return meals, errors.New(fmt.Sprintf("unable to prepare SELECT meal statement. Error %s", err.Error()))
	}

	var mealId int
	var mealName string
	var mealCategory string

	for rows.Next() {
		err = rows.Scan(&mealId, &mealName, &mealCategory)
		if err != nil {
			return meals, errors.New(fmt.Sprintf("unable to query meals table. Error %s", err.Error()))
		}

		meals = append(meals, Meal{
			ID:       mealId,
			Name:     mealName,
			Category: mealCategory,
		})
	}

	return meals, nil
}
