import React from "react";
import { Layout, Avatar, Dropdown } from "antd";
import {
  UserOutlined,
  SettingOutlined,
  LogoutOutlined,
} from "@ant-design/icons";

const { Header } = Layout;

const AppBar: React.FC = () => {
  const menuItems = [
    {
      key: "settings",
      icon: <SettingOutlined />,
      label: "Settings",
    },
    {
      key: "logout",
      icon: <LogoutOutlined />,
      label: "Logout",
    },
  ];

  return (
    <Header
      style={{
        display: "flex",
        justifyContent: "space-between",
        alignItems: "center",
        padding: "0 16px",
      }}
    >
      <div className="logo">
        {/* Add your logo or icon here */}
        <img
          src="/path-to-your-logo.png"
          alt="Logo"
          style={{ height: "32px" }}
        />
      </div>
      <Dropdown menu={{ items: menuItems }} placement="bottomRight">
        <Avatar icon={<UserOutlined />} />
      </Dropdown>
    </Header>
  );
};

export default AppBar;
