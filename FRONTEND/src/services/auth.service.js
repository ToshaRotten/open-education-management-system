import config from '@/config/config'
import router from '@/router'
import axios from 'axios'

const API_URL = config.apiUrl + config.apiPort + '/user/'
class AuthService {
	login(firstName, hash) {
		return axios
			.post(API_URL + 'auth', {
				firstName: firstName,
				hash: hash,
			})
			.then(response => {
				const token = response.data.token
				localStorage.setItem('user-token', token)
				router.push(`/dashboard`)
			})
			.catch(err => {
				console.log(err.response)
			})
	}
	register(firstName, lastName, thirdName, phone, email, DOB, role, hash) {
		return axios
			.post(API_URL + 'register', {
				firstName: firstName,
				lastName: lastName,
				thirdName: thirdName,
				phone: phone,
				email: email,
				DOB: DOB,
				role: role,
				hash: hash,
			})
			.then(() => {})
			.catch(err => {
				console.log(err.response)
			})
	}
}
export default new AuthService()