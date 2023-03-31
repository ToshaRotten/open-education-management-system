import axios from 'axios';
import router from '@/router';

const API_URL = 'http://localhost:5671/user/';
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
      });
  }
}
export default new AuthService();