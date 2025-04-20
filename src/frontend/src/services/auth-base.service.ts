import type { ApiResponse } from '../types/api';
import { TokenService } from './token.service';
import { BaseService } from './base.service';

export abstract class AuthBaseService extends BaseService {
  protected override async request<T>(
    endpoint: string,
    options: RequestInit = {}
  ): Promise<ApiResponse<T>> {
    const accessToken = TokenService.getAccessToken();
    
    if (!accessToken) {
      throw new Error('No access token available');
    }

    const headers = {
      Authorization: `Bearer ${accessToken}`,
      ...options.headers,
    };

    return super.request<T>(endpoint, {
      ...options,
      headers,
    });
  }
}