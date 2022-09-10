import { useEffect, useState } from "react";
import { useParams } from "react-router-dom";
import workspaceService from "../services/workspaces"
import metaService from "../services/meta"
import Tab from 'react-bootstrap/Tab';
import Tabs from 'react-bootstrap/Tabs';
import MetaData from "../components/meta/data";
import MetaDefinition from "../components/meta/definition";


const EmptyWorkspaceDetails = (props) => {
  return (
    <div className="mt-5 text-center cover-container d-flex w-100 h-100 p-3 mx-auto flex-column">
      <main className="px-3">
        <h1>Start by creating forms</h1>
        <p className="lead">Create form and share link with your users</p>
        <p className="lead">
          <a href={`/workspace/${props.id}/add`} className="btn btn-lg btn-primary">Create</a>
        </p>
      </main>
    </div>
  )
}

const SelectedMeta = (props) => {
  return (
    <Tabs
      defaultActiveKey="definition"
      id="uncontrolled-tab-example"
      className="mb-3">
      <Tab eventKey="definition" title="Definition">
        <MetaDefinition {...props} />
      </Tab>
      <Tab eventKey="data" title="Data">
        <MetaData {...props} />
      </Tab>
    </Tabs>
  )
}

const WorkspaceDetails = () => {
  const params = useParams();

  const [meta, setMeta] = useState([])
  const [page, setPage] = useState(0)
  const [refresh, setRefresh] = useState(0)
  const [pages, setPages] = useState([])
  const [selectedMeta, setSelectedMeta] = useState({})
  
  useEffect(() => {
    if (params.id == null) return

    workspaceService
      .getMetaModels(params.id, page)
      .then(res => {
        setMeta(res.data.content)
        if (res.data.content.length > 0) {
          setSelectedMeta(res.data.content[0])
        }

        let pageCount = []
        for (let i = 0; i < res.data.pagination.totalPages; i++) {
          pageCount.push(i)
        }
        setPages(pageCount)
      })
  }, [page, refresh])

  const deleteForm = (metaId) => {
      metaService.delete(metaId)
        .then(() => {
          setRefresh(refresh + 1)
        })
  }

  
  return (
    <div className="p-4">
      { meta.length === 0 
        ? <EmptyWorkspaceDetails id={params.id} />
        : <div> 
            <div className="d-flex m-4">
              <a href={`/workspace/${params.id}/add`} className="btn btn-lg btn-primary">New Form</a>
            </div>
            <div className="container-fluid p-4">
              <div className="row">
                <div className="col">
                <ul className="list-group">
                  {meta.map(item => 
                    <a href="#" key={item.id} className={item.id == selectedMeta.id ? "list-group-item list-group-item-action active" : "list-group-item list-group-item-action"} onClick={() => setSelectedMeta(item)}>{item.name}</a>
                  )}
                </ul>
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
              <div className="col-9 text-center">
                  { selectedMeta ? <SelectedMeta key={selectedMeta?.id} deleteForm={deleteForm} {...selectedMeta} /> : <></> }
              </div>
            </div>
          </div>
        </div>
      }
    </div>
  )
}

export default WorkspaceDetails