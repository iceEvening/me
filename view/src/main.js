import Vue from 'vue'
import App from './App.vue'
import router from './router'
import store from './store'
import utils from './utils'
import axios from 'axios'
import mavonEditor from 'mavon-editor'
import markdownitTOC from 'markdown-it-toc'
import emoji from 'markdown-it-emoji'
import 'mavon-editor/dist/css/index.css'
import './plugins/element.js'
import './assets/thirdparty/iconfont/iconfont.css'

Vue.config.productionTip = false
Vue.use(utils)
Vue.use(mavonEditor)
mavonEditor.markdownIt.use(markdownitTOC).use(emoji)

new Vue({
  el: '#app',
  router,
  store,
  axios,
  render: (h) => h(App)
})
