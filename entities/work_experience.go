package entities

import (
	"time"

	"gorm.io/gorm"
)

type WorkExperience struct {
	WorkExperienceUid string `gorm:"index;unique;type:varchar(22)" json:"education_uid"`
	UserUid           string `gorm:"index;type:varchar(22)" json:"-"`

	CompanyName string `gorm:"type:varchar(30)" json:"company"`
	Position    string `gorm:"type:varchar(50)" json:"position"`

	StartDate time.Time
	EndDate   time.Time

	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
