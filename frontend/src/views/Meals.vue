<template>
    <div>
        <h1 class="title">{{title}}</h1>
    
        <div class="category-select">
            <v-select dense :items="categories" label="category" v-model="selectedCategory" />
            <AddMealDialog class="add-button" v-on:onAddMeal=onAddMeal />
        </div>
        <ul>
            <li v-for="meal in filteredMeals" :key="meal.id">
                <div class="meal"
                    v-bind:class="{'selected': selectedMeal === meal}" 
                    v-on:click="setSelectedMeal(meal)"
                >
                    <span>
                        {{meal.name}} 
                    </span>
                    <div 
                        class="meal-buttons"
                        v-bind:class="{'meal-buttons-visible': selectedMeal === meal}"
                    >
                        <EditMealDialog
                            v-bind:mealToEdit=meal
                            v-on:onEditMeal=onEditMeal />
                        <DeleteMealDialog
                            v-bind:mealToDelete=meal
                            v-on:onDeleteMeal=onDeleteMeal />
                    </div>
                </div>
            </li>
        </ul>
    </div>
</template>

<style scoped>
.title {
    width: 100%;
    text-align: center;
    margin-bottom: 10px;
}

.add-button {
    margin-left: 20px;
}

.meal {
    display: flex;
    justify-content: space-between;
}

.meal-buttons {
    min-width: 100px;
    display: flex;
    visibility: hidden;
}

.meal-buttons-visible {
    min-width: 100px;
    display: flex;
    visibility: visible;
}

.category-select {
    display: flex;
    margin: 20px;
}

.selected{
    color: #65aaf4;
}

ul {
    list-style-type: none;
}
</style>

<script>
import AddMealDialog from '@/components/AddMealDialog.vue'
import EditMealDialog from '@/components/EditMealDialog.vue'
import DeleteMealDialog from '@/components/DeleteMealDialog.vue'

export default {
    components: {
        AddMealDialog,
        EditMealDialog,
        DeleteMealDialog,
    },

    data() {
        return {
            title: "Meals",
            selectedCategory: '',
            categories: ['all', 'fast food', 'low carb'],
            selectedMeal: null,
            meals: [],

            endpoint: '/meals'
        }
    },

    created() {
        this.listMeals();
    },

    methods: {
        listMeals() {
            let url = this.endpoint;

            if (this.selectedCategory != 'all') {
                url = url+'?category='+this.selectedCategory
            }
            this.$httpclient
                .get(url)
                .then(res => {
                    if (res.status === 200 && res.data) {
                        this.meals = res.data
                        this.sort()
                    } else {
                        console.log(res);
                    }
                })
                .catch(err => {
                    console.log("---ERROR---");
                    console.log(err); 
                });
        },
        sort() {
            if (this.meals) {
                this.meals = this.meals.slice(0).sort((a, b) => a.name > b.name ? 1 : -1);
            }
        },
        onAddMeal(data) {
            let url = this.endpoint;

            const category = data['category'];

            if (category !== '' || category !== 'all') {
                url = url+'?category='+category
            }

            this.$httpclient
                .post(url, data)
                .then(res => {
                    if (res.status === 201 && res.data) {
                        this.listMeals();
                    } else {
                        console.log(res);
                    }
                })
                .catch(err => {
                    console.log("---ERROR---");
                    console.log(err); 
                });
        },
        onEditMeal(data) {
            const url = this.endpoint;

            this.$httpclient
                .put(url, data)
                .then(res => {
                    if (res.status === 200 && res.data) {
                        this.listMeals();
                    } else {
                        console.log(res);
                    }
                })
                .catch(err => {
                    console.log("---ERROR---");
                    console.log(err); 
                });
        },
        onDeleteMeal(id) {
            const url = this.endpoint+'/'+id;

            this.$httpclient
                .delete(url)
                .then(res => {
                    if (res.status === 204) {
                        this.listMeals();
                    } else {
                        console.log(res);
                    }
                })
                .catch(err => {
                    console.log("---ERROR---");
                    console.log(err); 
                });
        },
        setSelectedMeal(meal) {
            this.selectedMeal = meal;
        }
    },

    computed: {
        filteredMeals() {
            if (this.selectedCategory === '' || this.selectedCategory === 'all') {
                return this.meals;
            } else {
                return this.meals.filter(meal => meal.category === this.selectedCategory);
            }
        }
    }
}
</script>