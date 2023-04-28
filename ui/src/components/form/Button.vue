<template>
  <button class="btn">
    <i v-bind:class="{[`ri-${prefixIcon}-line`]: true}" v-if="prefixIcon"></i>
    <slot/>
    <i v-bind:class="{[`ri-${suffixIcon}-line`]: true}" v-if="suffixIcon"></i>
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
    }
  },
  data: () => ({}),
  mounted() {
  },
  methods: {}
})
</script>

<style>
button.btn {
  @apply outline-none inline-flex gap-4 select-none justify-between items-center px-4 py-2 font-semibold leading-6 text-sm shadow rounded-md text-white bg-indigo-600 hover:bg-indigo-500 active:bg-indigo-600
}
</style>