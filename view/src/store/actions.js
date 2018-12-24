import Ajax from '../ajax'

const actions = {
  SIGN_UP: ({commit}, user) => Ajax({
    method: 'post',
    url: 'signup',
    data: user
  }).then((res) => {
    commit('SIGN_UP', user)
  }),
  
  SIGN_IN: ({commit}, user) => Ajax({
    method: 'post',
    url: 'signin',
    data: user
  }).then((res) => {
    commit('SIGN_IN', res.data)
  }),

  NEW_ARTICLE: ({commit}, article) => Ajax({
    method: 'post',
    url: 'article/new',
    data: {
      'user_id': article.user_id,
      'nickname': article.nickname,
      'title': article.title,
      'markdown': article.markdown,
      'html': article.html,
    }
  }),

  EDIT_ARTICLE: ({commit}, article) => Ajax({
    method: 'post',
    url: 'article/edit',
    data: {
      'id': article.id,
      'user_id': article.user_id,
      'nickname': article.nickname,
      'title': article.title,
      'markdown': article.markdown,
      'html': article.html,
    }
  }),

  DEL_ARTICLE: ({commit}, article) => Ajax({
    method: 'post',
    url: 'article/delete',
    data: {
      'id': article.id,
      'user_id': article.user_id,
    }
  }),

  ARTICLE: ({commit}, id) => Ajax({
    method: 'get',
    url: 'article/' + id
  }),

  CAREERS: ({commit}, careers) => Ajax({
    method: 'post',
    url: 'user/careers',
    data: careers,
  }),

  GET_CAREERS: ({commit}, id) => Ajax({
    method: 'get',
    url: 'user/careers/' + id,
  }),

  DEL_CAREER: ({commit}, career) => Ajax({
    method: 'post',
    url: 'user/delcareer',
    data: {
      'id': career.id,
      'user_id': career.user_id,
    }
  }),

  EDUCATION: ({commit}, careers) => Ajax({
    method: 'post',
    url: 'user/educations',
    data: careers,
  }),

  GET_EDUCATION: ({commit}, id) => Ajax({
    method: 'get',
    url: 'user/educations/' + id,
  }),

  DEL_EDUCATION: ({commit}, education) => Ajax({
    method: 'post',
    url: 'user/deleducation',
    data: {
      'id': education.id,
      'user_id': education.user_id,
    }
  }),

  EDIT_USER_PROFILE: ({commit}, user) => Ajax({
    method: 'post',
    url: 'user/edituser',
    data: user,
  }),

  GET_USER_PROFILE: ({commit}, id) => Ajax({
    method: 'get',
    url: 'user/profile/' + id,
  }),

  GET_USER_PROFILE_IMAGE: ({commit}, id) => Ajax({
    method: 'get',
    url: 'user/profileimg/' + id,
  }).then((res) => {
    commit('UPDATE_OWNIMG_IMG', res.data.img)
  }),

  ARTICLES: ({commit}, id) => Ajax({
    method: 'get',
    url: 'articles/' + id,
  }),
}
export default actions
