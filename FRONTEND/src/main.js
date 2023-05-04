import { createApp } from 'vue'
import { createVuetify } from 'vuetify'
import 'vuetify/styles'
import * as components from 'vuetify/components'
import * as directives from 'vuetify/directives'
import App from './App.vue'
import router from './router'
import store from './store/store'
import Toast from "vue-toastification";
import "vue-toastification/dist/index.css";

const vuetify = createVuetify({
  components,
  directives,
})
const options = {
  transition: "Vue-Toastification__bounce",
  maxToasts: 3,
  newestOnTop: true,
  position: "top-center",
  timeout: 2000,
  closeOnClick: true,
  pauseOnFocusLoss: true,
  pauseOnHover: false,
  draggable: true,
  draggablePercent: 0.7,
  showCloseButtonOnHover: false,
  hideProgressBar: true,
  closeButton: "button",
  icon: true,
  rtl: false,
};


createApp(App).use(vuetify).use(router).use(store).use(Toast, options).mount('#app')
