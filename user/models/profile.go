package user_models

import (
	"time"

	"gorm.io/gorm"
)

type Skill struct {
	gorm.Model
	Name      string  `json:"name,omitempty"`
	Level     float32 `json:"level,omitempty"`
	ProfileID uint    `json:"profile_id,omitempty"`
}
type Project struct {
	gorm.Model
	ProfileID uint      `json:"profile_id,omitempty"`
	From      time.Time `json:"from,omitempty"`
	To        time.Time `json:"to,omitempty"`
	Role      string    `json:"role,omitempty"`
	Company   string    `json:"company,omitempty"`
	Purpose   string    `json:"purpose,omitempty"`
	Tasks     string    `json:"tasks,omitempty"`
	SkillTags string    `json:"skill_tags,omitempty"`
}
type Profile struct {
	gorm.Model
	UserID            uint      `json:"user_id,omitempty"`
	Role              string    `json:"role,omitempty"`
	Location          string    `json:"location,omitempty"`
	Phone             string    `json:"phone,omitempty"`
	Website           string    `json:"website,omitempty"`
	Email             string    `json:"email,omitempty"`
	Github            string    `json:"github,omitempty"`
	Stackoverflow     string    `json:"stackoverflow,omitempty"`
	Linkedin          string    `json:"linkedin,omitempty"`
	Whoami            string    `json:"whoami,omitempty"`
	ProgrammingSkills []Skill   `json:"programming_skills,omitempty"`
	Tools             []Skill   `json:"tools,omitempty"`
	Projects          []Project `json:"projects,omitempty"`
}
