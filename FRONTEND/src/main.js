import { createApp } from 'vue'
import { createVuetify } from 'vuetify'
import * as components from 'vuetify/components'
import * as directives from 'vuetify/directives'
import App from './App.vue'
import router from './router'
import store from './store/store'

const vuetify = createVuetify({
  components,
  directives,
})

createApp(App).use(router).use(store).use(vuetify).mount('#app')
