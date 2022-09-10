import { Api } from "../api"

export default {
  get: async (metaId) => {
    return Api.meta.getData(metaId)
  },
  getModel: async (metaId) => {
    return Api.meta.getModel(metaId)
  },
  sendData: async (metaId, data) => {
    return Api.meta.addData(metaId, data)
  },
  saveMeta: async(workspaceId, formName, fields) => {
    return Api.meta.addMeta(workspaceId, formName, fields)
  },
  delete: async(metaId) => {
    return Api.meta.deleteMeta(metaId)
  }
}