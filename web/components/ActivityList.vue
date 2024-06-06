<script setup lang="ts">

const props = defineProps(['activities'])

const selectedActivity = ref('new')

function onGo() {
  navigateTo(`/activities/${selectedActivity.value}`)
}

</script>

<template>
  <div class="flex flex-col gap-3">
    <div v-for="activity in activities" class="flex align-items-center">
      <RadioButton v-model="selectedActivity" :inputId="activity.key" :value="activity.key"/>
      <label :for="activity.key" class="ml-2">{{ activity.name }}</label>
    </div>
    <div class="flex align-items-center">
      <RadioButton v-model="selectedActivity" inputId="new" value="new"/>
      <label for="new" class="ml-2">New Activity</label>
    </div>
    <Button label="Go" @click="onGo"/>
  </div>
</template>

<style scoped>

</style>