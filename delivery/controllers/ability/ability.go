package ability

import (
	"cv_app/delivery/controllers/common"
	"cv_app/entities"
	"cv_app/middlewares"
	"cv_app/repository/ability"
	"net/http"

	"github.com/labstack/echo/v4"
)

type AbilityController struct {
	repo ability.Ability
}

func New(repository ability.Ability) *AbilityController {
	return &AbilityController{
		repo: repository,
	}
}
func (tc *AbilityController) Create() echo.HandlerFunc {
	return func(c echo.Context) error {
		ability := AbilityRequestFormat{}
		userUid := middlewares.ExtractTokenUserUid(c)
		c.Bind(&ability)

		err := c.Validate(&ability)
		if err != nil {
			return c.JSON(http.StatusBadRequest, common.ResponseUser(http.StatusBadRequest, "There is some problem from input", nil))
		}

		res, err := tc.repo.Create(entities.Ability{
			UserUid:        userUid,
			AbilityName:    ability.AbilityName,
			AbilityMeasure: ability.AbilityMeasure,
			Note:           ability.Note,
		})

		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.ResponseUser(http.StatusInternalServerError, "There is some error on server", nil))
		}

		return c.JSON(http.StatusCreated, common.ResponseUser(http.StatusCreated, "Success create ability", res))
	}
}

func (tc *AbilityController) Get() echo.HandlerFunc {
	return func(c echo.Context) error {
		userUid := middlewares.ExtractTokenUserUid(c)

		res, err := tc.repo.Get(userUid)

		if err != nil {
			statusCode := http.StatusInternalServerError
			errorMessage := "There is some problem from the server"
			if err.Error() == "ability is empty" {
				statusCode = http.StatusOK
				errorMessage = err.Error()
			}
			return c.JSON(statusCode, common.ResponseUser(statusCode, errorMessage, nil))
		}

		return c.JSON(http.StatusOK, common.ResponseUser(http.StatusOK, "Success get all ability", res))
	}
}

func (tc *AbilityController) GetByUid() echo.HandlerFunc {
	return func(c echo.Context) error {
		abilityUid := c.Param("ability_uid")
		userUid := middlewares.ExtractTokenUserUid(c)

		res, err := tc.repo.GetByUid(userUid, abilityUid)

		if err != nil {
			statusCode := http.StatusInternalServerError
			errorMessage := "There is some problem from the server"
			if err.Error() == "ability not found" {
				statusCode = http.StatusNotFound
				errorMessage = err.Error()
			}
			return c.JSON(statusCode, common.ResponseUser(http.StatusNotFound, errorMessage, nil))
		}

		return c.JSON(http.StatusOK, common.ResponseUser(http.StatusOK, "Success get ability by uid", res))
	}
}

func (tc *AbilityController) Update() echo.HandlerFunc {
	return func(c echo.Context) error {
		var newAbility = UpdateAbilityRequestFormat{}
		abilityUid := c.Param("ability_uid")
		userUid := middlewares.ExtractTokenUserUid(c)
		c.Bind(&newAbility)

		err := c.Validate(&newAbility)
		if err != nil {
			return c.JSON(http.StatusBadRequest, common.ResponseUser(http.StatusBadRequest, "There is some problem from input", nil))
		}
		res, err := tc.repo.Update(abilityUid, entities.Ability{
			UserUid:        userUid,
			AbilityName:    newAbility.AbilityName,
			AbilityMeasure: newAbility.AbilityMeasure,
			Note:           newAbility.Note,
		})

		if err != nil {
			statusCode := http.StatusInternalServerError
			errorMessage := "There is some problem from the server"
			if err.Error() == "ability not found" {
				statusCode = http.StatusNotFound
				errorMessage = err.Error()
			}
			return c.JSON(statusCode, common.ResponseUser(http.StatusNotFound, errorMessage, nil))
		}

		return c.JSON(http.StatusOK, common.ResponseUser(http.StatusOK, "Success update ability", res))
	}
}

func (tc *AbilityController) Delete() echo.HandlerFunc {
	return func(c echo.Context) error {
		abilityUid := c.Param("ability_uid")
		userUid := middlewares.ExtractTokenUserUid(c)

		err := tc.repo.Delete(userUid, abilityUid)

		if err != nil {
			return c.JSON(http.StatusNotFound, common.ResponseUser(http.StatusNotFound, "not found", nil))
		}

		return c.JSON(http.StatusOK, common.ResponseUser(http.StatusOK, "Success delete ability", nil))
	}
}
