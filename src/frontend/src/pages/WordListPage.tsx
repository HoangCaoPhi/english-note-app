import React, { useEffect, useState } from 'react';
import { Layout, Table, Button, message, Input, Space } from 'antd';
import { PlusOutlined } from '@ant-design/icons';
import { useParams } from 'react-router-dom';
import type { Word } from '../services/word.service';
import { wordService } from '../services/word.service'; 
import CreateWordModal from '../components/Word/CreateWordModal';

const { Content } = Layout;
const { Search } = Input;

const WordListPage: React.FC = () => {
  const { groupId } = useParams<{ groupId: string }>();
  const [words, setWords] = useState<Word[]>([]);
  const [loading, setLoading] = useState(false);
  const [createModalOpen, setCreateModalOpen] = useState(false);
  const [pagination, setPagination] = useState({
    current: 1,
    pageSize: 10,
    total: 0,
  });
  const [searchText, setSearchText] = useState('');

  const fetchWords = async (page: number = 1) => {
    if (!groupId) return;
    
    try {
      setLoading(true);
      const response = await wordService.getWords(groupId, page, pagination.pageSize);
      if (response.data) {
        setWords(response.data.data);
        setPagination({
          ...pagination,
          current: page,
          total: response.data.pagination.total,
        });
      }
    } catch (error) {
      message.error('Failed to fetch words');
    } finally {
      setLoading(false);
    }
  };

  const handleTableChange = (pagination: any) => {
    fetchWords(pagination.current);
  };

  const handleCreateWord = async (values: any) => {
    try {
      await wordService.createWord({
        ...values,
        groupId,
      });
      message.success('Word created successfully');
      setCreateModalOpen(false);
      fetchWords();
    } catch (error) {
      message.error('Failed to create word');
    }
  };

  const columns = [
    {
      title: 'Word',
      dataIndex: 'word',
      key: 'word',
      render: (text: string, record: Word) => (
        <div>
          <div className="text-lg font-bold">{text}</div>
          <div className="text-gray-500">
            {record.pronunciations.map((p, i) => (
              <span key={i}>{p.ipa} ({p.region})</span>
            ))}
          </div>
        </div>
      ),
    },
    {
      title: 'Meanings',
      dataIndex: 'meanings',
      key: 'meanings',
      render: (meanings: Word['meanings']) => (
        <div>
          {meanings.map((m, i) => (
            <div key={i} className="mb-2">
              <div className="italic text-gray-600">{m.part_of_speech}</div>
              <div>{m.definition}</div>
              <ul className="list-disc ml-4">
                {m.examples.map((ex, j) => (
                  <li key={j} className="text-gray-600">{ex}</li>
                ))}
              </ul>
            </div>
          ))}
        </div>
      ),
    },
    {
      title: 'Synonyms/Antonyms',
      key: 'synonyms_antonyms',
      render: (record: Word) => (
        <div>
          {record.synonyms?.length > 0 && (
            <div className="mb-2">
              <span className="font-bold">Synonyms: </span>
              {record.synonyms.join(', ')}
            </div>
          )}
          {record.antonyms?.length > 0 && (
            <div>
              <span className="font-bold">Antonyms: </span>
              {record.antonyms.join(', ')}
            </div>
          )}
        </div>
      ),
    },
  ];

  useEffect(() => {
    fetchWords();
  }, [groupId]);

  return (
    <Content style={{ padding: '24px', background: '#f0f2f5' }}>
    <div style={{ 
      background: '#fff', 
      padding: '24px', 
      borderRadius: '8px',
      minHeight: 'calc(100vh - 112px)'
    }}>
      <div className="flex justify-between items-center mb-6">
        <h1 className="text-2xl font-bold">Words</h1>
        <Space>
          <Search
            placeholder="Search words"
            allowClear
            onChange={(e) => setSearchText(e.target.value)}
            style={{ width: 250 }}
          />
          <Button
            type="primary"
            icon={<PlusOutlined />}
            onClick={() => setCreateModalOpen(true)}
          >
            Add New Word
          </Button>
        </Space>
      </div>

      <Table
        columns={columns}
        dataSource={words}
        rowKey="id"
        pagination={pagination}
        onChange={handleTableChange}
        loading={loading}
      />

      <CreateWordModal
        open={createModalOpen}
        onCancel={() => setCreateModalOpen(false)}
        onSubmit={handleCreateWord}
      />
    </div>
  </Content>
  );
};

export default WordListPage;