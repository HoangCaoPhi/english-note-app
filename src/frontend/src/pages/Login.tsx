import { LockOutlined, UserOutlined } from "@ant-design/icons";
import { Button, Form, Input, message } from "antd";
import axios from "axios";
import { useState } from "react";
import { Link, useNavigate } from "react-router-dom";
import '@styles/login.css';

interface LoginFormValues {
  username: string;
  password: string;
}

const Login = () => {
  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");
  const navigate = useNavigate();

  const handleSubmit = async (values: LoginFormValues): Promise<void> => {
    try {
      const response = await axios.post(
        "http://localhost:8080/users/login",
        {
          username: values.username,
          password: values.password,
        }
      );

      localStorage.setItem("access_token", response.data.access_token);
      localStorage.setItem("refresh_token", response.data.refresh_token);

      message.success("Đăng nhập thành công");
      navigate("/dashboard");
    } catch (error) {
      message.error("Thông tin đăng nhập không chính xác");
    }
  };

  return (
    <div
      className="login-container"
      style={{ width: 400, margin: "0 auto", paddingTop: 50 }}
    >
      <h1>Đăng nhập</h1>
      <Form onFinish={handleSubmit} layout="vertical">
        <Form.Item
          label="Tên đăng nhập"
          name="username"
          rules={[
            {
              required: true,
              message: "Vui lòng nhập tên đăng nhập!",
            },
          ]}
        >
          <Input
            prefix={<UserOutlined />}
            placeholder="Nhập tên đăng nhập"
            value={username}
            onChange={(e) => setUsername(e.target.value)}
          />
        </Form.Item>

        <Form.Item
          label="Mật khẩu"
          name="password"
          rules={[
            { required: true, message: "Vui lòng nhập mật khẩu của bạn!" },
          ]}
        >
          <Input.Password
            prefix={<LockOutlined />}
            placeholder="Nhập mật khẩu"
            value={password}
            onChange={(e) => setPassword(e.target.value)}
          />
        </Form.Item>

        <Form.Item>
          <Button type="primary" htmlType="submit" block>
            Đăng nhập
          </Button>
        </Form.Item>
      </Form>
      <div style={{ textAlign: "center" }}>
        Chưa có tài khoản? <Link to="/register">Đăng ký ngay</Link>
      </div>
    </div>
  );
};

export default Login;
