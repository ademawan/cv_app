package education

import (
	"cv_app/entities"
	"errors"

	"github.com/lithammer/shortuuid"

	"gorm.io/gorm"
)

type EducationRepository struct {
	database *gorm.DB
}

func New(db *gorm.DB) *EducationRepository {
	return &EducationRepository{
		database: db,
	}
}

func (tr *EducationRepository) Create(education entities.Education) (entities.Education, error) {

	uid := shortuuid.New()
	education.EducationUid = uid

	if err := tr.database.Create(&education).Error; err != nil {
		return education, err
	}

	return education, nil
}

func (tr *EducationRepository) Get(userUid string) ([]entities.Education, error) {
	arrEducation := []entities.Education{}
	result := tr.database.Where("user_uid =?", userUid).Find(&arrEducation)

	if result.Error != nil {
		return nil, errors.New("failed to get educations")
	}
	if result.RowsAffected == 0 {
		return arrEducation, errors.New("education is empty")
	}

	return arrEducation, nil
}

func (tr *EducationRepository) GetByUid(userUid, educationUid string) (entities.Education, error) {
	education := entities.Education{}
	result := tr.database.Where("user_uid =? AND education_uid =?", userUid, educationUid).First(&education)

	if result.Error != nil {
		return education, errors.New("failed to get education")
	}
	if result.RowsAffected == 0 {
		return education, errors.New("education not found")
	}

	return education, nil
}

func (tr *EducationRepository) Update(educationUid string, newEducation entities.Education) (entities.Education, error) {

	var education entities.Education
	result := tr.database.Where("user_uid =? AND education_uid =?", newEducation.UserUid, educationUid).First(&education)
	if result.Error != nil {
		return entities.Education{}, errors.New("failed to update education")
	}
	if result.RowsAffected == 0 {
		return entities.Education{}, errors.New("education not found")
	}

	if err := tr.database.Model(&education).Updates(&newEducation).Error; err != nil {
		return entities.Education{}, errors.New("failed to update education")
	}

	return education, nil
}

func (tr *EducationRepository) Delete(userUid, educationUid string) error {
	result := tr.database.Where("user_uid =? AND education_uid =?", userUid, educationUid).Delete(&entities.Education{})
	if result.Error != nil {
		return result.Error
	}
	return nil

}
