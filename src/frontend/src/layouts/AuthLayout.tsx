import { useEffect } from 'react';
import { useNavigate, Outlet } from 'react-router-dom';
import axios from 'axios';
import { message } from 'antd';

const AuthLayout = () => {
  const navigate = useNavigate();

  const refreshToken = async () => {
    try {
      const refresh_token = localStorage.getItem('refresh_token');
      if (!refresh_token) {
        throw new Error('No refresh token');
      }

      const response = await axios.post('http://localhost:8080/auth/refresh', {
        refresh_token
      });

      localStorage.setItem('access_token', response.data.access_token);
      return response.data.access_token;
    } catch (error) {
      localStorage.removeItem('access_token');
      localStorage.removeItem('refresh_token');
      message.error('Phiên đăng nhập hết hạn');
      navigate('/login');
      throw error;
    }
  };

  useEffect(() => {
    // Add axios interceptor for handling token refresh
    const interceptor = axios.interceptors.response.use(
      response => response,
      async error => {
        const originalRequest = error.config;

        if (error.response?.status === 401 && !originalRequest._retry) {
          originalRequest._retry = true;
          try {
            const access_token = await refreshToken();
            originalRequest.headers['Authorization'] = `Bearer ${access_token}`;
            return axios(originalRequest);
          } catch (error) {
            return Promise.reject(error);
          }
        }
        return Promise.reject(error);
      }
    );

    return () => {
      axios.interceptors.response.eject(interceptor);
    };
  }, []);

  return <Outlet />;
};

export default AuthLayout;