import { Api } from '../api'
import { EmptyFieldsException, InvalidEmailException, InvalidPasswordException} from './exceptions'

export default {
  execute: (email, password) => {
    function isEmpty (field) { return field == null || field.trim().length == 0 }

    return new Promise((resolve, reject) => {
      if (isEmpty(email) && isEmpty(password)) {
        reject(EmptyFieldsException("Fields are required"))
        return
      } else if (isEmpty(email)) {
        reject(InvalidEmailException("Email is required"))
        return
      } else if (isEmpty(password)) {
        reject(InvalidPasswordException("Password is required"))
        return
      }

      resolve(Api.auth.login(email, password)) 
    })
  }  
} 