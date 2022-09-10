import { useEffect, useState } from "react"
import { useParams } from "react-router-dom";
import Button from 'react-bootstrap/Button';
import metaService from "../services/meta"

const DataSentConfirmation = () => {
  return (
    <div className="mt-5 text-center cover-container d-flex w-100 h-100 p-3 mx-auto flex-column">
      <main className="px-3">
        <h1>Form completed</h1>
        <p className="lead">Thank you for submitting your response.</p>
      </main>
    </div>
  )
}

const DataForm = () => {
  const params = useParams();

  const [model, setModel] = useState({})
  const [input, setInput] = useState({})

  const [dataSent, setDataSent] = useState(false)
  const [dataSentError, setDataSentError] = useState(false)

  const onFieldChange = (event, field) => {
    let inp = structuredClone(input)
    let valid = true
    
    if (field.validation != "" && field.validation != null) {
      const regex = new RegExp(field.validation);
      valid = regex.test(event.target.value)
    } 

    inp[field.label] = {
      valid: valid,
      value: event.target.value
    }

    setInput(inp)
  }

  const submit = () => {
    let keys = Object.keys(input)
    
    if (model.fields.length != keys.length) {
      alert("All fields are required")
      return
    }

    let isValid = true

    let data = Object.entries(input).reduce((prev, [k,v]) => {
      console.log(prev, k, v)
      prev[k] = v.value
      isValid = isValid && v.valid
      return prev
    }, {})

    console.log(data, isValid)

    if (!isValid) {
      alert("All fields are mandatory")
      
    } else {
      metaService
      .sendData(model.id, data)    
      .then(_ => {
        setDataSent(true)
        setDataSentError(false)
      })
      .catch(_ => setDataSentError(true))
    }
  }

  useEffect(() => {
    if (params.id == null) return

    metaService.getModel(params.id)
      .then(res => setModel(res.data))
  }, [])


  return (
    <div className="container p-4">
    { dataSent ? <DataSentConfirmation />
    : model.id != null 
      ? <div className="text-center">
        <h1 className="my-5">{model.name}</h1>
        { model.fields.map(field => 
          <div className="row" key={field.id}>
            <div className="col">
            <div className="form-floating mb-3">
              <input 
                type={field.type} 
                value={input[field.label]?.value || ""} 
                className={input[field.label] == undefined ? "form-control" : (input[field.label]?.valid ? "form-control is-valid" : "form-control is-invalid")}
                id="floatingInput" 
                onChange={(e) => onFieldChange(e, field)} />
              <label htmlFor="floatingInput">{field.label}</label>
              { field.validation != "" && field.validation != null 
                  ? <div class="invalid-feedback text-start">{ field.validation }</div>
                  : <></>
              }
            </div>
            </div>
          </div>
        )}
        <Button className="me-2" variant="primary" onClick={() => submit()}>Submit your response</Button>
      </div> 
      : <p>Loading...</p>}
    </div>
  )
}

export default DataForm