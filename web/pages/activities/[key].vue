<script setup lang="ts">
const config = useRuntimeConfig()
const route = useRoute()

const { data, error } = await useFetch(`${config.public.apiBase}/activity/${route.params.key}`)


</script>

<template>
  <div>
    <Message v-if="error != null" severity="error">
      Either '{{route.params.key}}' is not a valid activity key, or you do not have permission to
      view this activity.
    </Message>

    <NuxtLink v-if="error != null" to="/activities">Go back to activities list</NuxtLink>

    <div v-if="error == null">
      <h1>{{data.name}}, {{data.location}}</h1>
      <OrgChart/>
    </div>
  </div>
</template>

<style scoped>

</style>