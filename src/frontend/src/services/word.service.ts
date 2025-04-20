import { AuthBaseService } from './auth-base.service';
import type { ApiResponse } from '../types/api';

interface Pronunciation {
  ipa: string;
  region: string;
}

interface Meaning {
  part_of_speech: string;
  definition: string;
  examples: string[];
}

export interface Word {
  id: string;
  groupId: string;
  userId: string;
  word: string;
  language: string;
  pronunciations: Pronunciation[];
  meanings: Meaning[];
  synonyms: string[];
  antonyms: string[];
  createdAt: number;
}

interface WordsResponse {
  data: Word[];
  pagination: {
    page: number;
    limit: number;
    total: number;
  };
}

interface CreateWordRequest {
  groupId: string;
  word: string;
  language?: string;
  pronunciations: Pronunciation[];
  meanings?: Meaning[];
  synonyms?: string[];
  antonyms?: string[];
}

interface CreateWordResponse {
  id: string;
}

export class WordService extends AuthBaseService {
  private static instance: WordService;

  private constructor() {
    super();
  }

  public static getInstance(): WordService {
    if (!WordService.instance) {
      WordService.instance = new WordService();
    }
    return WordService.instance;
  }

  public async getWords(groupId: string, page: number = 1, limit: number = 10): Promise<ApiResponse<WordsResponse>> {
    return this.get<WordsResponse>(`/words?groupId=${groupId}&page=${page}&limit=${limit}`);
  }

  public async createWord(data: CreateWordRequest): Promise<ApiResponse<CreateWordResponse>> {
    return this.post<CreateWordResponse>('/words', data);
  }
}

export const wordService = WordService.getInstance();