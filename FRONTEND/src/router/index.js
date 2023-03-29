import { createRouter, createWebHistory } from 'vue-router'
import MainPage from '../components/MainPage'
import ModulesPage from '../components/ModulesPage'
import AuthPage from '../components/AuthPage'
import RegisterPage from '../components/RegisterPage'

const routes = [
  {
    path: '/dashboard',
    name: 'dashboard',
    component: MainPage
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
