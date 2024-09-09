import React from "react";
import { Dropdown, MenuProps, Button } from "antd";
import {
  SoundOutlined,
  FileOutlined,
  EditOutlined,
  PlusOutlined,
  CarryOutFilled,
} from "@ant-design/icons";

const items: MenuProps["items"] = [
  {
    key: "1",
    label: "Announcement",
    icon: <SoundOutlined />,
  },
  {
    key: "2",
    label: "Assignment",
    icon: <CarryOutFilled />,
  },
  {
    key: "3",
    label: "Material",
    icon: <FileOutlined />,
  },
  {
    key: "4",
    label: "Quiz",
    icon: <EditOutlined />,
  },
];

const AddDropDown: React.FC = () => {
  return (
    <Dropdown menu={{ items }}>
      <Button icon={<PlusOutlined />} type="primary">
        Create Something 😊
      </Button>
    </Dropdown>
  );
};

export default AddDropDown;
