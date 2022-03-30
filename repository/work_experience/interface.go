package work_experience

import "cv_app/entities"

type WorkExperience interface {
	Get(userUid string) ([]entities.WorkExperience, error)
	GetByUid(userUid, experienceUid string) (entities.WorkExperience, error)
	Create(newWorkExperience entities.WorkExperience) (entities.WorkExperience, error)
	Update(experienceUid string, newWorkExperience entities.WorkExperience) (entities.WorkExperience, error)
	Delete(userUid, experienceUid string) error
}
