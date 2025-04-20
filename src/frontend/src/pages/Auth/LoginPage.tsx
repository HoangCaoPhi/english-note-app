import React from 'react';
import { Typography } from 'antd';
import LoginForm from '../../components/Auth/LoginForm';
import { Link } from 'react-router-dom';

const { Title } = Typography;

const LoginPage: React.FC = () => {
  return (
    <div className="min-h-screen flex flex-col items-center justify-center bg-gray-50 py-12 px-4 sm:px-6 lg:px-8">
      <div className="max-w-md w-full space-y-8">
        <div className="text-center">
          <Title level={2}>Login to your account</Title>
          <p className="mt-2 text-gray-600">
            Don't have an account? <Link to="/register">Register</Link>
          </p>
        </div>
        <LoginForm />
      </div>
    </div>
  );
};

export default LoginPage;