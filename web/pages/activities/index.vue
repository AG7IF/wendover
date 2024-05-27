<script setup lang="ts">
  const { data: activityList } = await useFetch('http://localhost:8080/api/v1/activity')
</script>

<template>
<div>
  <ActivityList v-if="activityList.length > 0" :activities="activityList"/>
  <CreateActivity v-if="activityList.length == 0"/>
</div>
</template>

<style scoped>

</style>