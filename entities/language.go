package entities

import (
	"time"

	"gorm.io/gorm"
)

type Language struct {
	LanguageUid     string `gorm:"index;unique;type:varchar(22)" json:"language_uid"`
	LanguageName    string `gorm:"type:varchar(22)" json:"language_name"`
	UserUid         string `gorm:"index;type:varchar(22)" json:"-"`
	LanguageMeasure string `gorm:"type:enum('hight','medium','low')" json:"language_measure"`

	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
