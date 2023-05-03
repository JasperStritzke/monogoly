<template>
  <Button class="relative min-w-fit join"
          v-on:click="openButton(true)"
          v-on:keydown.enter="openButton(false)"
          v-on:keydown.tab.prevent
          v-bind:data-opened="opened"
          v-bind:data-step="currentStep"
          v-bind:loading="isLoading"
  >
    <div class="relative ml-4 flex h-full w-full flex-row items-center justify-between overflow-hidden py-2 pr-4">
      <span v-bind:class="{[`opacity-0`]: opened}">Join a game</span>
      <div class="absolute my-2 flex w-24 w-full flex-col justify-center gap-1 field-wrapper" data-field-name="code"
           v-if="opened">
        <div class="flex items-center gap-2">
          <i class="text-sm opacity-50 ri-hashtag"></i>
          <input ref="codeInput" placeholder="Code" type="tel" class="w-full font-normal" v-model="code" minlength="4"
                 maxlength="4" v-bind:disabled="isLoading"
          />
        </div>
      </div>
      <div class="absolute my-2 flex w-24 w-full flex-col justify-center gap-1 field-wrapper" data-field-name="password"
           v-if="opened">
        <div class="flex items-center gap-2">
          <i class="text-sm opacity-50 ri-lock-line"></i>
          <input ref="passwordInput" placeholder="Password" type="password" class="w-full font-normal"
                 v-model="password" minlength="1" v-bind:disabled="isLoading"/>
        </div>
      </div>
      <div class="absolute my-2 flex w-24 w-full flex-col justify-center gap-1 field-wrapper" data-field-name="username"
           v-if="opened">
        <div class="flex items-center gap-2">
          <i class="text-sm opacity-50 ri-user-line"></i>
          <input ref="usernameInput" placeholder="Username" class="w-full font-normal"
                 v-model="username" minlength="2" maxlength="10" v-bind:disabled="isLoading"/>
        </div>
      </div>
    </div>
    <span class="mr-4 w-3.5" v-if="opened"></span>
    <div class="mr-4 icon" v-bind:data-valid="isValid" @click="openButton(false)">
      <i class="ri-arrow-right-line" v-if="!isLoading"></i>
      <i class="animate-spin ri-loader-4-line" v-if="isLoading"></i>
    </div>
  </Button>
</template>
<script setup lang="ts">
import Button from "./form/Button.vue";
import {nextTick, ref, watch} from "vue";
import useGameStore from "../store/game.store";

const codeInput = ref<HTMLInputElement>()
const passwordInput = ref<HTMLInputElement>()
const usernameInput = ref<HTMLInputElement>()

const opened = ref(false)

type CurrentStep = "code" | "password" | "username"
const currentStep = ref<CurrentStep>("code")
const isLoading = ref(false)

let code = ref("")
let password = ref("")
let username = ref("")

let isValid = ref(false)

watch([code, password, username, currentStep], (value) => {
  const [localCode, localPassword, localUsername] = value

  switch (currentStep.value) {
    case "code":
      isValid.value = !!localCode && localCode.length == 4;
      return
    case "password":
      isValid.value = !!localPassword && localPassword.length > 0
      return
    case "username":
      isValid.value = !!localUsername && localUsername.length >= 2
  }

})

/**
 *
 * @param clickFromContainer indicates, that the button should only be opened with the code - not any further.
 */
function openButton(clickFromContainer: boolean) {
  if (isLoading.value) return

  //make entire button clickable to focus
  if (clickFromContainer && opened.value) return

  if (!opened.value) {
    opened.value = true
    nextTick(() => codeInput.value!.focus())
    return
  }

  if (!isValid.value) return

  if (currentStep.value == "code") {
    currentStep.value = "password"
    setTimeout(() => passwordInput.value!.focus(), 150)
    return
  }

  if (currentStep.value == "password") {
    currentStep.value = "username"
    setTimeout(() => usernameInput.value!.focus(), 150)
    return
  }

  if (!clickFromContainer) {
    console.log("Trying to join")
    tryJoin()
  }
}

async function tryJoin() {
  isLoading.value = true

  const game = await useGameStore().joinGame(username.value, code.value, password.value)

  if (game) {
    //redirect to game
    return
  }

  code.value = ""
  password.value = ""
  username.value = ""

  currentStep.value = "code"
  isLoading.value = false

  setTimeout(() => codeInput.value!.focus(), 150)
}
</script>

<style scoped lang="scss">
button {
  padding: 0 !important;
}

button input {
  padding-top: 10px;
  padding-bottom: 10px;
}

button[data-opened=true] {
  @apply hover:bg-indigo-600
}

button input {
  @apply bg-transparent outline-none placeholder-white placeholder:font-normal
}

button div.field-wrapper {
  @apply translate-x-0 transition-transform duration-150 ease-in-out
}

button[data-step="code"] div.field-wrapper {
  &[data-field-name="code"] {
    @apply translate-x-0
  }

  &[data-field-name="password"] {
    @apply translate-x-full
  }

  &[data-field-name="username"] {
    @apply translate-x-full
  }
}

button[data-step="password"] div.field-wrapper {
  &[data-field-name="code"] {
    @apply -translate-x-full
  }

  &[data-field-name="password"] {
    @apply translate-x-0
  }

  &[data-field-name="username"] {
    @apply translate-x-full
  }
}

button[data-step="username"] div.field-wrapper {
  &[data-field-name="code"] {
    @apply -translate-x-96
  }

  &[data-field-name="password"] {
    @apply -translate-x-96
  }

  &[data-field-name="username"] {
    @apply translate-x-0
  }
}

button div.icon {
  @apply translate-x-0 transition duration-150 ease-in-out grayscale-0 opacity-100
}

button[data-opened=true] div.icon[data-valid="false"] {
  @apply opacity-20
}

button[data-opened=true] div.icon {
  @apply bg-indigo-600 rounded-md absolute right-1 h-full aspect-square flex justify-center items-center translate-x-16 shadow-lg hover:bg-indigo-500 active:bg-indigo-600
}
</style>