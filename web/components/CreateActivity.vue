<script setup lang="ts">
  const activityKey = ref('')
  const activityName = ref('')
  const activityLocation = ref('')
  const activityStart = ref(new Date())
  const activityEnd = ref(new Date())
  const cadetStudentFee = ref(0)
  const cadetCadreFee = ref(0)
  const seniorStudentFee = ref(0)
  const seniorCadreFee = ref(0)

  async function onSubmit() {
    const response = await $fetch('http://localhost:8080/api/v1/activity', {
      method: 'POST',
      body: {
        key: activityKey.value,
        name: activityName.value,
        location: activityLocation.value,
        start: activityStart.value,
        end: activityEnd.value,
        cadet_student_fee: Math.trunc(parseFloat(cadetStudentFee.value) * 100),
        cadet_cadre_fee: Math.trunc(parseFloat(cadetCadreFee.value) * 100),
        senior_student_fee: Math.trunc(parseFloat(seniorStudentFee.value) * 100),
        senior_cadre_fee: Math.trunc(parseFloat(seniorCadreFee.value) * 100),
      }
    })

    navigateTo(`/activities/${activityKey.value}`)
  }
</script>

<template>
<v-card title="New Activity">
  <v-form @submit.prevent="onSubmit">
    <v-text-field label="Key" v-model="activityKey"/>
    <v-text-field label="Name" v-model="activityName"/>
    <v-text-field label="Location" v-model="activityLocation"/>
    <v-row>
      <v-col>
        <v-label>Start Date</v-label>
        <v-date-picker v-model="activityStart"/>
      </v-col>
      <v-col>
        <v-label>End Date</v-label>
        <v-date-picker v-model="activityEnd"/>
      </v-col>
    </v-row>
    <v-text-field label="Cadet Student Fee" type="number" v-model="cadetStudentFee"/>
    <v-text-field label="Cadet Cadre Fee" type="number" v-model="cadetCadreFee"/>
    <v-text-field label="Senior Student Fee" type="number" v-model="seniorStudentFee"/>
    <v-text-field label="Senior Cadre Fee" type="number" v-model="seniorCadreFee"/>
    <v-btn type="submit">Create</v-btn>
  </v-form>
</v-card>

</template>

<style scoped>

</style>