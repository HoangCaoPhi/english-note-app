import { useState } from "react";
import axios from "axios";
import { Input, Button, Form, message } from "antd";
import { UserOutlined, LockOutlined, MailOutlined } from "@ant-design/icons";
import { Link, useNavigate } from "react-router-dom";
import "@styles/register.css";

const Register = () => {
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [username, setUsername] = useState("");
  const navigate = useNavigate();

  const handleSubmit = async (values: {
    email: string;
    password: string;
    username: string;
  }) => {
    try {
      const response = await axios.post(
        "http://localhost:8080/users/register",
        {
          email: values.email,
          password: values.password,
          username: values.username,
        }
      );

      message.success("Đăng ký thành công");
      navigate("/login");
    } catch (err) {
      message.error("Đã có lỗi xảy ra, vui lòng thử lại!");
    }
  };

  return (
    <div className="register-container">
      <h1>Đăng ký</h1>
      <Form onFinish={handleSubmit} layout="vertical">
        <Form.Item
          label="Tên người dùng"
          name="username"
          rules={[{ required: true, message: "Vui lòng nhập tên người dùng!" }]}
        >
          <Input
            prefix={<UserOutlined />}
            placeholder="Nhập tên người dùng"
            value={username}
            onChange={(e) => setUsername(e.target.value)}
          />
        </Form.Item>

        <Form.Item
          label="Email"
          name="email"
          rules={[
            {
              required: true,
              message: "Vui lòng nhập email của bạn!",
              type: "email",
            },
          ]}
        >
          <Input
            prefix={<MailOutlined />}
            placeholder="Nhập email"
            value={email}
            onChange={(e) => setEmail(e.target.value)}
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
            Đăng ký
          </Button>
        </Form.Item>
      </Form>
      <div style={{ textAlign: "center" }}>
        Đã có tài khoản? <Link to="/login">Đăng nhập ngay</Link>
      </div>
    </div>
  );
};

export default Register;
