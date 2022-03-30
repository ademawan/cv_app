package education

import "cv_app/entities"

type Education interface {
	Get(userUid string) ([]entities.Education, error)
	GetByUid(userUid, educationUid string) (entities.Education, error)
	Create(newEducation entities.Education) (entities.Education, error)
	Update(educationUid string, newEducation entities.Education) (entities.Education, error)
	Delete(userUid, educationUid string) error
}
