import Vue from 'vue'
import App from './App.vue'
import ElementUI from 'element-ui'
import 'element-ui/lib/theme-chalk/index.css'
import locale from 'element-ui/lib/locale/lang/en'
import router from './router/index.js'
import './assets/theme/scrollbar.css'
import './assets/theme/container.css'
import './assets/theme/chess.css'
import store from './store'

Vue.config.productionTip = false
Vue.config.productionTip = false
Vue.use(ElementUI, { locale })

new Vue({
  render: h => h(App),
  router,
  store
}).$mount('#app')
