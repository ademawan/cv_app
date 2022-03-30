package work_experience

import "time"

//----------------------------------------------------
//REQUEST FORMAT
//----------------------------------------------------
type WorkExperienceRequestFormat struct {
	UserUid     string
	CompanyName string    `json:"company_name" form:"company_name" validate:"required,min=2,max=25"`
	Position    string    `json:"position" form:"position" validate:"required,min=2,max=25"`
	StartDate   time.Time `json:"start_date" form:"start_date" validate:"required"`
	EndDate     time.Time `json:"end_date" form:"end_date" validate:"required"`
}

type UpdateWorkExperienceRequestFormat struct {
	UserUid     string
	CompanyName string `json:"company_name" form:"company_name" validate:"omitempty,min=2,max=25"`
	Position    string `json:"position" form:"position" validate:"omitempty,min=2,max=25"`
	StartDate   time.Time
	EndDate     time.Time
}
