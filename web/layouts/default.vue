<script setup>
import { ref } from 'vue'

const drawer = ref(false)
</script>

<template>
  <v-app id="inspire">
    <v-navigation-drawer v-model="drawer">
      <!--  -->
    </v-navigation-drawer>

    <v-app-bar>
      <v-app-bar-nav-icon @click="drawer = !drawer"></v-app-bar-nav-icon>

      <v-app-bar-title>CAPForms.net</v-app-bar-title>
    </v-app-bar>

    <v-main>
      <slot/>
    </v-main>
  </v-app>
</template>
