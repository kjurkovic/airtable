import { authorizedClient } from './client';
import { API_WORKSPACE } from './routes';

export const Workspace = {
  get: () => {
    return authorizedClient.get(API_WORKSPACE, {
      validateStatus: (status) => status === 200,
    });
  },
};
