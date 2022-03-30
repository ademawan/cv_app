package work_experience

import (
	"cv_app/entities"
	"errors"

	"github.com/lithammer/shortuuid"

	"gorm.io/gorm"
)

type WorkExperienceRepository struct {
	database *gorm.DB
}

func New(db *gorm.DB) *WorkExperienceRepository {
	return &WorkExperienceRepository{
		database: db,
	}
}

func (tr *WorkExperienceRepository) Create(experience entities.WorkExperience) (entities.WorkExperience, error) {
	uid := shortuuid.New()
	experience.WorkExperienceUid = uid
	if err := tr.database.Create(&experience).Error; err != nil {
		return experience, err
	}

	return experience, nil
}

func (tr *WorkExperienceRepository) Get(userUid string) ([]entities.WorkExperience, error) {
	arrWorkExperience := []entities.WorkExperience{}
	result := tr.database.Where("user_uid =?", userUid).Find(&arrWorkExperience)

	if result.Error != nil {
		return nil, errors.New("failed to get work_experiences")
	}
	if result.RowsAffected == 0 {
		return arrWorkExperience, errors.New("experience is empty")
	}

	return arrWorkExperience, nil
}

func (tr *WorkExperienceRepository) GetByUid(userUid, workExperienceUid string) (entities.WorkExperience, error) {
	experience := entities.WorkExperience{}
	result := tr.database.Where("user_uid =? AND work_ekperience_uid =?", userUid, workExperienceUid).First(&experience)

	if result.Error != nil {
		return experience, errors.New("failed to get experience")
	}
	if result.RowsAffected == 0 {
		return experience, errors.New("experience not found")
	}

	return experience, nil
}

func (tr *WorkExperienceRepository) Update(workExperienceUid string, newWorkExperience entities.WorkExperience) (entities.WorkExperience, error) {

	var experience entities.WorkExperience
	result := tr.database.Where("user_uid =? AND work_ekperience_uid =?", newWorkExperience.UserUid, workExperienceUid).First(&experience)
	if result.Error != nil {
		return entities.WorkExperience{}, errors.New("failed to update experience")
	}
	if result.RowsAffected == 0 {
		return entities.WorkExperience{}, errors.New("experience not found")
	}

	if err := tr.database.Model(&experience).Updates(&newWorkExperience).Error; err != nil {
		return entities.WorkExperience{}, errors.New("failed to update experience")
	}

	return experience, nil
}

func (tr *WorkExperienceRepository) Delete(userUid, workExperienceUid string) error {
	result := tr.database.Where("user_uid =? AND work_ekperience_uid =?", userUid, workExperienceUid).Delete(&entities.WorkExperience{})
	if result.Error != nil {
		return result.Error
	}
	return nil

}
