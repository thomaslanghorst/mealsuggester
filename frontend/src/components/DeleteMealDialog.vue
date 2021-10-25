<template>
  <div class="text-center">
    <v-dialog
      v-model="dialog"
      width="500"
    >
      <template v-slot:activator="{ on, attrs }">
        <v-btn
          text
          small
          v-bind="attrs"
          v-on="on"
          color="red lighten-2"
        >
          <v-icon small dark>mdi-delete</v-icon>
        </v-btn>
      </template>

      <v-card>
        <v-card-title class="headline grey lighten-2">
          Delete {{this.mealToDelete.name}} ?
        </v-card-title>

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
            color="red"
            text
            @click="ok"
          >
            OK
          </v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>
  </div>
</template>

<style scoped>
</style>

<script>
export default {
    name: 'DeleteMealDialog',
     props: {
        mealToDelete: Object,
    },

    data() {
        return {
          dialog: false,
        }
    },

    methods: {
      cancel() {
        this.dialog = false;
      },

      ok() {
        this.dialog = false;
        this.$emit("onDeleteMeal", this.mealToDelete.id);
      },
    },
}
</script>