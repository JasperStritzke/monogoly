<script setup lang="ts">
import AlertContainer from "./components/alert/AlertContainer.vue";
</script>

<template>
  <p class="text-sm opacity-20 select-none absolute bottom-4 left-0 right-0 text-center">made with &heartsuit; by <a class="underline" target="_blank" href="https://github.com/JasperStritzke">Jasper Stritzke</a></p>

  <router-view/>

  <AlertContainer/>
</template>