import React, { useEffect } from 'react';
import { Form, Input, Button, message } from 'antd';
import { useDispatch, useSelector } from 'react-redux';
import { login, clearError } from '../../store/slices/authSlice';
import type { LoginRequest } from '../../types/auth'; 
import { useNavigate } from 'react-router-dom';
import { AppDispatch, RootState } from '../../types/store';

const LoginForm: React.FC = () => {
  const dispatch = useDispatch<AppDispatch>();
  const navigate = useNavigate();
  const { loading, error } = useSelector((state: RootState) => state.auth);
  const [messageApi] = message.useMessage();

  useEffect(() => {
    if (error) {
      messageApi.error(error);
      dispatch(clearError());
    }
  }, [error, messageApi, dispatch]);

  const onFinish = async (values: LoginRequest) => {
    const result = await dispatch(login(values));
    if (login.fulfilled.match(result)) {
      messageApi.success('Login successful');
      navigate('/');
    }
  };

  return (
    <Form
      name="login"
      onFinish={onFinish}
      layout="vertical"
      className="max-w-md mx-auto mt-8"
    >
      <Form.Item
        label="Username"
        name="username"
        rules={[{ required: true, message: 'Please input your username!' }]}
      >
        <Input />
      </Form.Item>

      <Form.Item
        label="Password"
        name="password"
        rules={[{ required: true, message: 'Please input your password!' }]}
      >
        <Input.Password />
      </Form.Item>

      <Form.Item>
        <Button type="primary" htmlType="submit" loading={loading} block>
          Login
        </Button>
      </Form.Item>
    </Form>
  );
};

export default LoginForm;