package entities

import (
	"time"

	"gorm.io/gorm"
)

type Ability struct {
	AbilityUid     string         `gorm:"index;unique;type:varchar(22)" json:"ability_uid"`
	UserUid        string         `gorm:"index;type:varchar(22)" json:"-"`
	AbilityName    string         `gorm:"type:varchar(30)" json:"ability_name"`
	AbilityMeasure string         `gorm:"type:enum('ekspert','hight','medium','low')" json:"ability_measure"`
	Note           string         `json:"note"`
	CreatedAt      time.Time      `json:"created_at"`
	UpdatedAt      time.Time      `gorm:"-" json:"-"`
	DeletedAt      gorm.DeletedAt `gorm:"index" json:"-"`
}
