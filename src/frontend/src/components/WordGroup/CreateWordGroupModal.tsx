import React from 'react';
import { Modal, Form, Input } from 'antd';

interface CreateWordGroupModalProps {
  open: boolean;
  loading: boolean;
  onCancel: () => void;
  onSubmit: (values: { name: string }) => void;
}

const CreateWordGroupModal: React.FC<CreateWordGroupModalProps> = ({
  open,
  onCancel,
  onSubmit,
  loading,
}) => {
  const [form] = Form.useForm();

  return (
    <Modal
      title="Create New Word Group"
      open={open}
      onCancel={onCancel}
      onOk={() => form.submit()}
      confirmLoading={loading}
    >
      <Form
        form={form}
        layout="vertical"
        onFinish={(values) => {
          onSubmit(values);
          form.resetFields();
        }}
      >
        <Form.Item
          name="name"
          label="Group Name"
          rules={[{ required: true, message: 'Please input the group name!' }]}
        >
          <Input placeholder="Enter group name" />
        </Form.Item>
      </Form>
    </Modal>
  );
};

export default CreateWordGroupModal;
