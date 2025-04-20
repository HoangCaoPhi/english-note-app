import React, { useEffect } from 'react';
import { Form, Input, Button, message } from 'antd';
import { useDispatch, useSelector } from 'react-redux';
import { register, clearError } from '../../store/slices/authSlice';
import type { RegisterRequest } from '../../types/auth'; 
import { useNavigate } from 'react-router-dom';
import { AppDispatch, RootState } from '../../types/store';

const RegisterForm: React.FC = () => {
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

  const onFinish = async (values: RegisterRequest) => {
    const result = await dispatch(register(values));
    if (register.fulfilled.match(result)) {
      messageApi.success('Registration successful');
      navigate('/login');
    }
  };

  return (
    <Form
      name="register"
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
        label="Email"
        name="email"
        rules={[
          { required: true, message: 'Please input your email!' },
          { type: 'email', message: 'Please input a valid email!' }
        ]}
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
          Register
        </Button>
      </Form.Item>
    </Form>
  );
};

export default RegisterForm;