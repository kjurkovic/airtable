import { useEffect, useState } from "react";
import { Link } from "react-router-dom";
import workspaceService from "../services/workspaces"

const Workspace = () => {

  const [workspaces, setWorkspaces] = useState([])
  const [fetchError, setFetchError] = useState(false)

  useEffect(() => {
    workspaceService.get()
      .then(data => setWorkspaces(data))
      .catch(_ => setFetchError(true))
  }, [])

  const handleCardClick = (workspace) => {
    
  }

  return (
    <div class="d-flex p-4">
      {setFetchError ? <p className="text-danger">Došlo je do greške prilikom dohvata radnih okruženja.</p> : <></>}
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
