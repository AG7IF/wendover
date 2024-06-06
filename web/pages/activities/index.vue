<script setup lang="ts">
  const config = useRuntimeConfig()
  const { data: activityList } = await useFetch(`${config.public.apiBase}/activity`)

</script>

<template>
<div>
  <ActivityList :activities="activityList"/>
</div>
</template>

<style scoped>

</style>