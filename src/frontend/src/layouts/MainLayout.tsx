import { Layout, Menu, Button, theme } from 'antd';
import { Outlet, useNavigate } from 'react-router-dom';
import { LogoutOutlined, BookOutlined, DashboardOutlined } from '@ant-design/icons';

const { Header, Content, Sider } = Layout;

const MainLayout = () => {
  const navigate = useNavigate();
  const {
    token: { colorBgContainer, borderRadiusLG },
  } = theme.useToken();

  const handleLogout = () => {
    localStorage.removeItem('access_token');
    localStorage.removeItem('refresh_token');
    navigate('/login');
  };

  return (
    <Layout style={{ minHeight: '100vh' }}>
      <Header style={{ padding: '0 16px', display: 'flex', alignItems: 'center', justifyContent: 'space-between' }}>
        <div style={{ color: 'white', fontSize: '20px' }}>
          English Note
        </div>
        <Button 
          type="text" 
          icon={<LogoutOutlined />} 
          onClick={handleLogout}
          style={{ color: 'white' }}
        >
          Đăng xuất
        </Button>
      </Header>
      <Layout>
        <Sider width={200} style={{ background: colorBgContainer }}>
          <Menu
            mode="inline"
            defaultSelectedKeys={['dashboard']}
            style={{ height: '100%', borderRight: 0 }}
            items={[
              {
                key: 'dashboard',
                icon: <DashboardOutlined />,
                label: 'Tổng quan',
                onClick: () => navigate('/dashboard')
              },
              {
                key: 'wordgroups',
                icon: <BookOutlined />,
                label: 'Nhóm từ vựng',
                onClick: () => navigate('/word-groups')
              }
            ]}
          />
        </Sider>
        <Layout style={{ padding: '24px' }}>
          <Content
            style={{
              padding: 24,
              margin: 0,
              background: colorBgContainer,
              borderRadius: borderRadiusLG,
              minHeight: 280,
            }}
          >
            <Outlet />
          </Content>
        </Layout>
      </Layout>
    </Layout>
  );
};

export default MainLayout;