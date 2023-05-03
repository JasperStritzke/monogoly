import {createApp} from 'vue'
import {createPinia} from "pinia";

import 'remixicon/fonts/remixicon.css'
import './style/main.scss'

import App from './App.vue'
import router from "./plugins/router";
import axios from "axios";
import {clickOutSide} from "./directives/clickOutSide";

axios.defaults.baseURL = "/api/"

createApp(App)
    .use(router)
    .directive("click-outside", clickOutSide)
    .use(createPinia())
    .mount('#app')