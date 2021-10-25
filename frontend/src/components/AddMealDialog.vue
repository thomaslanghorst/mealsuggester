<template>
  <div class="text-center">
    <v-dialog
      v-model="dialog"
      width="500"
    >
      <template v-slot:activator="{ on, attrs }">
        <v-btn
          color="blue lighten-2"
          dark
          v-bind="attrs"
          v-on="on"
        >
          Add
        </v-btn>
      </template>

      <v-card>
        <v-card-title class="headline grey lighten-2">
          Add new meal
        </v-card-title>

        <v-card-text class="meal-info">
            <v-select dense :items="categories" label="category" v-model="selectedCategory" />
          <v-text-field label="name" v-model="mealname" placeholder="Name..."></v-text-field>
        </v-card-text>

        <v-divider></v-divider>

        <v-card-actions>
          <v-spacer></v-spacer>
            <v-btn
            text
            @click="cancel"
          >
            Cancel
          </v-btn>
          <v-btn
            color="primary"
            text
            @click="save"
          >
            Save
          </v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>
  </div>
</template>

<style scoped>
.meal-info {
    margin-top: 20px;
}
</style>

<script>
export default {
    name: 'AddMealDialog',

    data() {
        return {
            mealname: '',
            dialog: false,
            selectedCategory: '',
            categories: ['fast food', 'low carb'],
        }
    },

    methods: {
        cancel() {
          this.mealname = '';
          this.selectedCategory = '';
          this.dialog = false;
        },

        save() {
          const data = {"category":this.selectedCategory, "name":this.mealname};
          this.selectedCategory = '';
          this.mealname = '';
          this.dialog = false;
          this.$emit("onAddMeal", data);
        },
    }
}
</script>