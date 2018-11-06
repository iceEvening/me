import Vue from 'vue'
import Router from 'vue-router'
import Me from '@/components/Me'
import Signup from '@/components/Signup'
import Signin from '@/components/Signin'
import New from '@/components/New'
import Article from '@/components/Article'
import Careers from '@/components/Careers'
import Educations from '@/components/Educations'
import Profile from '@/components/Profile'
import Articles from '@/components/Articles'
import EditArticle from '@/components/EditArticle'

Vue.use(Router)

const router = new Router({
  base: '/', 
  mode: 'history',
  routes: [
    {
      path: '/me',
      name: 'me',
      component: Me,
      meta: {
        title: 'Homepage'
      }
    },
    {
      path: '/',
      redirect: '/me'
    },
    {
      path: '*',
      redirect: '/me'
    },
    {
      path: '/signup',
      name: 'signup',
      component: Signup,
      meta: {
        title: 'Signup'
      }
    },
    {
      path: '/signin',
      name: 'signin',
      component: Signin,
      meta: {
        title: 'Signin'
      }
    },
    {
      path: '/new',
      name: 'new',
      component: New,
      meta: {
        title: 'New'
      }
    },
    {
      path: '/article/:id',
      name: 'article',
      component: Article,
      meta: {
        title: 'Article'
      }
    },
    {
      path: '/careers',
      name: 'careers',
      component: Careers,
      meta: {
        title: 'Careers'
      }
    },
    {
      path: '/educations',
      name: 'educations',
      component: Educations,
      meta: {
        title: 'Educations'
      }
    },
    {
      path: '/profile',
      name: 'profile',
      component: Profile,
      meta: {
        title: 'Profile'
      }
    },
    {
      path: '/articles',
      name: 'articles',
      component: Articles,
      meta: {
        title: 'Articles'
      }
    },
    {
      path: '/editarticle/:id',
      name: 'editarticle',
      component: EditArticle,
      meta: {
        title: 'Edit article'
      }
    },
  ]
})

router.beforeEach((to, from, next) => {
  if (to.meta.title) {
    document.title = 'Me|' + to.meta.title
  }
  next()
})

export default router
