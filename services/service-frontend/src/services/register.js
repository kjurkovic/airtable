import { Api } from '../api'
import { EmptyFieldsException, InvalidEmailException, InvalidPasswordException, InvalidNameException} from './exceptions'

export default {
  execute: (firstName, lastName, email, password) => {
    function isEmpty (field) { return field == null || field.trim().length == 0 }

    return new Promise((resolve, reject) => {
      if (isEmpty(firstName), isEmpty(lastName) && isEmpty(email) && isEmpty(password)) {
        reject(EmptyFieldsException("Fields are required"))
        return
      } else if (isEmpty(firstName) || isEmpty(lastName)) {
        reject(InvalidNameException("Name is required"))
        return
      } else if (isEmpty(email)) {
        reject(InvalidEmailException("Email is required"))
        return
      } else if (isEmpty(password)) {
        reject(InvalidPasswordException("Password is required"))
        return
      }

      resolve(Api.auth.register(firstName, lastName, email, password)) 
    })
  }  
} 