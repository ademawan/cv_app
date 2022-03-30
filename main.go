package main

import (
	"cv_app/configs"
	abc "cv_app/delivery/controllers/ability"
	acc "cv_app/delivery/controllers/achievement"
	ac "cv_app/delivery/controllers/auth"
	edc "cv_app/delivery/controllers/education"
	ll "cv_app/delivery/controllers/language"
	uc "cv_app/delivery/controllers/user"
	wc "cv_app/delivery/controllers/work_experience"
	"cv_app/delivery/routes"
	abilityRepo "cv_app/repository/ability"
	achievementRepo "cv_app/repository/achievement"
	authRepo "cv_app/repository/auth"
	educationRepo "cv_app/repository/education"
	languageRepo "cv_app/repository/language"
	userRepo "cv_app/repository/user"
	workExperienceRepo "cv_app/repository/work_experience"
	"cv_app/utils"
	"fmt"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"

	"github.com/labstack/gommon/log"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

func main() {
	config := configs.GetConfig()

	db := utils.InitDB(config)

	authRepo := authRepo.New(db)
	userRepo := userRepo.New(db)
	abilityRepo := abilityRepo.New(db)
	achievementRepo := achievementRepo.New(db)
	educationRepo := educationRepo.New(db)
	languageRepo := languageRepo.New(db)
	workExperienceRepo := workExperienceRepo.New(db)

	authController := ac.New(authRepo)
	userController := uc.New(userRepo)
	abilityController := abc.New(abilityRepo)
	achievementController := acc.New(achievementRepo)
	educationController := edc.New(educationRepo)
	languageController := ll.New(languageRepo)
	workExperienceController := wc.New(workExperienceRepo)

	e := echo.New()
	e.Validator = &CustomValidator{validator: validator.New()}

	routes.RegisterPath(
		e,
		authController,
		userController,
		abilityController,
		achievementController,
		educationController,
		languageController,
		workExperienceController,
	)

	log.Fatal(e.Start(fmt.Sprintf(":%d", config.Port)))
}
