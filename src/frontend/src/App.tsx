import { BrowserRouter, Routes, Route, Navigate } from 'react-router-dom';
import Login from '@pages/Login';
import Register from '@pages/auth/Register';
import AuthLayout from '@layouts/AuthLayout';
import MainLayout from '@layouts/MainLayout';
import Dashboard from '@pages/Dashboard';
import WordList from '@pages/WordList';

const App = () => {
  return (
    <BrowserRouter>
      <Routes>
        {/* Public routes */}
        <Route path="/login" element={<Login />} />
        <Route path="/register" element={<Register />} />

        {/* Protected routes */}
        <Route element={<AuthLayout />}>
          <Route element={<MainLayout />}>
            <Route path="/dashboard" element={<Dashboard />} />
            <Route path="/word-group/:groupId" element={<WordList />} />
          </Route>
        </Route>

        {/* Redirect to login if no route matches */}
        <Route path="*" element={<Navigate to="/login" replace />} />
      </Routes>
    </BrowserRouter>
  );
};

export default App;



