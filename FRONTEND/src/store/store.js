import config from '@/config/config'
import axios from 'axios'
import Vuex from 'vuex'
const API_URL = config.apiUrl + config.apiPort

export default new Vuex.Store({
	state: {
		status: '',
		token: localStorage.getItem('token') || '',
		user: {},
	},
	mutations: {
		auth_request(state) {
			state.status = 'loading'
		},
		auth_success(state, token, user) {
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
		},
	},
	actions: {
		login({ commit }, user) {
			return new Promise((resolve, reject) => {
				commit('auth_request')
				axios({ url: API_URL + '/user/auth', data: user, method: 'POST' })
					.then(resp => {
						console.log(resp)
						const token = resp.data.token
						const user = resp.data.user // ДАТА ПУСТАЯ
						console.log(user)
						localStorage.setItem('token', token)
						axios.defaults.headers.common['Authorization'] = token
						commit('auth_success', token, user)
						console.log(this.state)
						resolve(resp)
					})
					.catch(err => {
						commit('auth_error')
						localStorage.removeItem('token')
						reject(err)
					})
			})
		},
	},
	getters: {},
})