import Vue from 'vue'
import Vuex from 'vuex'
import * as getters from './getters'
import mutations from './mutations'
import actions from './actions'

Vue.use(Vuex)

const state = {
  ownerID: 1,
  ownImg: localStorage.ownImg || "",
  //active menu
  activeIndex: "",
  //user info
  user_id: localStorage.user_id || "",
  email: localStorage.email || "",
  nickname: localStorage.nickname || "",
  token: localStorage.token || "",
  update: localStorage.update || "",
  //login check
  notSignedIn: localStorage.notSignedIn == "true",
  expired: localStorage.expired || "false",
  //article info
  articleHtml: "",

  openSignup: false,
}
  
export default new Vuex.Store({
  state,
  getters,
  mutations,
  actions,
})
  