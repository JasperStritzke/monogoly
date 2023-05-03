<template>
  <SidePanel v-model="createGamePanel">
    <template #default="{closePanel}">
      <div class="flex justify-between">
        <h1 class="text-lg font-bold">Create a new game</h1>
        <i class="cursor-pointer text-lg font-medium ri-close-line" @click="closePanel"></i>
      </div>

      <div class="flex flex-col gap-1">
        <label class="text-xs font-bold text-slate-900">Currency <span class="font-normal opacity-50">- no affect on gameplay</span></label>
        <Select v-model="selectedCurrency" v-bind:values="values">
          <template #value="{value}">
            <div class="flex flex-row gap-2">
              <span class="font-medium text-neutral-500" v-html="value.symbol"/>
              <span class="font-medium text-slate-900">{{ value.code }}</span>
            </div>
          </template>

          <template #select-value="{value}">
            <div class="cursor-pointer select-none px-3 py-2 group hover:bg-indigo-600 hover:text-white">
              <div class="flex flex-row gap-2">
                <span class="font-medium text-neutral-500 group-hover:text-white" v-html="value.symbol"/>
                <span class="font-medium text-slate-900 group-hover:text-white">{{ value.code }}</span>
              </div>
            </div>
          </template>
        </Select>
      </div>
      <Field
          label="Amount Over Go" v-model="amountGo" type="number" prefix-icon="money-dollar-circle"
          v-bind:rules="[RequiredRule, MinMaxRule(1)]"
      />

      <Field
          label="Initial Balance" v-model="initialBalance" type="number" prefix-icon="money-dollar-circle"
          v-bind:rules="[RequiredRule, MinMaxRule(1)]"
      />

      <Field
          label="Game Password" placeholder="Your password here..." v-model="password" type="text" prefix-icon="lock-2"
          v-bind:rules="[RequiredRule, MinMaxLengthRule(-1, 32)]"
      />

      <Field
          label="Nickname" placeholder="Enter your nickname..." v-model="name" type="text" prefix-icon="lock-2"
          v-bind:rules="[RequiredRule]"
      />

      <Button suffix-icon="arrow-right" class="mt-4" v-on:click="createGame" v-bind:data-can-start="isValid">Start
        game
      </Button>
    </template>
  </SidePanel>

  <div class="flex h-full flex-col items-center justify-center gap-5">
    <div>
      <p class="text-center text-lg">
        <span class="font-bold">Monogoly</span> - play Monopoly <span class="hidden sm:inline-block"><span
          class="font-bold text-indigo-600 underline underline-offset-2">without</span> the cash</span>
      </p>
    </div>
    <div class="flex flex-col gap-5 sm:w-fit sm:flex-row md:items-center">
      <Button suffix-icon="add" @click="createGamePanel = !createGamePanel">
        Create a game
      </Button>
      <p class="text-center font-medium">or</p>
      <JoinGameButton/>
    </div>
  </div>
</template>

<script setup lang="ts">
import Button from "../components/form/Button.vue";
import JoinGameButton from "../components/JoinGameButton.vue";
import SidePanel from "../components/form/SidePanel.vue";
import Select from "../components/form/Select.vue";
import {nextTick, ref, watch} from "vue";
import currencies, {Currency} from "../util/currency";
import Field from "../components/form/Field.vue";
import {MinMaxLengthRule, MinMaxRule, RequiredRule, useUiForm} from "../components/form/form.util";
import useGameStore from "../store/game.store";

const selectedCurrency = ref(undefined)
const values = ref<Array<Currency>>(currencies)

const createGamePanel = ref(false)

const validate = useUiForm(ref(false))

const amountGo = ref(2000)
const initialBalance = ref(1500)
const password = ref("")
const name = ref("");

const isValid = ref(true)

watch([selectedCurrency, amountGo, initialBalance, password, name], () => nextTick(() => isValid.value = validate()))

function createGame() {
  if (!validate()) return
  useGameStore().createGame({
    name: name.value,
    currency: selectedCurrency.value!,
    amountGo: amountGo.value,
    initialBalance: initialBalance.value,
    password: password.value
  })
}
</script>
<style scoped>
[data-can-start="false"] {
  @apply pointer-events-none opacity-50 grayscale invert
}
</style>
