import { Form, Input } from 'antd';
import { Word, WordFormData } from '../types/word';

interface WordFormProps {
  form: any;
  initialValues?: Partial<Word>;
}

export const WordForm = ({ form, initialValues }: WordFormProps) => {
  return (
    <Form
      form={form}
      layout="vertical"
      initialValues={initialValues}
    >
      <Form.Item
        name="word"
        label="Từ vựng"
        rules={[{ required: true, message: 'Vui lòng nhập từ vựng' }]}
      >
        <Input />
      </Form.Item>
      <Form.Item
        name="meaning"
        label="Nghĩa"
        rules={[{ required: true, message: 'Vui lòng nhập nghĩa của từ' }]}
      >
        <Input />
      </Form.Item>
      <Form.Item
        name="pronunciation"
        label="Phát âm"
      >
        <Input />
      </Form.Item>
      <Form.Item
        name="example"
        label="Ví dụ"
      >
        <Input.TextArea rows={4} />
      </Form.Item>
    </Form>
  );
};
