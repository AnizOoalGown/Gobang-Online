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
import VueI18n from 'vue-i18n'

Vue.config.productionTip = false
Vue.config.productionTip = false
Vue.use(ElementUI, { locale })
Vue.use(VueI18n)

const i18n = new VueI18n({
  locale: 'en',
  messages: {
    'en': require('./constants/lang/en'),
    'zh': require('./constants/lang/zh')
  }
})

new Vue({
  render: h => h(App),
  router,
  store,
  i18n
}).$mount('#app')
