import { createRouter, createWebHistory } from 'vue-router'
import AuthPage from '../components/AuthPage'
import IndexPage from '../components/IndexPage'
import MainPage from '../components/MainPage'
import ModulesPage from '../components/ModulesPage'
import RegisterPage from '../components/RegisterPage'
import SchedulePage from '../components/modules/SchedulePage'

const routes = [
  {
    path: '/',
    name: 'index',
    component: IndexPage
  },
  {
    path: '/sch',
    name: 'schedule',
    component: SchedulePage
  },
  {
    path: '/dashboard',
    name: 'dashboard',
    component: MainPage,
    props: true
  },
  {
    path: '/modules',
    name: 'modules',
    component: ModulesPage
  },
  {
    path: '/auth',
    name: 'auth',
    component: AuthPage
  },
  {
    path: '/register',
    name: 'register',
    component: RegisterPage
  }
]

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes
})

export default router
