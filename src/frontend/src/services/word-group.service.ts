import { AuthBaseService } from './auth-base.service';
import type { ApiResponse } from '../types/api';

export interface WordGroup {
  id: string;
  userId: string;
  name: string;
  createdAt: number;
  // Thêm các trường default
  imageUrl: string;
  description: string;
  wordCount: number;
}

interface CreateWordGroupRequest {
  name: string;
}

interface CreateWordGroupResponse {
  id: string;
}

export class WordGroupService extends AuthBaseService {
  private static instance: WordGroupService;

  private constructor() {
    super();
  }

  public static getInstance(): WordGroupService {
    if (!WordGroupService.instance) {
      WordGroupService.instance = new WordGroupService();
    }
    return WordGroupService.instance;
  }

  public async getWordGroups(): Promise<ApiResponse<WordGroup[]>> {
    return this.get<WordGroup[]>('/word-groups');
  }

  public async createWordGroup(data: CreateWordGroupRequest): Promise<ApiResponse<CreateWordGroupResponse>> {
    return this.post<CreateWordGroupResponse>('/word-groups', data);
  }
}

export const wordGroupService = WordGroupService.getInstance();