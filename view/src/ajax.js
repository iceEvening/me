import axios from 'axios'
import store from './store'

axios.interceptors.request.use(
  config => {
    config.baseURL = 'http://www.sanghongwei:8081/'
    config.timeout = 1800000
    config.headers = {
      'Content-Type': 'application/x-www-form-urlencoded;charset=utf-8',
    }
    if (store.state.token != '') {
      config.headers['X-USER-TOKEN'] = store.state.token
    }
    console.log('token' + store.state.token)
    console.log(config)
    return config
  },
  error => {
    return Promise.reject(error);
  }
)

axios.interceptors.response.use(
  response => {
    if (response.headers['x-user-token'] != '' && typeof(response.headers['x-user-token']) != 'undefined') {
      store.state.token = response.headers['x-user-token']
    }
    if (response.data.msg="授权已过期" && response.data.status == -1) {
      store.state.expired = true
    }
    return response
  },
  error => {
    return Promise.reject(error)
  }
)

export default axios
