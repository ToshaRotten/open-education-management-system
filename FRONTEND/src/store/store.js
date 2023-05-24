import config from '@/config/config'
import router from '@/router'
import axios from 'axios'
import Vuex from 'vuex'
import createPersistedState from "vuex-persistedstate";
import { useToast } from 'vue-toastification'



export default new Vuex.Store({
	state: {
		status: '',
		token: localStorage.getItem('token') || '',
		user: {}
	},
	plugins: [createPersistedState()],
	mutations: {
		auth_request(state) {
			state.status = 'loading'
		},
		auth_success(state, {token, user}) {
			state.status = 'success'
			state.token = token
			state.user = user
		},
		auth_error(state) {
			state.status = 'error'
		},
		logout(state) {
			state.status = ''
			state.token = ''
			state.user = {}
			const toast = useToast();
			toast.info("Вы успешно вышли из аккаунта!", {timeout: 2000})
		},
	},
	actions: {
		login({ commit }, user) {
			return new Promise((resolve, reject) => {
				commit('auth_request')
				axios({ url: config.AUTHSERVICE_URL + '/user/read', data: user, method: 'POST' })
					.then(resp => {
						const toast = useToast();
						console.log(resp)
						toast.success("Вы успешно вошли в аккаунт!", {timeout: 2000})
						const token = resp.data.token
						const user = resp.data[0]
						localStorage.setItem('token', token)
						axios.defaults.headers.common['Authorization'] = token
						commit('auth_success', {token, user})
						resolve(resp)
						router.push('/dashboard')
					})
					.catch(err => {
						const toast = useToast();
						toast.error(err.message, {timeout: 2000})
						commit('auth_error')
						localStorage.removeItem('token')
						reject(err)
					})
			})

		},

	},
	getters: {
		loadData(state) {
			return state.user
		},
		isAuthenticated(state) {
			console.log(state.user)
			if (Object.entries(state.user).length === 0) {
				return false
			}
			else return true
		},

	}
})
