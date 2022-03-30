package achievement

import (
	"cv_app/entities"
	"errors"

	"github.com/lithammer/shortuuid"

	"gorm.io/gorm"
)

type AchievementRepository struct {
	database *gorm.DB
}

func New(db *gorm.DB) *AchievementRepository {
	return &AchievementRepository{
		database: db,
	}
}

func (tr *AchievementRepository) Create(achievement entities.Achievement) (entities.Achievement, error) {

	uid := shortuuid.New()
	achievement.AchievementUid = uid

	if err := tr.database.Create(&achievement).Error; err != nil {
		return achievement, err
	}

	return achievement, nil
}

func (tr *AchievementRepository) Get(userUid string) ([]entities.Achievement, error) {
	arrAchievement := []entities.Achievement{}
	result := tr.database.Where("user_uid =?", userUid).Find(&arrAchievement)

	if result.Error != nil {
		return nil, errors.New("failed to get achievements")
	}
	if result.RowsAffected == 0 {
		return arrAchievement, errors.New("achievement is empty")
	}

	return arrAchievement, nil
}

func (tr *AchievementRepository) GetByUid(userUid, achievementUid string) (entities.Achievement, error) {
	achievement := entities.Achievement{}
	result := tr.database.Where("user_uid =? AND achievement_uid =?", userUid, achievementUid).First(&achievement)

	if result.Error != nil {
		return achievement, errors.New("failed to get achievement")
	}
	if result.RowsAffected == 0 {
		return achievement, errors.New("achievement not found")
	}

	return achievement, nil
}

func (tr *AchievementRepository) Update(achievementUid string, newAchievement entities.Achievement) (entities.Achievement, error) {

	var achievement entities.Achievement
	result := tr.database.Where("user_uid =? AND achievement_uid =?", newAchievement.UserUid, achievementUid).First(&achievement)
	if result.Error != nil {
		return entities.Achievement{}, errors.New("failed to update achievement")
	}
	if result.RowsAffected == 0 {
		return entities.Achievement{}, errors.New("achievement not found")
	}

	if err := tr.database.Model(&achievement).Updates(&newAchievement).Error; err != nil {
		return entities.Achievement{}, errors.New("failed to update achievement")
	}

	return achievement, nil
}

func (tr *AchievementRepository) Delete(userUid, achievementUid string) error {
	result := tr.database.Where("user_uid =? AND achievement_uid =?", userUid, achievementUid).Delete(&entities.Achievement{})
	if result.Error != nil {
		return result.Error
	}
	return nil

}
