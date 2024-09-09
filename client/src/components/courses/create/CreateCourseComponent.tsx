import React, { useState } from "react";
import { Card, Col, Row, Segmented } from "antd";
import { useLocation } from "react-router-dom";

type ComponentType = "course" | "quiz" | "material" | "assignment";

const CreateCourseComponent: React.FC = () => {
  const location = useLocation();
  let componentType = "assignment" as ComponentType;
  if (location.state) componentType = location.state.compType as ComponentType;
  const [component, setComponent] = useState<ComponentType>(componentType);

  alert(componentType);

  const segments = [
    { label: "Assignment", value: "assignment" },
    { label: "Material", value: "material" },
    { label: "Questions", value: "questions" },
  ];

  return (
    <Card
      bordered={false}
      activeTabKey={component}
      onTabChange={(key) => setComponent(key as ComponentType)}
      extra={
        <Segmented
          options={segments}
          onChange={(key) => setComponent(key as ComponentType)}
          value={component}
        />
      }
    >
      <Row gutter={[16, 16]}>
        <Col span={18} />
      </Row>
    </Card>
  );
};

export default CreateCourseComponent;
