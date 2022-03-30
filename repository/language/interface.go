package language

import "cv_app/entities"

type Language interface {
	Get(userUid string) ([]entities.Language, error)
	GetByUid(userUid, languageUid string) (entities.Language, error)
	Create(newLanguage entities.Language) (entities.Language, error)
	Update(languageUid string, newLanguage entities.Language) (entities.Language, error)
	Delete(userUid, languageUid string) error
}
