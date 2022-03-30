package language

import (
	"errors"
	"cv_app/entities"

	"github.com/lithammer/shortuuid"

	"gorm.io/gorm"
)

type LanguageRepository struct {
	database *gorm.DB
}

func New(db *gorm.DB) *LanguageRepository {
	return &LanguageRepository{
		database: db,
	}
}

func (tr *LanguageRepository) Create(language entities.Language) (entities.Language, error) {

	uid := shortuuid.New()
	language.LanguageUid = uid
	if err := tr.database.Create(&language).Error; err != nil {
		return language, err
	}

	return language, nil
}

func (tr *LanguageRepository) Get(userUid string) ([]entities.Language, error) {
	arrLanguage := []entities.Language{}
	result := tr.database.Where("user_uid =?", userUid).Find(&arrLanguage)

	if result.Error != nil {
		return nil, errors.New("failed to get languages")
	}
	if result.RowsAffected == 0 {
		return arrLanguage, errors.New("language is empty")
	}

	return arrLanguage, nil
}

func (tr *LanguageRepository) GetByUid(userUid, languageUid string) (entities.Language, error) {
	language := entities.Language{}
	result := tr.database.Where("user_uid =? AND language_uid =?", userUid, languageUid).First(&language)

	if result.Error != nil {
		return language, errors.New("failed to get language")
	}
	if result.RowsAffected == 0 {
		return language, errors.New("language not found")
	}

	return language, nil
}

func (tr *LanguageRepository) Update(languageUid string, newLanguage entities.Language) (entities.Language, error) {

	var language entities.Language
	result := tr.database.Where("user_uid =? AND language_uid =?", newLanguage.UserUid, languageUid).First(&language)
	if result.Error != nil {
		return entities.Language{}, errors.New("failed to update language")
	}
	if result.RowsAffected == 0 {
		return entities.Language{}, errors.New("language not found")
	}

	if err := tr.database.Model(&language).Updates(&newLanguage).Error; err != nil {
		return entities.Language{}, errors.New("failed to update language")
	}

	return language, nil
}

func (tr *LanguageRepository) Delete(userUid, languageUid string) error {
	result := tr.database.Where("user_uid =? AND language_uid =?", userUid, languageUid).Delete(&entities.Language{})
	if result.Error != nil {
		return result.Error
	}
	return nil

}
