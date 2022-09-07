export const LS_ACCESS_TOKEN = "thAT"
export const LS_REFRESH_TOKEN = "thRT"

export const setAuthTokens = (accessToken, refreshToken) => {
  localStorage.setItem(LS_ACCESS_TOKEN, accessToken);
  localStorage.setItem(LS_REFRESH_TOKEN, refreshToken);
};

export const getAccessToken = () => {
  return localStorage.getItem(LS_ACCESS_TOKEN);
};

// TODO: move this to session cookie or somewhere else - this is not secure
export const getRefreshToken = () => {
  return localStorage.getItem(LS_REFRESH_TOKEN);
};

export const removeTokens = () => {
  localStorage.removeItem(LS_ACCESS_TOKEN);
  localStorage.removeItem(LS_REFRESH_TOKEN);
};
