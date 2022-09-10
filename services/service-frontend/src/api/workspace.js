import { authorizedClient } from './client';
import { API_META, API_WORKSPACE } from './routes';

export const Workspace = {
  get: () => {
    return authorizedClient.get(API_WORKSPACE, {
      validateStatus: (status) => status === 200,
    });
  },
  delete: (workspaceId) => {
    return authorizedClient.delete(`${API_WORKSPACE}/${workspaceId}`, {
      validateStatus: (status) => status === 204,
    });
  },
  getMetaModels: (workspaceId, page) => {
    return authorizedClient.get(`${API_META}/work/${workspaceId}?page=${page}`)
  },
  saveWorkspace: (name) => {
    return authorizedClient.post(`${API_WORKSPACE}`, {
      name,
    },{
      validateStatus: (status) => status === 200,
    })
  },
  updateWorkspace: (id, name) => {
    return authorizedClient.put(`${API_WORKSPACE}/${id}`, {
      name,
    },{
      validateStatus: (status) => status === 200,
    })
  }
};
