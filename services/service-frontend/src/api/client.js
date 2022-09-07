import axios from 'axios';
import config from '../config';
import createAuthRefreshInterceptor from 'axios-auth-refresh';
import { API_REFRESH_TOKEN } from './routes';

import { getAccessToken, getRefreshToken, setAuthTokens, removeTokens } from '../storage/localstorage';

const authClient = axios.create({
  baseURL: config.baseUrl,
});

const unauthClient = axios.create({
  baseURL: config.baseUrl,
});

authClient.interceptors.request.use((request) => {
  const accessToken = getAccessToken()
  if (accessToken != null) {
    request.headers['Authorization'] = `Bearer ${accessToken}`;
  }
  return request;
})

const refreshTokenRetrier = failedRequest => unauthClient.post(
    `${config.baseUrl}${API_REFRESH_TOKEN}`, { token: getRefreshToken() }, { validateStatus: (status) => status === 200 }
  ).then(tokenRefreshResponse => {
    setAuthTokens(tokenRefreshResponse.data.access_token, tokenRefreshResponse.data.refresh_token)
    failedRequest.response.config.headers['Authorization'] = `Bearer ${tokenRefreshResponse.data.access_token}`;
    return Promise.resolve();
  }).catch((err) => {
    console.log(err);
    removeTokens()
    // TODO: window.location.assign(routes.login);
  });

createAuthRefreshInterceptor(authClient, refreshTokenRetrier);

export const authorizedClient = authClient
export const unauthorizedClient = unauthClient