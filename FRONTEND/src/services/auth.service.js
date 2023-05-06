import config from '@/config/config'
import router from '@/router'
import axios from 'axios'
import { useToast } from 'vue-toastification'

const API_URL = config.apiUrl + config.apiPort + '/user/'
class AuthService {
	login(firstName, hash) {
		return axios
			.post(API_URL + 'read', {
				firstName: firstName,
				hash: hash,
			})
			.then(response => {
				this.$toast.success("Всё топ")
				const user = response.data.token
				// const token = response.data.user
				localStorage.setItem('user', user)
				router.push(`/dashboard`)
			})
			.catch(err => {
				console.log(err.response)
			})
	}
	register(firstName, lastName, thirdName, phone, email, DOB, role, hash) {
		
		let users = {users:[
			{
				firstName: firstName,
				lastName: lastName,
				thirdName: thirdName,
				phone: phone,
				email: email,
				DOB: DOB,
				role: role,
				hash: hash
			}
		]}
		return axios
			.post(API_URL + 'create', users)
			.then(response => {
				console.log(response)
				const toast = useToast()
				toast.success("Вы успешно зарегистрировались!", {timeout: 2000})
				router.push('/auth')
			})
			.catch(err => {
				const toast = useToast()
				toast.error(err.message, {timeout: 2000})
			})
	}
}
export default new AuthService()
