export const errors =  {
  INVALID_NAME: "INVALID_NAME",
  INVALID_EMAIL: "INVALID_EMAIL",
  INVALID_PASSWORD: "INVALID_PASSWORD",
  EMPTY_FIELDS: "EMPTY_FIELDS",
}

export const createException = function(message, code) {
  const error = new Error(message)
  error.code = code
  return error
}
