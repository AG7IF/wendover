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
  <div>
    <h1>New Activity</h1>

    <table>
      <tr>
        <td><label for="activityKey" class="ml-2">Activity Key</label></td>
        <td><InputText id="activityKey" type="text" v-model="activityKey"/></td>
      </tr>

      <tr>
        <td><label for="activityName" class="ml-2">Activity Name</label></td>
        <td><InputText id="activityName" type="text" v-model="activityName"/></td>
      </tr>

      <tr>
        <td><label for="activityLocation" class="ml-2">Location</label></td>
        <td><InputText id="activityLocation" type="text" v-model="activityLocation"/></td>
      </tr>

      <tr>
        <td><label for="activityStart" class="ml-2">Start Date</label></td>
        <td><Calendar id="activityStart" v-model="activityStart" dateFormat="dd-M-yy" showIcon/></td>
      </tr>

      <tr>
        <td><label for="activityEnd" class="ml-2">End Date</label></td>
        <td><Calendar id="activityEnd" v-model="activityEnd" dateFormat="dd-M-yy" showIcon/></td>
      </tr>

      <tr>
        <td><label for="cadetStudentFee" class="ml-2">Cadet Student Fee</label></td>
        <td><InputText id="cadetStudentFee" type="number" v-model="cadetStudentFee"/></td>
      </tr>

      <tr>
        <td><label for="cadetCadreFee" class="ml-2">Cadet Cadre Fee</label></td>
        <td><InputText id="cadetCadreFee" type="number" v-model="cadetCadreFee"/></td>
      </tr>

      <tr>
        <td><label for="seniorStudentFee" class="ml-2">Senior Student Fee</label></td>
        <td><InputText id="seniorStudentFee" type="number" v-model="seniorStudentFee"/></td>
      </tr>

      <tr>
        <td><label for="seniorCadreFee" class="ml-2">Senior Cadre Fee</label></td>
        <td><InputText id="seniorCadreFee" type="number" v-model="seniorCadreFee"/></td>
      </tr>

      <tr>
        <td><Button label="Create" @click="onSubmit"/></td>
      </tr>
    </table>

  </div>

</template>

<style scoped>

</style>