<template>
  <div class="flex flex-col h-full items-center justify-center gap-5">
    <div>
      <p class="text-lg text-center">
        <span class="font-bold">Monogoly</span> - play Monopoly <span class="hidden sm:inline-block"><span class="underline underline-offset-2 text-indigo-600 font-bold">without</span> the cash</span>
      </p>
    </div>
    <div class="flex flex-col sm:flex-row md:items-center gap-5 sm:w-fit">
      <CreateGameButton />
      <p class="font-medium text-center">or</p>
      <JoinGameButton />
    </div>
  </div>
</template>

<script setup lang="ts">
import CreateGameButton from "../components/CreateGameButton.vue";
import JoinGameButton from "../components/JoinGameButton.vue";
</script>
