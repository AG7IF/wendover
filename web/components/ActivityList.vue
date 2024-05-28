<script setup lang="ts">

const props = defineProps(['activities'])

const selectedActivity = ref('new')

function onSubmit() {
  navigateTo(`/activities/${selectedActivity.value}`)
}

</script>

<template>
  <v-form @submit.prevent="onSubmit">
    <v-radio-group v-model="selectedActivity">
      <v-radio v-for="activity in activities" :label="activity.name" :value="activity.key"/>
      <v-radio label="Create New" value="new"/>
    </v-radio-group>
    <v-btn type="submit">Go</v-btn>
  </v-form>
</template>

<style scoped>

</style>