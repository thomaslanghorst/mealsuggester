<template>
    <div>
        <h1 class="title">{{title}}</h1>
    
        <div class="selects-wrapper">
            <v-select class="selector-left" dense :items="categories" label="category" v-model="selectedCategory" />
            <v-select class="selector-right" dense :items="mealsCount" label="number of meals" v-model="selectedCount" />
        </div>
         <div class="button-wrapper">
            <v-btn 
                class="suggest-button"
                color="primary"
                @click="suggestMeals">
                Surprise me!
            </v-btn>
        </div>
        <ul>
            <li v-for="meal in meals" :key="meal.id">
                  <span>{{meal.name}}</span>
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

.selects-wrapper {
    display: flex;
}
.selector-left {
    margin-left: 20px;
    margin-right: 10px;
}

.selector-right {
    margin-right: 20px;
    margin-left: 10px;
}

.button-wrapper {
    margin: 20px;
}

.suggest-button {
    width: 100%;
}

ul {
    list-style-type: none;
}
</style>

<script>
export default {
    data() {
        return {
            title: "Meal suggestions",
            selectedCategory: '',
            categories: ['all', 'fast food', 'low carb'],
            mealsCount: [1, 2, 3, 4, 5, 6, 7],
            selectedCount: 0,
            meals: [],

            endpoint: '/suggestions'
        }
    },

    created() {
        this.selectedCategory = 'all';
        this.selectedCount = 5;
    },

    methods: {
        sort() {
            if (this.meals) {
                this.meals = this.meals.slice(0).sort((a, b) => a.name > b.name ? 1 : -1);
            }
        },
        suggestMeals() {
            let url = this.endpoint+'?count='+this.selectedCount;

            if (this.selectedCategory !== 'all') {
                url = url+'&category='+this.selectedCategory
            }
            console.log(url)
            this.$httpclient
                .get(url)
                .then(res => {
                    if (res.status === 200 && res.data) {
                        this.meals = res.data
                        this.sort()
                    } else {
                        this.meals = []
                    }
                })
                .catch(err => {
                    console.log("---ERROR---");
                    console.log(err); 
                });
        }
    },
}
</script>