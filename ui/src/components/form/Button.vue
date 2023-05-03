<template>
  <button class="btn" v-bind="$props" v-bind:data-loading="loading">
    <i v-bind:class="{[`ri-${prefixIcon}-line`]: true}" v-if="prefixIcon"></i>
    <span class="w-full h-full flex justify-center flex-row items-center"><slot/></span>
    <i v-bind:class="{[`ri-${suffixIcon}-line`]: true}" v-if="suffixIcon"></i>

    <i class="ri-loader-3-line loading-icon" v-if="loading"></i>
  </button>
</template>

<script lang="ts">
import {defineComponent} from "vue";

export default defineComponent({
  props: {
    prefixIcon: {
      type: String,
      default: undefined,
    },
    suffixIcon: {
      type: String,
      default: undefined,
    },
    loading: {
      type: Boolean,
      default: false,
    }
  },
  data: () => ({}),
  methods: {}
})
</script>

<style lang="scss">
button.btn {
  @apply relative outline-none h-10 inline-flex gap-5 select-none justify-center items-center px-4 py-2 font-semibold leading-6 text-sm shadow rounded-md text-white bg-indigo-600 hover:bg-indigo-500 active:bg-indigo-600
}

button[data-loading="true"] {
  @apply pointer-events-none;

  & > *:not(.loading-icon) {
    opacity: 0;
  }
}

button i.loading-icon {
  @apply text-lg font-bold fixed animate-spin text-lg;
}
</style>