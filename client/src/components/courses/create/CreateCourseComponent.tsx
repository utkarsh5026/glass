import React, { useState } from "react";
import { Card, Segmented } from "antd";
import { useLocation } from "react-router-dom";
import CreateAssignment from "./CreateAssignment";
import CreateMaterial from "./CreateMaterial";

type ComponentType = "course" | "quiz" | "material" | "assignment";

const CreateCourseComponent: React.FC = () => {
  const location = useLocation();
  let componentType = "assignment" as ComponentType;
  if (location.state) componentType = location.state.compType as ComponentType;
  const [component, setComponent] = useState<ComponentType>(componentType);

  const segments = [
    {
      label: "Assignment",
      value: "assignment",
      component: <CreateAssignment />,
    },
    { label: "Material", value: "material" },
    { label: "Quiz", value: "quiz" },
  ];

  return (
    <Card
      title={`Create ${component} 😊`}
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
      {component === "assignment" ? <CreateAssignment /> : null}
      {component === "material" ? <CreateMaterial /> : null}
    </Card>
  );
};

export default CreateCourseComponent;
