<template>
  <div class="flex w-full flex-col gap-1">
    <label v-bind:for="id" class="text-xs font-bold text-slate-900">
      {{ label }} <span class="text-red-500" v-if="!!error">- {{ error }}</span>
    </label>
    <div class="relative flex-grow">
      <input
          v-bind:aria-invalid="!!error"
          v-bind:id="id"
          class="flex h-10 w-full cursor-pointer select-none appearance-none items-center justify-between rounded-md border border-gray-200 bg-white p-2 pl-7 shadow-sm outline-none focus:border-indigo-600 placeholder:text-sm"
          v-bind:placeholder="placeholder"
          v-bind:type="type"
          v-model="computedValue"
      />
      <i class="absolute bottom-1/2 left-2 translate-y-1/2 font-medium text-neutral-500"
         v-bind:class="{[`ri-${prefixIcon}-line`]: !!prefixIcon}"></i>
    </div>
  </div>
</template>
<script lang="ts" setup>
import useId from "../../hooks/useId";
import {computed, nextTick, toRef} from "vue";
import {FieldRules, useField} from "./form.util";

type Props = {
  type?: "number" | "text" | "password" | "tel",
  placeholder?: string,
  label: string,
  modelValue: any,
  prefixIcon: string,
  rules?: FieldRules,
}

const id = useId()

const emit = defineEmits<{ (e: 'update:modelValue', value: any): void }>()
const props = withDefaults(
    defineProps<Props>(),
    {
      placeholder: "",
      type: "text",
      required: false,
      rules: () => [],
    }
)

const modelValue = toRef(props, "modelValue")
const rules = toRef(props, "rules")

const {validate, error, isLazyValidation} = useField(modelValue, rules)

const computedValue = computed({
  get(): any {
    return props.modelValue
  },
  set(newValue: any) {
    emit('update:modelValue', newValue)

    //otherwise validate won't know about the change that was made to the state
    nextTick(() => {
      if (!isLazyValidation?.value) validate()
    })
  }
})
</script>
<style scoped lang="scss">
input[aria-invalid="true"] {
  @apply border-red-500;
}
</style>
