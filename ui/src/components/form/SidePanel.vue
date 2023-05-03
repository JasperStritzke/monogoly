<template>
  <Teleport to="#app">
    <Transition name="opacity">
      <div class="fixed inset-0 bg-black opacity-50" v-if="modelValue">
      </div>
    </Transition>
    <Transition name="slide-right">
      <div
          v-if="modelValue"
          class="fixed top-0 right-0 bottom-0 flex w-full sm:w-96 select-none flex-col gap-3 bg-slate-50 p-5 shadow"
          v-click-outside="closePanel"
      >
        <slot :closePanel="closePanel"/>
      </div>
    </Transition>
  </Teleport>
</template>

<script setup lang="ts">
import {toRef} from "vue";

const props = defineProps<{ modelValue: boolean }>()
const modelValue = toRef(props, "modelValue")
const emit = defineEmits<{ (e: 'update:modelValue', value: boolean): void }>()

function closePanel() {
  emit('update:modelValue', false)
}
</script>

<style scoped>
</style>