import { useEffect, useState } from 'react';
import { Table, Button, Modal, Form, Input, Space, Popconfirm, Typography } from 'antd';
import { PlusOutlined, EditOutlined, DeleteOutlined, SearchOutlined } from '@ant-design/icons';
import { useParams } from 'react-router-dom';
import { Word, WordFormData } from '../types/word';
import { useWordList } from '../hooks/useWordList';
import { WordForm } from '../components/WordForm';
import '@styles/WordList.css';

const { Title } = Typography;

const WordList = () => {
  const { groupId } = useParams<{ groupId: string }>();
  const [form] = Form.useForm();
  const [isModalVisible, setIsModalVisible] = useState(false);
  
  const {
    words,
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
  } = useWordList(groupId!);

  useEffect(() => {
    if (groupId) {
      fetchWords();
    }
  }, [groupId, fetchWords]);

  const handleSubmit = async (values: WordFormData) => {
    let success;
    if (selectedWord) {
      success = await updateWord(selectedWord._id, values);
    } else {
      success = await createWord(values);
    }
    if (success) {
      handleCloseModal();
    }
  };

  const handleEdit = (record: Word) => {
    setSelectedWord(record);
    form.setFieldsValue(record);
    setIsModalVisible(true);
  };

  const handleCloseModal = () => {
    setIsModalVisible(false);
    setSelectedWord(null);
    form.resetFields();
  };

  // Table columns configuration
  const columns = [
    {
      title: 'Từ vựng',
      dataIndex: 'word',
      sorter: (a: Word, b: Word) => a.word.localeCompare(b.word),
      width: '20%',
    },
    {
      title: 'Nghĩa',
      dataIndex: 'meaning',
      key: 'meaning',
      width: '25%',
    },
    {
      title: 'Phát âm',
      dataIndex: 'pronunciation',
      key: 'pronunciation',
      width: '15%',
    },
    {
      title: 'Ví dụ',
      dataIndex: 'example',
      key: 'example',
      width: '30%',
      ellipsis: true,
    },
    {
      title: 'Thao tác',
      key: 'action',
      width: '10%',
      render: (_: any, record: Word) => (
        <Space>
          <Button
            icon={<EditOutlined />}
            onClick={() => handleEdit(record)}
          />
          <Popconfirm
            title="Bạn có chắc chắn muốn xóa?"
            onConfirm={() => deleteWord(record._id)}
          >
            <Button icon={<DeleteOutlined />} danger />
          </Popconfirm>
        </Space>
      ),
    },
  ];

  return (
    <div className="word-list-container">
      {/* Header with search and add button */}
      <div className="word-list-header">
        <Title level={3}>Danh sách từ vựng</Title>
        <Space>
          <Input
            placeholder="Tìm kiếm từ vựng..."
            prefix={<SearchOutlined />}
            value={searchText}
            onChange={e => setSearchText(e.target.value)}
            allowClear
          />
          <Button
            type="primary"
            icon={<PlusOutlined />}
            onClick={() => setIsModalVisible(true)}
          >
            Thêm từ mới
          </Button>
        </Space>
      </div>

      {/* Table */}
      <Table
        className="word-list-table"
        columns={columns}
        dataSource={words}
        rowKey="_id"
        loading={loading}
        pagination={{
          showSizeChanger: true,
          showTotal: (total) => `Tổng số: ${total} từ`,
        }}
      />

      {/* Modal for create/edit */}
      <Modal
        title={selectedWord ? "Sửa từ vựng" : "Thêm từ vựng mới"}
        open={isModalVisible}
        onCancel={handleCloseModal}
        confirmLoading={submitting}
        onOk={form.submit}
      >
        <WordForm 
          form={form}
          initialValues={selectedWord || undefined}
        />
      </Modal>
    </div>
  );
};

export default WordList;




