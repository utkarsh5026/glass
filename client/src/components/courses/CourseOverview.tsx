import React from "react";
import { Row, Col, Tabs } from "antd";
import {
  AppstoreOutlined,
  FileOutlined,
  UsergroupAddOutlined,
  MessageOutlined,
  CalendarOutlined,
} from "@ant-design/icons";
import Announcement from "./announcement/Announcement";
import CoursePeople from "./people/CoursePeople";
import Activity from "../admin/activity/Activity";

const CourseOverview: React.FC = () => {
  const onTabChange = (key: string) => {
    console.log(key);
  };
  return (
    <Row
      style={{
        border: "1px solid white",
        padding: "10px",
        borderRadius: "10px",
        width: "90vw",
        height: "90vh",
      }}
    >
      <Col span={24}>
        <Tabs onChange={onTabChange}>
          <Tabs.TabPane tab="Announcements" key="1" icon={<AppstoreOutlined />}>
            <Announcement />
          </Tabs.TabPane>
          <Tabs.TabPane
            tab="Files"
            key="2"
            icon={<FileOutlined />}
          ></Tabs.TabPane>
          <Tabs.TabPane tab="People" key="3" icon={<UsergroupAddOutlined />}>
            <CoursePeople />
          </Tabs.TabPane>
          <Tabs.TabPane
            tab="Chat"
            key="4"
            icon={<MessageOutlined />}
          ></Tabs.TabPane>
          <Tabs.TabPane tab="Activity" key="5" icon={<CalendarOutlined />}>
            <Activity />
          </Tabs.TabPane>
        </Tabs>
      </Col>
    </Row>
  );
};

export default CourseOverview;
