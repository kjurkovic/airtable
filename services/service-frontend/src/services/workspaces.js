import { Api } from "../api"

export default {
  get: async () => {
    return Api.workspace.get()
      .then(response => response.data)
  }
}