import React, { useState } from "react";
import { Form, Input, DatePicker, Upload, Button, Row, Col } from "antd";
import { UploadOutlined } from "@ant-design/icons";
import Description from "./components/Description";

const { RangePicker } = DatePicker;

/**
 * CreateAssignment component for creating a new assignment
 *
 * This component renders a form for creating a new assignment. It includes
 * fields for the assignment title, description, date range, and file uploads.
 *
 * @component
 * @returns {JSX.Element} Rendered CreateAssignment component
 */
const CreateAssignment: React.FC = () => {
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
            label="Assignment Title"
            rules={[{ required: true, message: "Please input the title!" }]}
          >
            <Input />
          </Form.Item>

          <Form.Item
            name="description"
            label="Assignment Description"
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
            name="dateRange"
            label="Start and End Date/Time"
            rules={[
              { required: true, message: "Please select the date/time range!" },
            ]}
          >
            <RangePicker
              showTime={{ format: "HH:mm" }}
              format="YYYY-MM-DD HH:mm"
            />
          </Form.Item>
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
          Create Assignment
        </Button>
      </Form.Item>
    </Form>
  );
};

export default CreateAssignment;
