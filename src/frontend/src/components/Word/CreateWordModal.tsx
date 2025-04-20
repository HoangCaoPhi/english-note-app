import React from 'react';
import { Modal, Form, Input, Button, Space, Select } from 'antd';
import { MinusCircleOutlined, PlusOutlined } from '@ant-design/icons';

interface CreateWordModalProps {
  open: boolean;
  onCancel: () => void;
  onSubmit: (values: any) => void;
}

const CreateWordModal: React.FC<CreateWordModalProps> = ({
  open,
  onCancel,
  onSubmit,
}) => {
  const [form] = Form.useForm();

  const handleSubmit = () => {
    form.submit();
  };

  const initialValues = {
    word: '',
    pronunciations: [{ ipa: '', region: 'US' }],
    meanings: [{ 
      definition: '',
      part_of_speech: '',
      examples: ['']
    }]
  };

  const partOfSpeechOptions = [
    { value: 'noun', label: 'Noun' },
    { value: 'verb', label: 'Verb' },
    { value: 'adjective', label: 'Adjective' },
    { value: 'adverb', label: 'Adverb' },
    { value: 'pronoun', label: 'Pronoun' },
    { value: 'preposition', label: 'Preposition' },
    { value: 'conjunction', label: 'Conjunction' },
    { value: 'interjection', label: 'Interjection' }
  ];

  return (
    <Modal
      title="Add New Word"
      open={open}
      onCancel={onCancel}
      width="90vw"
      style={{ maxWidth: 1200, top: 20 }}
      footer={[
        <Button key="cancel" onClick={onCancel}>
          Cancel
        </Button>,
        <Button key="submit" type="primary" onClick={handleSubmit}>
          Save Word
        </Button>,
      ]}
      maskClosable={false}
    >
      <Form
        form={form}
        layout="vertical"
        initialValues={initialValues}
        onFinish={(values) => {
          onSubmit(values);
          form.resetFields();
        }}
        className="overflow-y-auto pr-2"
        style={{ maxHeight: 'calc(100vh - 200px)' }}
      >
        <div className="flex gap-4 items-start">
          <div className="flex-1">
            <Form.Item
              name="word"
              label="Word"
              rules={[{ required: true, message: 'Please input the word!' }]}
            >
              <Input placeholder="Enter word" size="large" />
            </Form.Item>

            <Form.List
              name="pronunciations"
              initialValue={[{ ipa: '', region: 'US' }]}
            >
              {(fields, { add, remove }) => (
                <>
                  {fields.map(({ key, name, ...restField }) => (
                    <Space key={key} style={{ display: 'flex', marginBottom: 8 }} align="baseline">
                      <Form.Item
                        {...restField}
                        name={[name, 'ipa']}
                        rules={[{ required: true, message: 'Missing IPA' }]}
                      >
                        <Input placeholder="IPA pronunciation" />
                      </Form.Item>
                      <Form.Item
                        {...restField}
                        name={[name, 'region']}
                        rules={[{ required: true, message: 'Missing region' }]}
                      >
                        <Input placeholder="Region (e.g., US, UK)" />
                      </Form.Item>
                      {fields.length > 1 && (
                        <MinusCircleOutlined onClick={() => remove(name)} />
                      )}
                    </Space>
                  ))}
                  <Form.Item>
                    <Button type="dashed" onClick={() => add()} block icon={<PlusOutlined />}>
                      Add Pronunciation
                    </Button>
                  </Form.Item>
                </>
              )}
            </Form.List>

            <Form.List
              name="meanings"
              initialValue={[{ 
                definition: '',
                part_of_speech: '',
                examples: ['']
              }]}
            >
              {(fields, { add, remove }) => (
                <>
                  {fields.map(({ key, name, ...restField }) => (
                    <div key={key} style={{ marginBottom: 16, padding: 16, border: '1px solid #f0f0f0', borderRadius: 8 }}>
                      <div className="flex gap-4">
                        <Form.Item
                          {...restField}
                          name={[name, 'definition']}
                          label="Definition"
                          className="flex-1"
                          rules={[{ required: true, message: 'Definition is required' }]}
                        >
                          <Input.TextArea 
                            placeholder="Enter definition" 
                            autoSize={{ minRows: 2 }}
                            size="large"
                          />
                        </Form.Item>
                        
                        <Form.Item
                          {...restField}
                          name={[name, 'part_of_speech']}
                          label="Part of Speech"
                          style={{ width: 200 }}
                        >
                          <Select
                            placeholder="Select part of speech"
                            options={partOfSpeechOptions}
                            allowClear
                          />
                        </Form.Item>
                      </div>
                      
                      <Form.List name={[name, 'examples']}>
                        {(exampleFields, { add: addExample, remove: removeExample }) => (
                          <>
                            {exampleFields.map((exampleField, index) => (
                              <Space key={exampleField.key} style={{ display: 'flex', marginBottom: 8 }}>
                                <Form.Item
                                  {...exampleField}
                                  validateTrigger={['onChange', 'onBlur']}
                                  style={{ flex: 1 }}
                                >
                                  <Input placeholder="Example sentence" />
                                </Form.Item>
                                <MinusCircleOutlined onClick={() => removeExample(index)} />
                              </Space>
                            ))}
                            <Button type="dashed" onClick={() => addExample()} block icon={<PlusOutlined />}>
                              Add Example
                            </Button>
                          </>
                        )}
                      </Form.List>
                      
                      {fields.length > 1 && (
                        <Button 
                          type="text" 
                          danger 
                          onClick={() => remove(name)} 
                          block
                          style={{ marginTop: 8 }}
                        >
                          Delete this meaning
                        </Button>
                      )}
                    </div>
                  ))}
                  <Form.Item>
                    <Button type="dashed" onClick={() => add()} block icon={<PlusOutlined />}>
                      Add Meaning
                    </Button>
                  </Form.Item>
                </>
              )}
            </Form.List>

            <div className="flex gap-4">
              <Form.List name="synonyms">
                {(fields, { add, remove }) => (
                  <div className="flex-1">
                    {fields.map(({ key, name, ...restField }) => (
                      <Space key={key} style={{ display: 'flex', marginBottom: 8 }} align="baseline">
                        <Form.Item
                          {...restField}
                          name={name}
                        >
                          <Input placeholder="Synonym" />
                        </Form.Item>
                        <MinusCircleOutlined onClick={() => remove(name)} />
                      </Space>
                    ))}
                    <Form.Item>
                      <Button type="dashed" onClick={() => add()} block icon={<PlusOutlined />}>
                        Add Synonym
                      </Button>
                    </Form.Item>
                  </div>
                )}
              </Form.List>

              <Form.List name="antonyms">
                {(fields, { add, remove }) => (
                  <div className="flex-1">
                    {fields.map(({ key, name, ...restField }) => (
                      <Space key={key} style={{ display: 'flex', marginBottom: 8 }} align="baseline">
                        <Form.Item
                          {...restField}
                          name={name}
                        >
                          <Input placeholder="Antonym" />
                        </Form.Item>
                        <MinusCircleOutlined onClick={() => remove(name)} />
                      </Space>
                    ))}
                    <Form.Item>
                      <Button type="dashed" onClick={() => add()} block icon={<PlusOutlined />}>
                        Add Antonym
                      </Button>
                    </Form.Item>
                  </div>
                )}
              </Form.List>
            </div>
          </div>
        </div>
      </Form>
    </Modal>
  );
};

export default CreateWordModal;




