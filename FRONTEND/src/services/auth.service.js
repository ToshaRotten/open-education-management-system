import axios from 'axios';

const API_URL = 'http://localhost:5671/user/';
class AuthService {
  login(firstName, hash) {
    return axios
      .post(API_URL + 'auth', {
        firstName: firstName,
        lastName: '1',
        thirdName: '1',
        DOB: '1',
        role: '1',
        hash: hash
      }, )
      .then(response => {
        console.log(response);
      }).catch(err => {
        console.log(err.response)
      });
  }
  register(firstName, hash) {
    return axios.post(API_URL + 'register', {
        firstName: firstName,
        lastName: '1',
        thirdName: '1',
        DOB: '1',
        role: '1',
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