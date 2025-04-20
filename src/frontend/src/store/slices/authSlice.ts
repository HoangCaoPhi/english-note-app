import { createSlice, createAsyncThunk } from '@reduxjs/toolkit';
import type { AuthState, LoginRequest, RegisterRequest } from '../../types/auth';
import { authService } from '../../services/auth.service';
import { TokenService } from '../../services/token.service';

export const register = createAsyncThunk<void, RegisterRequest>(
  'auth/register',
  async (data, { rejectWithValue }) => {
    try {
      await authService.register(data);
    } catch (error) {
      return rejectWithValue(error instanceof Error ? error.message : 'Registration failed');
    }
  }
);

export const login = createAsyncThunk<
  { access_token: string; refresh_token: string },
  LoginRequest
>(
  'auth/login',
  async (data, { rejectWithValue }) => {
    try {
      const response = await authService.login(data);
      if (!response.data?.access_token || !response.data?.refresh_token) {
        throw new Error('Invalid response format');
      }
      
      // Store tokens in localStorage
      TokenService.setTokens(
        response.data.access_token,
        response.data.refresh_token
      );
      
      return {
        access_token: response.data.access_token,
        refresh_token: response.data.refresh_token,
      };
    } catch (error) {
      return rejectWithValue(error instanceof Error ? error.message : 'Login failed');
    }
  }
);

const initialState: AuthState = {
  isAuthenticated: TokenService.hasTokens(),
  accessToken: TokenService.getAccessToken(),
  refreshToken: TokenService.getRefreshToken(),
  loading: false,
  error: null,
};

const authSlice = createSlice({
  name: 'auth',
  initialState,
  reducers: {
    logout: (state) => {
      state.isAuthenticated = false;
      state.accessToken = null;
      state.refreshToken = null;
      state.error = null;
      TokenService.removeTokens();
    },
    clearError: (state) => {
      state.error = null;
    },
  },
  extraReducers: (builder) => {
    builder
      // Register
      .addCase(register.pending, (state) => {
        state.loading = true;
        state.error = null;
      })
      .addCase(register.fulfilled, (state) => {
        state.loading = false;
      })
      .addCase(register.rejected, (state, action) => {
        state.loading = false;
        state.error = action.payload as string;
      })
      // Login
      .addCase(login.pending, (state) => {
        state.loading = true;
        state.error = null;
      })
      .addCase(login.fulfilled, (state, action) => {
        state.loading = false;
        state.isAuthenticated = true;
        state.accessToken = action.payload.access_token;
        state.refreshToken = action.payload.refresh_token;
      })
      .addCase(login.rejected, (state, action) => {
        state.loading = false;
        state.error = action.payload as string;
        state.isAuthenticated = false;
        state.accessToken = null;
        state.refreshToken = null;
        TokenService.removeTokens();
      });
  },
});

export const { logout, clearError } = authSlice.actions;
export default authSlice.reducer;

