package achievement

import "cv_app/entities"

type Achievement interface {
	Get(userUid string) ([]entities.Achievement, error)
	GetByUid(userUid, achievementUid string) (entities.Achievement, error)
	Create(newAchievement entities.Achievement) (entities.Achievement, error)
	Update(achievementUid string, newAchievement entities.Achievement) (entities.Achievement, error)
	Delete(userUid, achievementUid string) error
}
