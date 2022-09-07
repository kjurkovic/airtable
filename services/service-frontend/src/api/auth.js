import { unauthorizedClient} from './client';
import { setAuthTokens } from '../storage/localstorage';
import { API_FORGOT_PASSWORD, API_LOGIN, API_REGISTER } from './routes';

export const Auth = {
  login: (email, password) => {
    return unauthorizedClient.post(API_LOGIN, {
      email,
      password,
    }, {
      validateStatus: (status) => status === 200,
    }).then((response) => {
      setAuthTokens(response.data.accessToken, response.data.refreshToken);
      return response;
    });
  },
  register: (firstName, lastName, email, password) => {
    return unauthorizedClient.post(API_REGISTER, {
      firstName,
      lastName,
      email,
      password,
    }, {
      validateStatus: (status) => status === 200,
    }).then((response) => {
      setAuthTokens(response.data.accessToken, response.data.refreshToken);
      return response;
    });
  },
  forgotPassword: (email) => {
    return unauthorizedClient.post(API_FORGOT_PASSWORD, {
      email:email,
    }, {
      validateStatus: (status) => status === 204,
    });
  }
};
