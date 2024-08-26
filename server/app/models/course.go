package models

import "gorm.io/gorm"

type Course struct {
	gorm.Model
	Name        string `json:"name" gorm:"not null"`
	Description string `json:"description"`
	StartDate   string `json:"start_date"`
	EndDate     string `json:"end_date"`
	MaxStudents int    `json:"max_students"`
	Difficulty  string `json:"difficulty"`
	Category    string `json:"category"`
	IsActive    bool   `json:"is_active"`
}

func (c *Course) TableName() string {
	return "courses"
}
