package meals

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

type MealHandler struct {
	service MealsServiceInterface
}

func NewMealHandler(s MealsServiceInterface) *MealHandler {
	return &MealHandler{
		service: s,
	}
}

func (h *MealHandler) ListMeals(w http.ResponseWriter, r *http.Request) {
	categoryParam := r.URL.Query().Get("category")

	meals, err := h.service.ListMeals(categoryParam)
	if err != nil {
		log.Warnf(err.Error())
		sendError(w, http.StatusInternalServerError, err.Error())
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(200)
	err = json.NewEncoder(w).Encode(meals)
	if err != nil {
		log.Warnf("error encoding meals. Error: %s", err.Error())
		sendError(w, http.StatusInternalServerError, "error encoding meals")
		return
	}
}

func (h *MealHandler) CreateMeal(w http.ResponseWriter, r *http.Request) {
	var m Meal

	err := json.NewDecoder(r.Body).Decode(&m)
	defer r.Body.Close()

	if err != nil {
		log.Warnf(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("error while decoding meal"))
		return
	}

	id, err := h.service.CreateMeal(m)
	if err != nil {
		log.Warnf(err.Error())
		sendError(w, http.StatusInternalServerError, err.Error())
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Add("Content-Type", "application/json")
	w.Write([]byte(fmt.Sprintf("{\"id\":\"%d\"}", id)))
	return
}

func (h *MealHandler) DeleteMeal(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	idParam := vars["id"]

	id, err := strconv.Atoi(idParam)
	if err != nil {
		log.Warnf(err.Error())
		sendError(w, http.StatusInternalServerError, "no id parameter provided: meals/{id}")
		return
	}

	err = h.service.DeleteMeal(id)
	if err != nil {
		log.Warnf(err.Error())
		sendError(w, http.StatusInternalServerError, err.Error())
		return
	}

	w.WriteHeader(http.StatusNoContent)
	return
}

func (h *MealHandler) EditMeal(w http.ResponseWriter, r *http.Request) {
	var m Meal

	err := json.NewDecoder(r.Body).Decode(&m)
	defer r.Body.Close()

	if err != nil {
		log.Warnf(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("error while decoding meal"))
		return
	}

	err = h.service.EditMeal(m.ID, m)
	if err != nil {
		log.Warnf(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(200)
	err = json.NewEncoder(w).Encode(m)
	if err != nil {
		log.Warnf("error encoding meal. Error: %s", err.Error())
		sendError(w, http.StatusInternalServerError, "error encoding meal")
		return
	}
}

func (h *MealHandler) SuggestMeals(w http.ResponseWriter, r *http.Request) {
	// default
	count := 5

	countParam := r.URL.Query().Get("count")
	categoryParam := r.URL.Query().Get("category")

	if countParam != "" {
		c, err := strconv.Atoi(countParam)
		if err != nil {
			log.Warnf(err.Error())
			sendError(w, http.StatusInternalServerError, "error parsing request parameter count")
			return
		}

		count = c
	}

	meals, err := h.service.SuggestMeals(count, categoryParam)
	if err != nil {
		log.Warnf("error suggesting meals. Error: %s", err.Error())
		sendError(w, http.StatusInternalServerError, err.Error())
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(200)
	err = json.NewEncoder(w).Encode(meals)
	if err != nil {
		log.Warnf("error encoding meals. Error: %s", err.Error())
		sendError(w, http.StatusInternalServerError, "error encoding meals")
		return
	}
}

func sendError(w http.ResponseWriter, statusCode int, msg string) {

	body := []byte(fmt.Sprintf("{\"error_message\":\"%s\"}", msg))

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(body)
}
