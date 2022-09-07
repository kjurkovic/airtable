import { Outlet } from "react-router-dom"

const Dashboard = () => {
  return (
    <div>
      <nav class="navbar bg-light">
        <div class="container-fluid">
          <a class="navbar-brand" href="/">Airtable</a>
        </div>
      </nav>
      <Outlet />
    </div>
  );
}

export default Dashboard
