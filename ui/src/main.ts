import {createApp} from 'vue'
import {createPinia} from "pinia";

import 'remixicon/fonts/remixicon.css'
import './style/main.scss'

import App from './App.vue'
import router from "./plugins/router";
import axios from "axios";

axios.defaults.baseURL = "/api/"

createApp(App)
    .use(router)
    .use(createPinia())
    .mount('#app')