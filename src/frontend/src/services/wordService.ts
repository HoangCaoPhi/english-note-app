import axios, { AxiosResponse } from 'axios';
import { API_ENDPOINTS } from '../constants/api';
import { Word, WordFormData } from '@/types/word';

interface ApiResponse<T> {
  data: T;
  message: string;
  status: number;
}

const getAuthHeader = () => {
  const token = localStorage.getItem('access_token');
  if (!token) {
    throw new Error('No authentication token found');
  }
  return {
    Authorization: `Bearer ${token}`
  };
};

export const wordService = {
  getWords: async (groupId: string): Promise<Word[]> => {
    try {
      const response: AxiosResponse<ApiResponse<Word[]>> = await axios.get(
        `${API_ENDPOINTS.WORDS}?groupId=${groupId}`,
        { headers: getAuthHeader() }
      );
      return response.data.data;
    } catch (error) {
      if (axios.isAxiosError(error)) {
        throw new Error(error.response?.data?.message || 'Failed to fetch words');
      }
      throw error;
    }
  },

  createWord: async (data: WordFormData & { groupId: string }): Promise<Word> => {
    try {
      const response: AxiosResponse<ApiResponse<Word>> = await axios.post(
        API_ENDPOINTS.WORDS,
        data,
        { headers: getAuthHeader() }
      );
      return response.data.data;
    } catch (error) {
      if (axios.isAxiosError(error)) {
        throw new Error(error.response?.data?.message || 'Failed to create word');
      }
      throw error;
    }
  },

  updateWord: async (id: string, data: Partial<Word>): Promise<Word> => {
    try {
      const response: AxiosResponse<ApiResponse<Word>> = await axios.put(
        `${API_ENDPOINTS.WORDS}/${id}`,
        data,
        { headers: getAuthHeader() }
      );
      return response.data.data;
    } catch (error) {
      if (axios.isAxiosError(error)) {
        throw new Error(error.response?.data?.message || 'Failed to update word');
      }
      throw error;
    }
  },

  deleteWord: async (id: string): Promise<void> => {
    try {
      await axios.delete(
        `${API_ENDPOINTS.WORDS}/${id}`,
        { headers: getAuthHeader() }
      );
    } catch (error) {
      if (axios.isAxiosError(error)) {
        throw new Error(error.response?.data?.message || 'Failed to delete word');
      }
      throw error;
    }
  }
};

