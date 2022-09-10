import { useEffect, useState } from "react"
import Button from 'react-bootstrap/Button';
import Form from 'react-bootstrap/Form';
import FloatingLabel from 'react-bootstrap/FloatingLabel';
import { useNavigate, useParams } from "react-router-dom";
import workspaceService from "../services/workspaces"

const DataSentConfirmation = () => {

  const navigate = useNavigate()

  const goBack = () => navigate(-1)

  return (
    <div className="mt-5 text-center cover-container d-flex w-100 h-100 p-3 mx-auto flex-column">
      <main className="px-3">
        <h1>Workspace saved</h1>
        <p className="lead">Checkout list of your workspaces</p>
        <p className="lead">
          <button className="btn btn-lg btn-primary" onClick={() => goBack()}>Show workspaces</button>
        </p>
      </main>
    </div>
  )
}


const CreateWorkspace = (props) => {

  const params = useParams()
  const [name, setName] = useState("")
  const [saveSuccess, setSaveSuccess] = useState(false)
  const [title, setTitle] = useState("")
  const [workspace, setWorkspace] = useState({})

  useEffect(() => {
    console.log(params.id)
    if (params.id != null) {
      setTitle("Update workspace")

      workspaceService
        .getById(params.id)
        .then(res => {
          setName(res.name)
          setWorkspace(res)
        })
    } else {
      setTitle("Create workspace")
    }
  }, [])

  const save = () => {
    if (name.trim().length == 0) {
      alert("Name is mandatory")
      return
    }

    if (params.id != null) {
      workspaceService.update(params.id, name).then(_ => setSaveSuccess(true))
    } else {
      workspaceService.save(name).then(_ => setSaveSuccess(true))
    }
  }
  
  return (
    <div className="container">
      { saveSuccess 
      ? <DataSentConfirmation />
      : <div>
        <h3 className="my-5">{title}</h3>
        
        <FloatingLabel
          controlId="floatingInput"
          label="Workspace name"
          className="mb-3"
        >
          <Form.Control type="text" value={name} onChange={e => setName(e.target.value)} />
        </FloatingLabel>
        
        <Button variant="primary" type="submit" onClick={() => save()}>Save workspace</Button>
        
      </div>
      }
    </div>
  )
}

export default CreateWorkspace