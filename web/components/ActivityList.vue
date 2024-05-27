<script setup lang="ts">

const props = defineProps(['activities'])

function submit() {

}

</script>

<template>
  <v-radio-group>
    <v-radio v-for="activity in activities" label="{{activity.name}}" value="{{activity.key}}"/>
  </v-radio-group>
</template>

<style scoped>

</style>