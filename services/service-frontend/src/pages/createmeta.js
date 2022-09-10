import { useNavigate, useParams } from "react-router-dom"
import Table from 'react-bootstrap/Table';
import { useEffect, useState } from "react";
import Form from 'react-bootstrap/Form';
import Button from 'react-bootstrap/Button';
import FloatingLabel from 'react-bootstrap/FloatingLabel';
import metaService from "../services/meta"

const validationExample = {
  color: "",
  date: "",
  datetime: "",
  email: "",
  number: "^[0-9]$+",
  tel: "^[0-9]$+",
  text: "",
  time: "",
  url: "",
}

const SaveConfirmation = () => {

  const navigate = useNavigate()

  const goBack = () => navigate(-1)

  return (
    <div className="mt-5 text-center cover-container d-flex w-100 h-100 p-3 mx-auto flex-column">
      <main className="px-3">
        <h1>Form saved</h1>
        <p className="lead">Checkout list of your forms</p>
        <p className="lead">
          <button className="btn btn-lg btn-primary" onClick={() => goBack()}>Show forms</button>
        </p>
      </main>
    </div>
  )
}


const CreateMeta = (props) => {
  let params = useParams()

  const [fields, setFields] = useState([])
  const [saveSuccess, setSaveSuccess] = useState(false)
  const [formName, setFormName] = useState("")
  const [type, setType] = useState("")
  const [label, setLabel] = useState("")
  const [validation, setValidation] = useState("")

  useEffect(() => {
    setValidation(validationExample[type])
  }, [type])

  const save = () => {
    if (formName.trim().length == 0 || fields.length == 0) {
      alert("Populate form name and add at least one type to save form")
      return
    }

    metaService
      .saveMeta(params.id, formName, fields)
      .then((res) => setSaveSuccess(true))

  }

  const removeField = (fieldLabel) => {
    let filtered = fields.filter((item) => item.label != fieldLabel)
    setFields(filtered)
  }

  const addField = () => {
    if (type.trim().length == 0 || label.trim().length == 0) {
      alert("Populate label and type to add new field")
      return
    }

    let fieldsCopy = fields.slice()
    fieldsCopy.push({
      label: label,
      type: type,
      validation: validation,
    })

    setFields(fieldsCopy)
    setType("")
    setLabel("")
    setValidation("")
  }

  return (
    <div className="container mt-5">
      { saveSuccess ? <SaveConfirmation /> : 

      (<div>
      <h1 className="my-3">Create new form</h1>
      <div className="d-flex my-2">
       <FloatingLabel controlId="floatingInput" label="Form name">
          <Form.Control type="text" value={formName} onChange={(e) => setFormName(e.target.value)} />
        </FloatingLabel>
        <Button className="ms-3" variant="primary" onClick={() => save()}>Save form</Button>
      </div>
      <Table bordered hover>
        <thead>
          <tr>
            <th>Label</th>
            <th>Type</th>
            <th>Validation</th>
            <th></th>
          </tr>
        </thead>
        <tbody>
          { fields.map( field => 
            <tr key={field.label}>
              <td>{field.label}</td>
              <td>{field.type}</td>
              <td>{field.validation}</td>
              <td className="text-center">
               <Button variant="outline-danger" onClick={() => removeField(field.label)}>
                <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" fill="currentColor" class="bi bi-x-octagon" viewBox="0 0 16 16">
                  <path d="M4.54.146A.5.5 0 0 1 4.893 0h6.214a.5.5 0 0 1 .353.146l4.394 4.394a.5.5 0 0 1 .146.353v6.214a.5.5 0 0 1-.146.353l-4.394 4.394a.5.5 0 0 1-.353.146H4.893a.5.5 0 0 1-.353-.146L.146 11.46A.5.5 0 0 1 0 11.107V4.893a.5.5 0 0 1 .146-.353L4.54.146zM5.1 1 1 5.1v5.8L5.1 15h5.8l4.1-4.1V5.1L10.9 1H5.1z"/>
                  <path d="M4.646 4.646a.5.5 0 0 1 .708 0L8 7.293l2.646-2.647a.5.5 0 0 1 .708.708L8.707 8l2.647 2.646a.5.5 0 0 1-.708.708L8 8.707l-2.646 2.647a.5.5 0 0 1-.708-.708L7.293 8 4.646 5.354a.5.5 0 0 1 0-.708z"/>
                </svg>
                </Button>
              </td>
            </tr>
          )}
        </tbody>
        <tfoot>
          <tr>
            <td>
              <FloatingLabel controlId="floatingInput" label="Label">
                <Form.Control type="text" value={label} onChange={(e) => setLabel(e.target.value)} />
              </FloatingLabel>
            </td>
            <td>
              <Form.Select size="lg" aria-label="Default select example" onChange={(e) => setType(e.target.value)}>
                <option>Select type</option>              
                <option value="color">color</option>
                <option value="date">date</option>
                <option value="datetime">datetime</option>
                <option value="email">email</option>
                <option value="number">number</option>
                <option value="tel">tel</option>
                <option value="text">text</option>
                <option value="time">time</option>
                <option value="url">url</option>
              </Form.Select>
            </td>
            <td colSpan={2}>
              <FloatingLabel controlId="floatingInput" label="Validation">
                <Form.Control type="text" value={validation} onChange={(e) => setValidation(e.target.value)} />
              </FloatingLabel>
            </td>
          </tr>
        </tfoot>
      </Table>
      <Button variant="outline-primary" onClick={() => addField()}>Add field</Button>
      </div>
      )}
    </div>
  )
}

export default CreateMeta
