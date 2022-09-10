import Table from 'react-bootstrap/Table';
import { useState } from 'react';
import DeleteModal from '../deletemodal';

const MetaDefinition = (props) => {

  const [modalShow, setModalShow] = useState(false);

  return (
    <div>
      { props.id != null ? <p>Share form link: <a href={`/data/${props.id}`}>{`${window.location.origin}/data/${props.id}`}</a></p> : <></> }
      <Table striped bordered hover>
        <thead>
          <tr>
            <th>Label</th>
            <th>Type</th>
            <th>Validation</th>
          </tr>
        </thead>
        <tbody>
          { props.fields != null ? props.fields.map( field => 
            <tr key={field.id}>
              <td>{field.label}</td>
              <td>{field.type}</td>
              <td>{field.validation}</td>
            </tr>
          ) : <></>}
        </tbody>
      </Table>
      <div className="text-end" >
        <DeleteModal 
          show={modalShow} 
          title={props.name} 
          message="This will remove all data contained within your form" 
          buttonText="Delete" 
          onDelete={() => props.deleteForm(props.id)} 
          onHide={() => setModalShow(false)} />
      </div>
    </div>
  )
}

export default MetaDefinition