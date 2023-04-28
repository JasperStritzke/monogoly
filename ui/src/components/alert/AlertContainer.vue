<template>
  <div class="flex flex-col gap-4 fixed bottom-5 w-96 left-1/2 -translate-x-1/2 z-20 pointer-events-none">
    <transition-group name="list">
      <Alert v-for="alert in alertsStore.alerts" :key="alert.id" :alert="alert"/>
    </transition-group>
  </div>
</template>

<script setup lang="ts">
import Alert from "./Alert.vue";
import useAlertStore from "../../store/alert.store";

const alertsStore = useAlertStore()
</script>
<style>
.list-move, /* apply transition to moving elements */
.list-enter-active,
.list-leave-active {
  transition: all 0.2s ease-out;
}

.list-enter-from,
.list-leave-to {
  opacity: 0;
  transform: translateY(30px) scale(0.8);
}

/* ensure leaving items are taken out of layout flow so that moving
   animations can be calculated correctly. */
.list-leave-active {
  position: absolute;
  z-index: -1;
}

</style>