<template>
  <v-app>
    <v-main >
      <div>
      <router-view/>
      <div class="buttons-wrapper">
        <v-btn
          class="button"
          color="primary lighten-2"
          tile
          @click="goToMeals">
          Meals
        </v-btn>
        <v-btn
          class="button"
          color="primary lighten-2"
          tile
          @click="goToSuggestions">
          Suggestions
        </v-btn>
      </div>
      </div>
     
    </v-main>
    
  </v-app>
</template>

<style scoped>
.buttons-wrapper {
  position: absolute;
  bottom: 0px;
  width: 100%;
}

.button {
  width: 50%;
}
</style>

<script>

export default {
  name: 'App',

  methods: {
    goToMeals() {
      this.$router.push('/');
    },
    goToSuggestions() {
      this.$router.push('/suggestions');
    },
  }
};
</script>
