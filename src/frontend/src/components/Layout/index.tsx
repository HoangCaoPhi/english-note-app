import React from 'react';
import { Layout } from 'antd';
import { Outlet } from 'react-router-dom';
import HeaderMenu from './HeaderMenu';

const { Content } = Layout;

const AppLayout: React.FC = () => {
  return (
    <Layout className="h-screen flex flex-col">
      <HeaderMenu />
      <Content className="flex-1 overflow-auto">
        <div className="container mx-auto px-4 py-6">
          <Outlet />
        </div>
      </Content>
    </Layout>
  );
};

export default AppLayout;


