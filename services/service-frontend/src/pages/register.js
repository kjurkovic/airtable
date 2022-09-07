import { useState } from "react";
import { useNavigate } from "react-router-dom";
import registerService from "../services/register"
import { errors } from "../services/error";

const Register = function() {

  const navigate = useNavigate()

  const [firstName, setFirstName] = useState("")
  const [lastName, setLastName] = useState("")
  const [email, setEmail] = useState("")
  const [password, setPassword] = useState("")
  const [firstNameError, setFirstNameError] = useState(false)
  const [lastNameError, setLastNameError] = useState(false)
  const [emailError, setEmailError] = useState(false)
  const [passwordError, setPasswordError] = useState(false)
  const [emptyError, setEmptyError] = useState(false)

  const handleRegisterSubmit = (event) => {
    event.preventDefault()
  
    registerService
      .execute(firstName, lastName, email, password)
      .then( _ => navigate('/'))
      .catch( e => {
        console.log(e)
        switch (e.code) {
          case errors.INVALID_NAME:
            setFirstNameError(true)
            setLastNameError(true)
            break;
          case errors.INVALID_EMAIL:
            setEmailError(true)
            break;
          case errors.INVALID_PASSWORD:
            setPasswordError(true)
            break;
          case errors.EMPTY_FIELDS:
            setEmptyError(true)
            break;
          default:
            setEmailError(true)
            setPasswordError(true)
        }
      })
  }

  return (
    <div className="container mt-5">
      <form>
        <h1 className="h3 mb-3 fw-normal">Register</h1>
        <div className="form-floating">
          <input type="firstName" className="form-control" id="floatingInput" value={firstName} onChange={ e => setFirstName(e.target.value)} />
          <label htmlFor="floatingInput">First Name</label>
          { firstNameError ? <p className="text-danger">Name is required</p> : <></>}
        </div>
        <div className="form-floating mt-2">
          <input type="lastName" className="form-control" id="floatingInput" value={lastName} onChange={ e => setLastName(e.target.value)} />
          <label htmlFor="floatingInput">Last Name</label>
          { lastNameError ? <p className="text-danger">Name is required</p> : <></>}
        </div>
        <div className="form-floating mt-2">
          <input type="email" className="form-control" id="floatingInput" value={email} onChange={ e => setEmail(e.target.value)} />
          <label htmlFor="floatingInput">Email address</label>
          { emailError ? <p className="text-danger">Email is required</p> : <></>}
        </div>
        <div className="form-floating mt-2">
          <input type="password" className="form-control" id="floatingPassword" value={password} onChange={ e => setPassword(e.target.value)} />
          <label htmlFor="floatingPassword">Password</label>
          { passwordError ? <p className="text-danger">Password has to have 12 characters, 1 upper case letter, 1 lower case, 1 number and 1 special character</p> : <></>}
        </div>
        <button className="w-100 btn btn-lg btn-primary mt-5" type="submit" onClick={ e => handleRegisterSubmit(e)}>Sign in</button>
        { emptyError ? <p className="text-danger text-center mt-2">All fields are required</p> : <></>}
      </form>

      <p className="mt-3 text-center">Have an account? <a href="/login" className="text-decoration-underline">Login.</a></p>
    </div>
  );
}

export default Register
