import React from 'react';
import { useNavigate } from 'react-router-dom';
import { Button, message } from 'antd';
import { PlusOutlined } from '@ant-design/icons';
import WordGroupCard from '../components/WordGroup/WordGroupCard';
import CreateWordGroupModal from '../components/WordGroup/CreateWordGroupModal';
import { wordGroupService } from '../services/word-group.service';
import type { WordGroup } from '../services/word-group.service';

const DashboardPage: React.FC = () => {
  const navigate = useNavigate();
  const [isModalOpen, setIsModalOpen] = React.useState(false);
  const [loading, setLoading] = React.useState(false);
  const [wordGroups, setWordGroups] = React.useState<WordGroup[]>([]);

  const fetchWordGroups = async () => {
    try {
      setLoading(true);
      const response = await wordGroupService.getWordGroups();
      if (response.data) {
        setWordGroups(response.data);
      }
    } catch (error) {
      message.error('Failed to fetch word groups');
    } finally {
      setLoading(false);
    }
  };

  React.useEffect(() => {
    fetchWordGroups();
  }, []);

  const handleCreateWordGroup = async (values: { name: string }) => {
    try {
      setLoading(true);
      await wordGroupService.createWordGroup(values);
      message.success('Word group created successfully');
      setIsModalOpen(false);
      fetchWordGroups(); // Refresh the list
    } catch (error) {
      message.error('Failed to create word group');
    } finally {
      setLoading(false);
    }
  };

  return (
    <div className="space-y-6">
      <div className="flex justify-between items-center">
        <h1 className="text-2xl font-bold text-gray-800">My Word Groups</h1>
        <Button
          type="primary"
          icon={<PlusOutlined />}
          onClick={() => setIsModalOpen(true)}
          className="bg-blue-600 hover:bg-blue-700"
        >
          Create New Group
        </Button>
      </div>

      <div className="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 lg:grid-cols-4 gap-6">
        {wordGroups.map((group) => (
          <WordGroupCard
            id={group.id}
            key={group.id}
            name={group.name}
            description={group.description}
            imageUrl={group.imageUrl}
            wordCount={group.wordCount}
            onClick={() => navigate(`/groups/${group.id}`)}
            className="h-full"
          />
        ))}
      </div>

      <CreateWordGroupModal
        open={isModalOpen}
        onCancel={() => setIsModalOpen(false)}
        onSubmit={handleCreateWordGroup}
        loading={loading}
      />
    </div>
  );
};

export default DashboardPage;





