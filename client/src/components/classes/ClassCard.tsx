import React from "react";
import { Card } from "antd";
import { SettingOutlined } from "@ant-design/icons";

const { Meta } = Card;

interface ClassCardProps {
  classTitle: string;
  creatorName: string;
}

const ClassCard: React.FC<ClassCardProps> = ({ classTitle }) => {
  return (
    <Card
      hoverable
      style={{ width: 240 }}
      actions={[<SettingOutlined key="setting" />]}
    >
      <Meta title={classTitle} description="www.instagram.com" />
    </Card>
  );
};

export default ClassCard;
