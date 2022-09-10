import Button from "react-bootstrap/esm/Button";
import { Outlet, useNavigate } from "react-router-dom"
import { removeTokens } from "../storage/localstorage"

const Dashboard = () => {
  const navigate = useNavigate()

  const signout = () => {
    removeTokens()
    navigate('/')
  }

  return (
    <div>
      <nav className="navbar bg-light">
        <div className="container-fluid">
          <a className="navbar-brand" href="/">Airtable</a>
          <Button variant="light" onClick={() => signout()}>Sign out</Button>
        </div>
      </nav>
      <Outlet />
    </div>
  );
}

export default Dashboard
