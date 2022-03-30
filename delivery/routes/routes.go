package routes

import (
	"cv_app/delivery/controllers/ability"
	"cv_app/delivery/controllers/achievement"
	"cv_app/delivery/controllers/auth"
	"cv_app/delivery/controllers/education"
	"cv_app/delivery/controllers/language"
	"cv_app/delivery/controllers/work_experience"

	"cv_app/delivery/controllers/user"
	"cv_app/middlewares"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func RegisterPath(e *echo.Echo,
	aa *auth.AuthController,
	uc *user.UserController,
	aab *ability.AbilityController,
	aac *achievement.AchievementController,
	edc *education.EducationController,
	lc *language.LanguageController,
	ww *work_experience.WorkExperienceController,

) {

	//CORS
	e.Use(middleware.CORS())

	//LOGGER
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}",
	}))

	//ROUTE REGISTER - LOGIN USERS
	e.POST("users/register", uc.Register())
	e.POST("users/login", aa.Login())

	//ROUTE USERS
	e.GET("/users/me", uc.GetByUid(), middlewares.JwtMiddleware())
	e.PUT("/users/me", uc.Update(), middlewares.JwtMiddleware())
	e.DELETE("/users/me", uc.Delete(), middlewares.JwtMiddleware())
	//ROUTE ABILITY
	e.POST("/users/me/abilities", aab.Create(), middlewares.JwtMiddleware())
	e.GET("/users/me/abilities", aab.GetByUid(), middlewares.JwtMiddleware())
	e.GET("/users/me/abilities/:ability_uid", aab.GetByUid(), middlewares.JwtMiddleware())
	e.PUT("/users/me/abilities/:ability_uid", aab.Update(), middlewares.JwtMiddleware())
	e.DELETE("/users/me/abilities/:ability_uid", aab.Delete(), middlewares.JwtMiddleware())
	//ROUTE ACHIVEMENT
	e.POST("/users/me/achievements", aac.Create(), middlewares.JwtMiddleware())
	e.GET("/users/me/achievements", aac.GetByUid(), middlewares.JwtMiddleware())
	e.GET("/users/me/achievements/:achievement_uid", aac.GetByUid(), middlewares.JwtMiddleware())
	e.PUT("/users/me/achievements/:achievement_uid", aac.Update(), middlewares.JwtMiddleware())
	e.DELETE("/users/me/achievements/:achievement_uid", aac.Delete(), middlewares.JwtMiddleware())
	//ROUTE EDUCATION
	e.POST("/users/me/educations", edc.Create(), middlewares.JwtMiddleware())
	e.GET("/users/me/educations", edc.GetByUid(), middlewares.JwtMiddleware())
	e.GET("/users/me/educations/:education_uid", edc.GetByUid(), middlewares.JwtMiddleware())
	e.PUT("/users/me/educations/:education_uid", edc.Update(), middlewares.JwtMiddleware())
	e.DELETE("/users/me/educations/:education_uid", edc.Delete(), middlewares.JwtMiddleware())
	//ROUTE LANGUAGE
	e.POST("/users/me/languages", lc.Create(), middlewares.JwtMiddleware())
	e.GET("/users/me/languages", lc.GetByUid(), middlewares.JwtMiddleware())
	e.GET("/users/me/languages/:language_uid", lc.GetByUid(), middlewares.JwtMiddleware())
	e.PUT("/users/me/languages/:language_uid", lc.Update(), middlewares.JwtMiddleware())
	e.DELETE("/users/me/languages/:language_uid", lc.Delete(), middlewares.JwtMiddleware())
	//ROUTE WORK EXPERIENCE
	e.POST("/users/me/work-experience", ww.Create(), middlewares.JwtMiddleware())
	e.GET("/users/me/work-experience", ww.GetByUid(), middlewares.JwtMiddleware())
	e.GET("/users/me/work-experience/:work_experience_uid", ww.GetByUid(), middlewares.JwtMiddleware())
	e.PUT("/users/me/work-experience/:work_experience_uid", ww.Update(), middlewares.JwtMiddleware())
	e.DELETE("/users/me/work-experience/:work_experience_uid", ww.Delete(), middlewares.JwtMiddleware())

}
