package ability

import "cv_app/entities"

type Ability interface {
	Get(userUid string) ([]entities.Ability, error)
	GetByUid(userUid, abilityUid string) (entities.Ability, error)
	Create(newAbility entities.Ability) (entities.Ability, error)
	Update(abilityUid string, newAbility entities.Ability) (entities.Ability, error)
	Delete(userUid, abilityUid string) error
}
