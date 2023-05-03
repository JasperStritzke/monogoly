<template>
  <div class="relative">
    <div
        class="flex h-10 cursor-pointer select-none items-center justify-between rounded-md border border-gray-200 bg-white p-2 shadow-sm"
        v-on:click="expanded = !expanded"
    >
      <slot name="value" v-bind:value="modelValue" v-if="modelValue"/>
      <i class="ri-expand-up-down-line"></i>
    </div>

    <Transition name="slide">
      <div
          class="absolute right-0 left-0 mt-2 flex w-full flex-col rounded-md bg-white py-2 shadow-md font-light z-10 sm:max-h-fit max-h-72 overflow-y-auto"
          v-if="expanded"
          v-click-outside="() => expanded = false"
      >
        <div
            v-for="value in values" v-bind:key="value.code"
            v-on:click="select(value)"
        >
          <slot name="select-value" v-bind:value="value"/>
        </div>
      </div>
    </Transition>
  </div>
</template>

<script setup lang="ts">
import {onMounted, ref} from "vue";

const expanded = ref(false)

const props = defineProps<{ values: Array<any>, modelValue: any, }>()
const emit = defineEmits<{ (e: 'update:modelValue', value: any): void }>()

onMounted(() => {
  if (!props.modelValue) {
    emit("update:modelValue", props.values[0])
  }
})

function select(value: any) {
  emit("update:modelValue", value)
  expanded.value = false
}
</script>

<style scoped>
.slide-enter-active,
.slide-leave-active {
  transition: 120ms ease;
  opacity: 1;
  transform: translateY(0);
}

.slide-enter-from,
.slide-leave-to {
  opacity: 0;
  transform: translateY(-5px);
}
</style>