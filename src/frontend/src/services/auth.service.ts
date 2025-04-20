import { BaseService } from './base.service';
import type { LoginRequest, RegisterRequest, AuthResponse } from '../types/auth';
import type { ApiResponse } from '../types/api';

export class AuthService extends BaseService {
  private static instance: AuthService;

  private constructor() {
    super();
  }

  public static getInstance(): AuthService {
    if (!AuthService.instance) {
      AuthService.instance = new AuthService();
    }
    return AuthService.instance;
  }

  public async login(data: LoginRequest): Promise<ApiResponse<AuthResponse['data']>> {
    return this.post<AuthResponse['data']>('/users/login', data);
  }

  public async register(data: RegisterRequest): Promise<ApiResponse<null>> {
    return this.post<null>('/users/register', data);
  }

  public async logout(): Promise<ApiResponse<null>> {
    return this.post<null>('/users/logout', {});
  }
}

export const authService = AuthService.getInstance();
