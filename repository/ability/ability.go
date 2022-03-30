package ability

import (
	"cv_app/entities"
	"errors"

	"github.com/lithammer/shortuuid"

	"gorm.io/gorm"
)

type AbilityRepository struct {
	database *gorm.DB
}

func New(db *gorm.DB) *AbilityRepository {
	return &AbilityRepository{
		database: db,
	}
}

func (tr *AbilityRepository) Create(ability entities.Ability) (entities.Ability, error) {

	uid := shortuuid.New()
	ability.AbilityUid = uid

	if err := tr.database.Create(&ability).Error; err != nil {
		return ability, err
	}

	return ability, nil
}

func (tr *AbilityRepository) Get(userUid string) ([]entities.Ability, error) {
	arrAbility := []entities.Ability{}
	result := tr.database.Where("user_uid =?", userUid).Find(&arrAbility)

	if result.Error != nil {
		return nil, errors.New("failed to get abilities")
	}
	if result.RowsAffected == 0 {
		return arrAbility, errors.New("ability is empty")
	}

	return arrAbility, nil
}

func (tr *AbilityRepository) GetByUid(userUid, abilityUid string) (entities.Ability, error) {
	ability := entities.Ability{}
	result := tr.database.Where("user_uid =? AND ability_uid =?", userUid, abilityUid).First(&ability)

	if result.Error != nil {
		return ability, errors.New("failed to get ability")
	}
	if result.RowsAffected == 0 {
		return ability, errors.New("ability not found")
	}

	return ability, nil
}

func (tr *AbilityRepository) Update(abilityUid string, newAbility entities.Ability) (entities.Ability, error) {

	var ability entities.Ability
	result := tr.database.Where("user_uid =? AND ability_uid =?", newAbility.UserUid, abilityUid).First(&ability)
	if result.Error != nil {
		return entities.Ability{}, errors.New("failed to update ability")
	}
	if result.RowsAffected == 0 {
		return entities.Ability{}, errors.New("ability not found")
	}

	if err := tr.database.Model(&ability).Updates(&newAbility).Error; err != nil {
		return entities.Ability{}, errors.New("failed to update ability")
	}

	return ability, nil
}

func (tr *AbilityRepository) Delete(userUid, abilityUid string) error {
	result := tr.database.Where("user_uid =? AND ability_uid =?", userUid, abilityUid).Delete(&entities.Ability{})
	if result.Error != nil {
		return result.Error
	}
	return nil

}
