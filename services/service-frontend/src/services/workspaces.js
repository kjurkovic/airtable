import { Api } from "../api"

export default {
  get: async () => {
    return Api.workspace.get()
  },
  delete: async (id) => {
    return Api.workspace.delete(id)
  },
  getMetaModels: async(workspaceId, page) => {
    return Api.workspace.getMetaModels(workspaceId, page)
  },
  getById: async (id) => {
    return Api.workspace.get()
      .then(res => res.data)
      .then(data => data.find(item => item.id == id))
  },
  save: async (name) => {
    return Api.workspace.saveWorkspace(name)
  },
  update: async (id, name) => {
    return Api.workspace.updateWorkspace(id, name)
  }
}