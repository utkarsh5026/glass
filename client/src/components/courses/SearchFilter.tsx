import React, { useState } from "react";
import { Input, Button, Drawer, Space, Select, Slider, Switch } from "antd";
import { SearchOutlined, FilterOutlined } from "@ant-design/icons";
import styled from "styled-components";

const { Search } = Input;
const { Option } = Select;

interface CourseSearchAndFiltersProps {
  onSearch: (value: string) => void;
  onFilterChange: (filters: FilterState) => void;
  categories: string[];
}

interface FilterState {
  category: string;
  difficulty: [number, number];
  isActive: boolean;
}

const CourseSearchAndFilters: React.FC<CourseSearchAndFiltersProps> = ({
  onSearch,
  onFilterChange,
  categories,
}) => {
  const [isFilterVisible, setIsFilterVisible] = useState(false);
  const [filters, setFilters] = useState<FilterState>({
    category: "All",
    difficulty: [1, 5],
    isActive: true,
  });

  const handleFilterChange = (newFilters: Partial<FilterState>) => {
    const updatedFilters = { ...filters, ...newFilters };
    setFilters(updatedFilters);
    onFilterChange(updatedFilters);
  };

  return (
    <>
      <SearchContainer>
        <StyledSearch
          placeholder="Search courses"
          onSearch={onSearch}
          enterButton={<SearchOutlined />}
        />
        <Button
          icon={<FilterOutlined />}
          onClick={() => setIsFilterVisible(true)}
        >
          Filters
        </Button>
      </SearchContainer>

      <Drawer
        title="Course Filters"
        placement="right"
        onClose={() => setIsFilterVisible(false)}
        open={isFilterVisible}
      >
        <Space direction="vertical" size="middle" style={{ width: "100%" }}>
          <div>
            <h4>Category</h4>
            <Select
              style={{ width: "100%" }}
              value={filters.category}
              onChange={(value) => handleFilterChange({ category: value })}
            >
              <Option value="All">All Categories</Option>
              {categories.map((category) => (
                <Option key={category} value={category}>
                  {category}
                </Option>
              ))}
            </Select>
          </div>

          <div>
            <h4>Difficulty</h4>
            <Slider
              range
              min={1}
              max={5}
              value={filters.difficulty}
              onChange={(value: number | number[]) =>
                handleFilterChange({ difficulty: value as [number, number] })
              }
              marks={{ 1: "1", 2: "2", 3: "3", 4: "4", 5: "5" }}
            />
          </div>

          <div>
            <h4>Status</h4>
            <Switch
              checked={filters.isActive}
              onChange={(checked) => handleFilterChange({ isActive: checked })}
              checkedChildren="Active"
              unCheckedChildren="All"
            />
          </div>
        </Space>
      </Drawer>
    </>
  );
};

const SearchContainer = styled.div`
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 16px;
`;

const StyledSearch = styled(Search)`
  width: 300px;
`;
export default CourseSearchAndFilters;
