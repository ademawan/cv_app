package entities

import (
	"time"

	"gorm.io/gorm"
)

type Achievement struct {
	AchievementUid string `gorm:"index;unique;type:varchar(22)" json:"achievement_uid"`
	UserUid        string `gorm:"index;type:varchar(22)" json:"-"`
	Title          string `gorm:"type:varchar(30)" json:"title"`
	Note           string `gorm:"type:varchar(250)" json:"note"`
	StartDate      time.Time
	EndDate        time.Time

	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
