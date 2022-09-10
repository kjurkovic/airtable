import { authorizedClient, unauthorizedClient } from './client';
import { API_DATA, API_META } from './routes';

export const Meta = {
  getData: (metaId) => {
    return authorizedClient.get(`${API_DATA}/${metaId}`, {
      validateStatus: (status) => status === 200,
    });
  },
  getModel: (metaId) => {
    return unauthorizedClient.get(`${API_META}/${metaId}`)
  },
  addData: (metaId, data) => {
    return unauthorizedClient.post(`${API_DATA}/${metaId}`, {
      content: data
    }, {
      validateStatus: (status) => status === 202,
    })
  },
  addMeta: (workspaceId, formName, fields) => {
    return authorizedClient.post(`${API_META}`, {
      name: formName,
      workspaceId: workspaceId,
      fields: fields,
    }, {
      validateStatus: (status) => status === 201,
    })
  },
  deleteMeta: (metaId) => {
    return authorizedClient.delete(`${API_META}/${metaId}`)
  }
};
