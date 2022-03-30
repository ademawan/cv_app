package education

import "time"

//----------------------------------------------------
//REQUEST FORMAT
//----------------------------------------------------
type EducationRequestFormat struct {
	UserUid   string
	Academy   string    `json:"academy" form:"academy" validate:"required,min=2,max=25"`
	Major     string    `json:"major" form:"major"`
	StartDate time.Time `json:"start_date" form:"start_date"`
	EndDate   time.Time `json:"end_date" form:"end_date"`
}

type UpdateEducationRequestFormat struct {
	UserUid   string
	Academy   string `json:"academy" form:"academy" validate:"omitempty,min=2,max=25"`
	Major     string `json:"major" form:"major"`
	StartDate time.Time
	EndDate   time.Time
}
