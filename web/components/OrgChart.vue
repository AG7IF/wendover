<script setup lang="ts">
const ds = reactive({
  'id': 1,
  'name': 'Group',
  'children': [
    {
      'id': 2,
      'name': 'Squadron 1',
      'children': [
        {
          'id': 3,
          'name': 'Alpha Flight',
        },
        {
          'id': 4,
          'name': 'Bravo Flight',
        },
        {
          'id': 5,
          'name': 'Charlie Flight',
        },
      ]
    },
    {
      'id': 3,
      'name': 'Squadron 2',
      'children': [
        {
          'id': 6,
          'name': 'Delta Flight',
        },
        {
          'id': 7,
          'name': 'Echo Flight',
        },
        {
          'id': 8,
          'name': 'Foxtrot Flight',
        },
      ]
    },
  ]
})

</script>

<template>
<OrganizationChart :value="ds">
  <template #default="slotProps">
    <span>{{ slotProps.node.name }}</span>
  </template>
</OrganizationChart>
</template>

<style scoped>

</style>