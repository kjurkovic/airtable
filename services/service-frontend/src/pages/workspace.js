import { useEffect, useState } from "react";
import { Link } from "react-router-dom";

const Workspace = () => {

  const [workspaces, setWorkspaces] = useState([])

  useEffect(() => {

  }, [])

  const handleCardClick = (workspace) => {

  }

  return (
    <div class="d-flex p-4">
      { workspaces.map( workspace => (
        <Link className="link-secondary text-decoration-none" to={`/workspace/${workspace.id}`} >
          <div className="pe-2">
            <div className="card bg-light bg-gradient" onClick={ _ => handleCardClick(workspace)}>
              <div className="card-body">
                <h5 className="card-title">{workspace.name}</h5>
              </div>
            </div>
          </div> 
        </Link>
      ))}
      
    </div>
  );
}

export default Workspace
