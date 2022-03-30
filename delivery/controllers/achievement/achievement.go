package achievement

import (
	"net/http"
	"cv_app/delivery/controllers/common"
	"cv_app/entities"
	"cv_app/middlewares"
	"cv_app/repository/achievement"

	"github.com/labstack/echo/v4"
)

type AchievementController struct {
	repo achievement.Achievement
}

func New(repository achievement.Achievement) *AchievementController {
	return &AchievementController{
		repo: repository,
	}
}
func (tc *AchievementController) Create() echo.HandlerFunc {
	return func(c echo.Context) error {
		achievement := AchievementRequestFormat{}
		userUid := middlewares.ExtractTokenUserUid(c)
		c.Bind(&achievement)

		err := c.Validate(&achievement)
		if err != nil {
			return c.JSON(http.StatusBadRequest, common.ResponseUser(http.StatusBadRequest, "There is some problem from input", nil))
		}

		res, err := tc.repo.Create(entities.Achievement{
			UserUid:   userUid,
			Title:     achievement.Title,
			Note:      achievement.Note,
			StartDate: achievement.StartDate,
			EndDate:   achievement.EndDate,
		})

		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.ResponseUser(http.StatusInternalServerError, "There is some error on server", nil))
		}

		return c.JSON(http.StatusCreated, common.ResponseUser(http.StatusCreated, "Success create achievement", res))
	}
}

func (tc *AchievementController) Get() echo.HandlerFunc {
	return func(c echo.Context) error {
		userUid := middlewares.ExtractTokenUserUid(c)

		res, err := tc.repo.Get(userUid)

		if err != nil {
			statusCode := http.StatusInternalServerError
			errorMessage := "There is some problem from the server"
			if err.Error() == "achievement is empty" {
				statusCode = http.StatusOK
				errorMessage = err.Error()
			}
			return c.JSON(statusCode, common.ResponseUser(statusCode, errorMessage, nil))
		}

		return c.JSON(http.StatusOK, common.ResponseUser(http.StatusOK, "Success get all achievement", res))
	}
}

func (tc *AchievementController) GetByUid() echo.HandlerFunc {
	return func(c echo.Context) error {
		achievementUid := c.Param("achievement_uid")
		userUid := middlewares.ExtractTokenUserUid(c)

		res, err := tc.repo.GetByUid(userUid, achievementUid)

		if err != nil {
			statusCode := http.StatusInternalServerError
			errorMessage := "There is some problem from the server"
			if err.Error() == "achievement not found" {
				statusCode = http.StatusNotFound
				errorMessage = err.Error()
			}
			return c.JSON(statusCode, common.ResponseUser(http.StatusNotFound, errorMessage, nil))
		}

		return c.JSON(http.StatusOK, common.ResponseUser(http.StatusOK, "Success get achievement by uid", res))
	}
}

func (tc *AchievementController) Update() echo.HandlerFunc {
	return func(c echo.Context) error {
		var newAchievement = UpdateAchievementRequestFormat{}
		achievementUid := c.Param("achievement_uid")
		userUid := middlewares.ExtractTokenUserUid(c)
		c.Bind(&newAchievement)

		err := c.Validate(&newAchievement)
		if err != nil {
			return c.JSON(http.StatusBadRequest, common.ResponseUser(http.StatusBadRequest, "There is some problem from input", nil))
		}
		res, err := tc.repo.Update(achievementUid, entities.Achievement{
			UserUid:   userUid,
			Title:     newAchievement.Title,
			Note:      newAchievement.Note,
			StartDate: newAchievement.StartDate,
			EndDate:   newAchievement.EndDate,
		})

		if err != nil {
			statusCode := http.StatusInternalServerError
			errorMessage := "There is some problem from the server"
			if err.Error() == "achievement not found" {
				statusCode = http.StatusNotFound
				errorMessage = err.Error()
			}
			return c.JSON(statusCode, common.ResponseUser(http.StatusNotFound, errorMessage, nil))
		}

		return c.JSON(http.StatusOK, common.ResponseUser(http.StatusOK, "Success update achievement", res))
	}
}

func (tc *AchievementController) Delete() echo.HandlerFunc {
	return func(c echo.Context) error {
		achievementUid := c.Param("achievement_uid")
		userUid := middlewares.ExtractTokenUserUid(c)

		err := tc.repo.Delete(userUid, achievementUid)

		if err != nil {
			return c.JSON(http.StatusNotFound, common.ResponseUser(http.StatusNotFound, "not found", nil))
		}

		return c.JSON(http.StatusOK, common.ResponseUser(http.StatusOK, "Success delete achievement", nil))
	}
}
