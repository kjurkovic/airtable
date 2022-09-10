import { Route, Routes, Navigate, useLocation } from "react-router-dom";
import { getRefreshToken } from "./storage/localstorage";
import Login from "./pages/login";
import Register from "./pages/register";
import Dashboard from "./pages/dashboard";
import Workspace from "./pages/workspace";
import NoMatch from "./pages/nomatch";
import CreateWorkspace from "./pages/createworkspace";
import WorkspaceDetails from "./pages/workspacedetails";
import DataForm from "./pages/data";
import CreateMeta from "./pages/createmeta";

function App() {
  return (
    <Routes>
      <Route path="/" element={
        <Protected>
          <Dashboard />
        </Protected>
      }>
        <Route index element={<Workspace />} />
        <Route path="/workspace/add" element={<CreateWorkspace />} />
        <Route path="/workspace/edit/:id" element={<CreateWorkspace />} />
        <Route path="/workspace/:id" element={<WorkspaceDetails />} />
        <Route path="/workspace/:id/add" element={<CreateMeta />} />
      </Route>
      <Route path="/login" element={<Login />} />
      <Route path="/register" element={<Register />} />
      <Route path="/data/:id" element={<DataForm />} />
      <Route path="*" element={<NoMatch />} />
    </Routes>    
  );
}

const Protected = (props) => {
  let location = useLocation();

  if (!getRefreshToken()) {
    // Redirect them to the /login page, but save the current location they were
    // trying to go to when they were redirected. This allows us to send them
    // along to that page after they login, which is a nicer user experience
    // than dropping them off on the home page.
    return <Navigate to="/login" state={{ from: location }} replace />;
  }

  return props.children;
}

export default App;
