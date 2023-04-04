import config from '@/config/config'
import router from '@/router'
import axios from 'axios'

const API_URL = config.apiUrl + config.apiPort + '/user/';
class AuthService {
  login(firstName, hash) {
    return axios
      .post(API_URL + 'auth', {
        firstName: firstName,
        hash: hash
      }, )
      .then(response => {
        console.log(response);
      }).catch(err => {
        console.log(err.response)
        if (err.response.statusText == "Found") {
          router.push(`/dashboard`)
        }
        else {
          alert("Неверное имя пользователя или пароль")
        }
      });
  }
  register(firstName, lastName, thirdName, phone, email, DOB, role, hash) {
    return axios.post(API_URL + 'register', {
        firstName: firstName,
        lastName: lastName,
        thirdName: thirdName,
        phone: phone,
        email: email,
        DOB: DOB,
        role: role,
        hash: hash
    },)
    .then(response => {
        console.log(response);
        
      }).catch(err => {
        console.log(err.response)
        if (err.response.statusText == "") {
          console.log("t")
        }
      });
  }
}
export default new AuthService();