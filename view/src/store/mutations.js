import {
  ACTIVE_INDEX,
  SIGN_UP,
  SET_ACTIVE_INDEX_TO_ME,
  SIGN_IN,
  SIGN_OUT,
  UPDATE_PROFILE_TIME,
  UPDATE_OWNIMG_IMG
} from './mutation-types'

const mutations = {
  [ACTIVE_INDEX] (state, index) {
    state.activeIndex = index
  },

  [SET_ACTIVE_INDEX_TO_ME] (state) {
    state.activeIndex = '/me'
  },

  [SIGN_UP] (state, user) {
    state.nickname = user.nickname
  },

  [UPDATE_PROFILE_TIME] (state, update) {
    state.update = update
    localStorage.update = update
  },

  [UPDATE_OWNIMG_IMG] (state, img) {
    state.ownImg = img
    localStorage.ownImg = img
  },

  [SIGN_IN] (state, data) {
    state.user_id = data.user_id
    state.email = data.email
    state.nickname = data.nickname
    state.token = data.token
    state.notSignedIn = false

    localStorage.user_id = data.user_id
    localStorage.email = data.email
    localStorage.nickname = data.nickname
    localStorage.token = data.token
    localStorage.notSignedIn = false
  },

  [SIGN_OUT] (state) {
    state.user_id = ""
    state.nickname = ""
    state.email = ""
    state.token = ""
    state.notSignedIn = true
    state.expired = false

    localStorage.user_id = ""
    localStorage.nickname = ""
    localStorage.email = ""
    localStorage.token = ""
    localStorage.notSignedIn = true
    localStorage.expired = false
  },
}
export default mutations