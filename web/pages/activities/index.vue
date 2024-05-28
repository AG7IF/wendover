<script setup lang="ts">
  const { data: activityList } = await useFetch('http://localhost:8080/api/v1/activity')


</script>

<template>
<div>
  <ActivityList :activities="activityList"/>
</div>
</template>

<style scoped>

</style>