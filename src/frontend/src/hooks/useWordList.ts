import { useState, useCallback } from 'react';
import { message } from 'antd';
import { Word, WordFormData } from '../types/word';
import { wordService } from '../services/wordService';

export const useWordList = (groupId: string) => {
  // States
  const [words, setWords] = useState<Word[]>([]);
  const [loading, setLoading] = useState(false);
  const [submitting, setSubmitting] = useState(false);
  const [selectedWord, setSelectedWord] = useState<Word | null>(null);
  const [searchText, setSearchText] = useState('');

  // Fetch words
  const fetchWords = useCallback(async () => {
    setLoading(true);
    try {
      const response = await wordService.getWords(groupId);
      setWords(response);
    } catch (error) {
      message.error('Không thể tải danh sách từ vựng');
    } finally {
      setLoading(false);
    }
  }, [groupId]);

  // CRUD operations
  const createWord = async (data: WordFormData) => {
    setSubmitting(true);
    try {
      await wordService.createWord({ ...data, groupId });
      message.success('Thêm từ vựng thành công');
      await fetchWords();
      return true;
    } catch (error) {
      message.error('Không thể thêm từ vựng');
      return false;
    } finally {
      setSubmitting(false);
    }
  };

  const updateWord = async (id: string, data: WordFormData) => {
    setSubmitting(true);
    try {
      await wordService.updateWord(id, data);
      message.success('Cập nhật từ vựng thành công');
      await fetchWords();
      return true;
    } catch (error) {
      message.error('Không thể cập nhật từ vựng');
      return false;
    } finally {
      setSubmitting(false);
    }
  };

  const deleteWord = async (id: string) => {
    try {
      await wordService.deleteWord(id);
      message.success('Xóa từ vựng thành công');
      await fetchWords();
      return true;
    } catch (error) {
      message.error('Không thể xóa từ vựng');
      return false;
    }
  };

  // Filter words based on search
  const filteredWords = words.filter(word => 
    word.word.toLowerCase().includes(searchText.toLowerCase()) ||
    word.meaning.toLowerCase().includes(searchText.toLowerCase())
  );

  return {
    words: filteredWords,
    loading,
    submitting,
    selectedWord,
    searchText,
    setSearchText,
    setSelectedWord,
    fetchWords,
    createWord,
    updateWord,
    deleteWord
  };
};

