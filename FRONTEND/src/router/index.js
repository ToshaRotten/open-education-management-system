import { createRouter, createWebHistory } from 'vue-router'
import AuthPage from '../components/AuthPage'
import IndexPage from '../components/IndexPage'
import MainPage from '../components/MainPage'
import ModulesPage from '../components/ModulesPage'
import RegisterPage from '../components/RegisterPage'
import GeneratePage from "@/components/modules/GeneratePage";
import JournalPage from "@/components/modules/JournalPage";
import HelpPage from "@/components/HelpPage.vue";
import SettingsPage from "@/components/SettingsPage.vue";
import MeetingPage from "@/components/modules/OnlineMeeting/MeetingPage.vue";
import ConversationPage from "@/components/modules/OnlineMeeting/ConversationPage.vue";
import store from '@/store/store'

const ifAuthenticated = (to, from, next) => {
  if (store.getters.isAuthenticated) {
    next()
    return
  }
  next('/auth')
}

const routes = [
  {
    path: '',
    name: 'index',
    component: IndexPage
  },
  {
    path: '/dashboard',
    name: 'dashboard',
    component: MainPage,
    props: true,
    beforeEnter: ifAuthenticated,
  },
  {
    path: '/support',
    name: 'help',
    component: HelpPage,
    beforeEnter: ifAuthenticated,
  },
  {
    path: '/settings',
    name: 'settings',
    component: SettingsPage,
    beforeEnter: ifAuthenticated,
  },
  {
    path: '/modules',
    name: 'modules',
    component: ModulesPage,
    beforeEnter: ifAuthenticated,
  },
  {
    name: 'generate',
    path: '/modules/generate',
    component: GeneratePage,
    beforeEnter: ifAuthenticated,
  },
  {
    name: 'journal',
    path: '/modules/journal',
    component: JournalPage,
    beforeEnter: ifAuthenticated,
  },
  {
    name: 'meet',
    path: '/modules/meet',
    component: MeetingPage,
    beforeEnter: ifAuthenticated,
  },
  {
    name: 'conversation',
    path: '/modules/meet/:roomId',
    component: ConversationPage,
    beforeEnter: ifAuthenticated,
  },
  {
    path: '/auth',
    name: 'auth',
    component: AuthPage,
  },
  {
    path: '/register',
    name: 'register',
    component: RegisterPage,
  }
]

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes,

})

export default router
