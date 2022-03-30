package ability

//----------------------------------------------------
//REQUEST FORMAT
//----------------------------------------------------
type AbilityRequestFormat struct {
	UserUid        string
	AbilityName    string `json:"ability" form:"ability" validate:"required,min=2,max=25"`
	AbilityMeasure string `json:"ability_measure" form:"ability_measure" validate:"required,min=2,max=25"`
	Note           string `json:"note" form:"note"`
}

type UpdateAbilityRequestFormat struct {
	UserUid        string
	AbilityName    string `json:"ability_name" form:"ability_name"  validate:"omitempty,min=2,max=25"`
	AbilityMeasure string `json:"ability_measure" form:"ability_measure" validate:"omitempty,required,min=2,max=25"`

	Note string `json:"note" form:"note"`
}
