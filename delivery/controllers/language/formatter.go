package language

//----------------------------------------------------
//REQUEST FORMAT
//----------------------------------------------------
type LanguageRequestFormat struct {
	UserUid         string
	LanguageName    string `json:"language_name" form:"language_name" validate:"required,min=2,max=25"`
	LanguageMeasure string `json:"language_measure" form:"language_measure" validate:"required,min=3,max=6"`
}

type UpdateLanguageRequestFormat struct {
	UserUid         string
	LanguageName    string `json:"language_name" form:"language_name" validate:"omitempty,min=2,max=25"`
	LanguageMeasure string `json:"language_measure" form:"language_measure" validate:"omitempty,min=3,max=6"`
}
