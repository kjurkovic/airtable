import { useState } from "react";
import { useNavigate } from "react-router-dom";
import loginService from "../services/login"
import { errors } from "../services/error";

const Login = function() {

  const navigate = useNavigate()

  const [email, setEmail] = useState("")
  const [password, setPassword] = useState("")
  const [emailError, setEmailError] = useState(false)
  const [passwordError, setPasswordError] = useState(false)
  const [emptyError, setEmptyError] = useState(false)

  const handleLoginSubmit = (event) => {
    event.preventDefault()
  
    loginService
      .execute(email, password)
      .then( _ => navigate('/'))
      .catch( e => {
        switch (e.code) {
          case errors.INVALID_EMAIL:
            setEmailError(true)
            break;
          case errors.INVALID_PASSWORD:
            setPasswordError(true)
            break;
          default:
            setEmptyError(true)
        }
      })
  }

  return (
    <div className="container mt-5">
      <form>
        <h1 className="h3 mb-3 fw-normal">Sign in</h1>

        <div className="form-floating">
          <input type="email" className="form-control" id="floatingInput" value={email} onChange={ e => setEmail(e.target.value)} />
          <label htmlFor="floatingInput">Email address</label>
          { emailError ? <p className="text-danger">Email is required</p> : <></>}
        </div>
        <div className="form-floating mt-2">
          <input type="password" className="form-control" id="floatingPassword" value={password} onChange={ e => setPassword(e.target.value)} />
          <label htmlFor="floatingPassword">Password</label>
          { passwordError ? <p className="text-danger">Password is required</p> : <></>}
        </div>
        <button className="w-100 btn btn-lg btn-primary mt-5" type="submit" onClick={ e => handleLoginSubmit(e)}>Sign in</button>
        { emptyError ? <p className="text-danger">All fields are required</p> : <></>}
      </form>

      <p className="mt-3 text-center">Don't have an account? <a href="/register" className="text-decoration-underline">Register.</a></p>
    </div>
  );
}

export default Login
