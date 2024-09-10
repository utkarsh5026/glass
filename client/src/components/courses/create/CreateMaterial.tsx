import React, { useState } from "react";
import { Form, Input, Upload, Button, Row, Col } from "antd";
import { UploadOutlined } from "@ant-design/icons";
import Description from "./components/Description";

/**
 * CreateMaterial component for creating new course material
 *
 * This component renders a form for creating new course material. It includes
 * fields for the material title, description, and file uploads.
 *
 * @component
 * @returns {JSX.Element} Rendered CreateMaterial component
 */
const CreateMaterial: React.FC = () => {
  const [form] = Form.useForm();
  const [markdown, setMarkdown] = useState<string>("");

  /**
   * Handles form submission
   *
   * @param {Object} values - The form values
   */
  const onFinish = (values: any) => {
    console.log("Form values:", values);
    // Handle form submission
  };

  return (
    <Form form={form} layout="vertical" onFinish={onFinish}>
      <Row gutter={[16, 16]}>
        <Col span={16}>
          <Form.Item
            name="title"
            label="Material Title"
            rules={[{ required: true, message: "Please input the title!" }]}
          >
            <Input />
          </Form.Item>

          <Form.Item
            name="description"
            label="Material Description"
            rules={[
              { required: true, message: "Please input the description!" },
            ]}
          >
            <Description
              markdown={markdown}
              onChange={setMarkdown}
              editorRef={null}
            />
          </Form.Item>
        </Col>
        <Col span={6}>
          <Form.Item
            name="files"
            label="Add Files"
            valuePropName="fileList"
            getValueFromEvent={(e) => {
              if (Array.isArray(e)) return e;
              return e && e.fileList;
            }}
          >
            <Upload>
              <Button icon={<UploadOutlined />}>Add Files</Button>
            </Upload>
          </Form.Item>
        </Col>
      </Row>
      <Form.Item>
        <Button type="primary" htmlType="submit">
          Create Material
        </Button>
      </Form.Item>
    </Form>
  );
};

export default CreateMaterial;
