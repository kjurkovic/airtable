import { useEffect, useState } from "react";
import { useNavigate } from "react-router-dom";
import workspaceService from "../services/workspaces"
import Button from 'react-bootstrap/Button';
import Card from 'react-bootstrap/Card';
import DeleteModal from '../components/deletemodal';

const EmptyWorkspace = () => {
  return (
    <div className="mt-5 text-center cover-container d-flex w-100 h-100 p-3 mx-auto flex-column">
      <main className="px-3">
        <h1>Start by creating workspace</h1>
        <p className="lead">Create workspace where you'll be able to define your forms.</p>
        <p className="lead">
          <a href="/workspace/add" className="btn btn-lg btn-primary">Create</a>
        </p>
      </main>
    </div>
  )
}

const Workspace = () => {

  const navigate = useNavigate()

  const [reload, setReload] = useState(0)
  const [workspaces, setWorkspaces] = useState([])
  const [fetchError, setFetchError] = useState(false)
  const [workspacesEmpty, setWorkspacesEmpty] = useState(false)
  const [modalShow, setModalShow] = useState(false);

  useEffect(() => {
    workspaceService.get()
      .then(res => {
        setFetchError(false)
        setWorkspacesEmpty(res.data.length == 0)
        setWorkspaces(res.data)
      })
      .catch(_ => setFetchError(true))
  }, [reload])

  const handleWorkspaceViewClick = (workspace) => navigate(`/workspace/${workspace.id}`)
  const handleWorkspaceEdit = (workspace) => navigate(`/workspace/edit/${workspace.id}`)
  const handleWorkspaceDelete = (workspaceId) => workspaceService.delete(workspaceId).then(_ => setReload(reload + 1))
  
  return (
    <div>
      { fetchError ? <p className="text-danger">Došlo je do greške prilikom dohvata radnih okruženja.</p> : <></>}
      { workspacesEmpty 
          ? <EmptyWorkspace /> 
          : <div> 
              <div className="m-4">
                <a href="/workspace/add" className="btn btn-lg btn-primary">New Workspace</a>
              </div>
              <div className="d-flex p-4">
                {workspaces.map(workspace => (
                  <Card className="me-2 pe-auto" key={workspace.id}>
                    <Card.Body>
                      <Card.Title>{workspace.name}</Card.Title>
                      <Card.Text></Card.Text>
                      <Button className="me-2" variant="outline-secondary" onClick={() => handleWorkspaceEdit(workspace)}>
                        <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" fill="currentColor" class="bi bi-pen" viewBox="0 0 16 16">
                          <path d="m13.498.795.149-.149a1.207 1.207 0 1 1 1.707 1.708l-.149.148a1.5 1.5 0 0 1-.059 2.059L4.854 14.854a.5.5 0 0 1-.233.131l-4 1a.5.5 0 0 1-.606-.606l1-4a.5.5 0 0 1 .131-.232l9.642-9.642a.5.5 0 0 0-.642.056L6.854 4.854a.5.5 0 1 1-.708-.708L9.44.854A1.5 1.5 0 0 1 11.5.796a1.5 1.5 0 0 1 1.998-.001zm-.644.766a.5.5 0 0 0-.707 0L1.95 11.756l-.764 3.057 3.057-.764L14.44 3.854a.5.5 0 0 0 0-.708l-1.585-1.585z"/>
                        </svg>
                      </Button>
                      <DeleteModal show={modalShow} title={workspace.name} message="This will remove all forms contained within your workspace" onDelete={() => handleWorkspaceDelete(workspace.id)} onHide={() => setModalShow(false)} />
                      <Button className="ms-2" variant="outline-primary" onClick={() => handleWorkspaceViewClick(workspace)}>
                        <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" fill="currentColor" class="bi bi-list-nested" viewBox="0 0 16 16">
                          <path fill-rule="evenodd" d="M4.5 11.5A.5.5 0 0 1 5 11h10a.5.5 0 0 1 0 1H5a.5.5 0 0 1-.5-.5zm-2-4A.5.5 0 0 1 3 7h10a.5.5 0 0 1 0 1H3a.5.5 0 0 1-.5-.5zm-2-4A.5.5 0 0 1 1 3h10a.5.5 0 0 1 0 1H1a.5.5 0 0 1-.5-.5z"/>
                        </svg>
                      </Button>
                    </Card.Body>
                    <Card.Footer>
                      <small className="text-muted">{workspace.id}</small>
                    </Card.Footer>
                </Card>
              ))}
            </div>
          </div>
    }
  </div>
  );
}

export default Workspace
