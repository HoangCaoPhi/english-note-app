import { PlusOutlined, BookOutlined } from "@ant-design/icons";
import { Button, Card, Col, Row, message, Statistic, Typography } from "antd";
import axios from "axios";
import { useEffect, useState } from "react";
import { useNavigate } from "react-router-dom";
import '@styles/dashboard.css';

const { Title } = Typography;

interface WordGroup {
  _id: string;
  name: string;
  description?: string;
  totalWords: number;
  createdAt: number;
  lastStudied?: number;
}

export default function Dashboard() {
  const [wordGroups, setWordGroups] = useState<WordGroup[]>([]);
  const [loading, setLoading] = useState(false);
  const [stats, setStats] = useState({
    totalGroups: 0,
    totalWords: 0,
    studiedToday: 0
  });
  const navigate = useNavigate();

  useEffect(() => {
    fetchWordGroups();
    fetchStats();
  }, []);

  const fetchWordGroups = async () => {
    try {
      setLoading(true);
      const response = await axios.get(
        "http://localhost:8080/word-groups",
        {
          headers: {
            Authorization: `Bearer ${localStorage.getItem("access_token")}`,
          },
        }
      );
      setWordGroups(response.data);
    } catch (error) {
      message.error("Không thể lấy dữ liệu nhóm từ vựng");
    } finally {
      setLoading(false);
    }
  };

  const fetchStats = async () => {
    try {
      const response = await axios.get(
        "http://localhost:8080/stats",
        {
          headers: {
            Authorization: `Bearer ${localStorage.getItem("access_token")}`,
          },
        }
      );
      setStats(response.data);
    } catch (error) {
      console.error("Không thể lấy thống kê:", error);
    }
  };

  const handleCreateGroup = () => {
    navigate("/word-groups/create");
  };

  return (
    <div className="dashboard-container">
      <div className="dashboard-header">
        <Title level={2}>Bảng điều khiển</Title>
        <Button
          type="primary"
          icon={<PlusOutlined />}
          onClick={handleCreateGroup}
          size="large"
        >
          Tạo nhóm từ vựng
        </Button>
      </div>

      <div className="stats-section">
        <Row gutter={16}>
          <Col span={8}>
            <Card>
              <Statistic
                title="Tổng số nhóm từ"
                value={stats.totalGroups}
                prefix={<BookOutlined />}
              />
            </Card>
          </Col>
          <Col span={8}>
            <Card>
              <Statistic
                title="Tổng số từ vựng"
                value={stats.totalWords}
                prefix={<BookOutlined />}
              />
            </Card>
          </Col>
          <Col span={8}>
            <Card>
              <Statistic
                title="Từ đã học hôm nay"
                value={stats.studiedToday}
                prefix={<BookOutlined />}
              />
            </Card>
          </Col>
        </Row>
      </div>

      <div className="word-groups-section">
        <Title level={3}>Nhóm từ vựng của bạn</Title>
        <Row gutter={[16, 16]}>
          {wordGroups.map((group) => (
            <Col xs={24} sm={12} md={8} lg={6} key={group._id}>
              <Card
                hoverable
                className="group-card"
                onClick={() => navigate(`/word-groups/${group._id}`)}
              >
                <div className="group-card-content">
                  <BookOutlined className="group-icon" />
                  <Title level={4}>{group.name}</Title>
                  {group.description && <p>{group.description}</p>}
                  <div className="group-stats">
                    <span>{group.totalWords} từ</span>
                    <span>•</span>
                    <span>
                      {new Date(group.createdAt).toLocaleDateString()}
                    </span>
                  </div>
                  {group.lastStudied && (
                    <div className="last-studied">
                      Học lần cuối: {new Date(group.lastStudied).toLocaleDateString()}
                    </div>
                  )}
                </div>
              </Card>
            </Col>
          ))}
        </Row>
      </div>
    </div>
  );
}
