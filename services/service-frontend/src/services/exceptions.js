import { errors, createException } from './error'

export function InvalidNameException(message) {
  return createException(message, errors.INVALID_NAME)
}
InvalidEmailException.prototype = Object.create(Error.prototype)

export function InvalidEmailException(message) {
  return createException(message, errors.INVALID_EMAIL)
}
InvalidEmailException.prototype = Object.create(Error.prototype)

export function InvalidPasswordException(message) {
  return createException(message, errors.INVALID_PASSWORD)
}

InvalidPasswordException.prototype = Object.create(Error.prototype) 

export function EmptyFieldsException(message) {
  return createException(message, errors.EMPTY_FIELDS)
}

EmptyFieldsException.prototype = Object.create(Error.prototype)