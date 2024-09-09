import React from "react";
import { Row, Col, Tabs, Button } from "antd";
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
    <div
      style={{
        height: "90vh",
        overflow: "auto",
        position: "relative",
      }}
    >
      <Row
        style={{
          padding: "10px",
          overflow: "auto",
        }}
      >
        <Col span={24}>
          <Tabs
            animated={{
              inkBar: true,
              tabPane: true,
            }}
            onChange={onTabChange}
            style={{
              position: "sticky",
              top: 0,
              zIndex: 1,
              backgroundColor: "white",
              paddingTop: "10px",
            }}
            tabBarExtraContent={{
              right: <Button>Add</Button>,
            }}
          >
            <Tabs.TabPane
              tab="Announcements"
              key="1"
              icon={<AppstoreOutlined />}
            >
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
    </div>
  );
};

export default CourseOverview;
