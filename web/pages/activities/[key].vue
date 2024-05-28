<script setup lang="ts">
const route = useRoute()

const { data: activity, error } = await useFetch(`http://localhost:8080/api/v1/activity/${route.params.key}`)

/*
const startDateStr = computed(() => {
  const month = activity.start.toLocaleString('en-US', {month:'short'})

  return `${activity.start.getDate()} ${month} ${activity.start.getFullYear()}`
})

const endDateStr = computed(() => {
  const month = activity.end.toLocaleString('en-US', {month:'short'})

  return `${activity.end.getDate()} ${month} ${activity.end.getFullYear()}`
})
 */

</script>

<template>
  <div>
    <v-alert v-if="error != null" title="Activity Not Found" type="error">
      Either '{{route.params.key}}' is not a valid activity key, or you do not have permission to
      view this activity.
    </v-alert>

    <NuxtLink v-if="error != null" to="/activities">Go back to activities list</NuxtLink>

    <v-card v-if="error == null">
      <v-card-title>{{activity.name}}, {{activity.location}}<!--: {{startDateStr}} to {{endDateStr}}--></v-card-title>
    </v-card>
  </div>
</template>

<style scoped>

</style>