package achievement

import "time"

//----------------------------------------------------
//REQUEST FORMAT
//----------------------------------------------------
type AchievementRequestFormat struct {
	UserUid   string
	Title     string    `json:"title" form:"title" validate:"required,min=2,max=25"`
	Note      string    `json:"note" form:"note"`
	StartDate time.Time `json:"start_date" form:"start_date" validate:"required"`
	EndDate   time.Time `json:"end_date" form:"end_date" validate:"required"`
}

type UpdateAchievementRequestFormat struct {
	UserUid   string
	Title     string `json:"title" form:"title" validate:"omitempty,min=2,max=25"`
	Note      string `json:"note" form:"note"`
	StartDate time.Time
	EndDate   time.Time
}
