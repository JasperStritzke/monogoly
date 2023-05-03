<template>
  <div
      class="pointer-events-auto flex w-96 select-none flex-row items-center gap-5 px-5 py-5 shadow-lg alert"
      v-bind:data-type="alert.type"
      v-on:mouseenter="stopTimer"
      v-on:mouseleave="startTimer"
  >
    <i class="border-r border-white pr-4 text-lg" v-bind:class="{[`ri-${icon}-line`]: true}"></i>
    <p class="flex-grow text-xs">{{ alert.message }}</p>
    <i class="cursor-pointer text-gray-500 ri-close-line" @click="close"></i>
  </div>
</template>

<script setup lang="ts">
const ALERT_CLOSE_TIME_MS = 4000

import {computed, onMounted, ref} from "vue";
import useAlertStore from "../../store/alert.store";

type Props = {
  alert: {
    id?: number
    type: "success" | "warning" | "error" | "info"
    message: string,
  }
}

const props = defineProps<Props>()
const closeTimerRef = ref()

function startTimer() {
  closeTimerRef.value = setTimeout(close, ALERT_CLOSE_TIME_MS)
}

function stopTimer() {
  clearInterval(closeTimerRef.value)
}

onMounted(startTimer)

function close() {
  useAlertStore().removeAlert(props.alert)
}

const icon = computed(() => {
  switch (props.alert.type) {
    default:
    case "success":
      return "checkbox-circle"
    case "warning":
      return "error-warning"
    case "error":
      return "alarm-warning"
    case "info":
      return "information"
  }
})
</script>

<style lang="scss" scoped>
div.alert {
  border-radius: var(--baseBorderRadius);

  p {
    overflow: hidden;
    white-space: nowrap;
    text-overflow: ellipsis;
  }

  &[data-type='success'] {
    @apply bg-green-200;

    i:not(.ri-close-line) {
      @apply text-green-600
    }
  }

  &[data-type='warning'] {
    @apply bg-orange-200;

    i:not(.ri-close-line) {
      @apply text-orange-600
    }
  }

  &[data-type='info'] {
    @apply bg-blue-200;

    i:not(.ri-close-line) {
      @apply text-black
    }
  }

  &[data-type='error'] {
    @apply bg-red-200;

    i:not(.ri-close-line) {
      @apply animate-pulse text-red-600;
    }
  }
}
</style>