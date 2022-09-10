import { useEffect, useState } from "react"
import metaService from "../../services/meta"
import Table from 'react-bootstrap/Table';

const MetaData = (props) => {

  const [data, setData] = useState([])
  const [page, setPage] = useState(0)
  const [pages, setPages] = useState([])
  
  useEffect(() => {
    metaService
      .get(props.id)
      .then(res => {
        let entries = res.data.content.map(item => item.content)
        setData(entries)

        let pageCount = []
        for (let i = 0; i < res.data.pagination.totalPages; i++) {
          pageCount.push(i)
        }
        setPages(pageCount)
      })
  }, [])

  
  return (
    <div>
      { props.fields != null 
          ? <Table striped bordered hover>
              <thead>
                <tr>
                  { props.fields.map((field) => <th>{field.label}</th>) }
                </tr>
              </thead>
              <tbody>
                { data.length != 0 ? data.map(item => 
                  <tr key={item}>
                    { props.fields.map((field) => <td key={`${item[field.label]}`}>{item[field.label]}</td>)}
                  </tr>
                ) : <></>}
              </tbody>
            </Table>
          : <p>No data available</p>
      }
        
      <div className="mt-3">
        { pages.length > 1 
            ? <nav aria-label="Page navigation example">
                <ul className="pagination">
                  { pages.map((item) => <li key={item} className={ item == page ? "page-item active" : "page-item"}><a className="page-link" href="#" onClick={() => setPage(item)}>{item + 1}</a></li>) }
                </ul>
              </nav>
            : <></>}
        </div>
    </div>
  )
}

export default MetaData