const TOKEN_KEY = {
  ACCESS_TOKEN: 'access_token',
  REFRESH_TOKEN: 'refresh_token',
};

export const TokenService = {
  getAccessToken: (): string | null => {
    return localStorage.getItem(TOKEN_KEY.ACCESS_TOKEN);
  },

  getRefreshToken: (): string | null => {
    return localStorage.getItem(TOKEN_KEY.REFRESH_TOKEN);
  },

  setTokens: (access_token: string, refresh_token: string): void => {
    localStorage.setItem(TOKEN_KEY.ACCESS_TOKEN, access_token);
    localStorage.setItem(TOKEN_KEY.REFRESH_TOKEN, refresh_token);
  },

  removeTokens: (): void => {
    localStorage.removeItem(TOKEN_KEY.ACCESS_TOKEN);
    localStorage.removeItem(TOKEN_KEY.REFRESH_TOKEN);
  },

  hasTokens: (): boolean => {
    return !!TokenService.getAccessToken() && !!TokenService.getRefreshToken();
  },
};