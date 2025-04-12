export interface Word {
  _id: string;
  word: string;
  meaning: string;
  pronunciation: string;
  example: string;
  groupId: string;
  createdAt: string;
  updatedAt: string;
}

export interface WordFormData {
  word: string;
  meaning: string;
  pronunciation?: string;  // optional
  example?: string;       // optional
}

