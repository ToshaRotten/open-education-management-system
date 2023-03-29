export default class User {
    constructor(firstName, lastName, thirdName, dob, role, hash) {
      this.firstName = firstName;
      this.lastName = lastName;
      this.thirdName = thirdName;
      this.dob = dob;
      this.role = role;
      this.hash = hash;
    }
  }