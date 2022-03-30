package work_experience

import (
	"net/http"
	"cv_app/delivery/controllers/common"
	"cv_app/entities"
	"cv_app/middlewares"
	"cv_app/repository/work_experience"

	"github.com/labstack/echo/v4"
)

type WorkExperienceController struct {
	repo work_experience.WorkExperience
}

func New(repository work_experience.WorkExperience) *WorkExperienceController {
	return &WorkExperienceController{
		repo: repository,
	}
}
func (tc *WorkExperienceController) Create() echo.HandlerFunc {
	return func(c echo.Context) error {
		workExperience := WorkExperienceRequestFormat{}
		userUid := middlewares.ExtractTokenUserUid(c)
		c.Bind(&workExperience)

		err := c.Validate(&workExperience)
		if err != nil {
			return c.JSON(http.StatusBadRequest, common.ResponseUser(http.StatusBadRequest, "There is some problem from input", nil))
		}

		res, err := tc.repo.Create(entities.WorkExperience{
			UserUid:     userUid,
			CompanyName: workExperience.CompanyName,
			Position:    workExperience.Position,
			StartDate:   workExperience.StartDate,
			EndDate:     workExperience.EndDate,
		})

		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.ResponseUser(http.StatusInternalServerError, "There is some error on server", nil))
		}

		return c.JSON(http.StatusCreated, common.ResponseUser(http.StatusCreated, "Success create work_experience", res))
	}
}

func (tc *WorkExperienceController) Get() echo.HandlerFunc {
	return func(c echo.Context) error {
		userUid := middlewares.ExtractTokenUserUid(c)

		res, err := tc.repo.Get(userUid)

		if err != nil {
			statusCode := http.StatusInternalServerError
			errorMessage := "There is some problem from the server"
			if err.Error() == "work_experience is empty" {
				statusCode = http.StatusOK
				errorMessage = err.Error()
			}
			return c.JSON(statusCode, common.ResponseUser(statusCode, errorMessage, nil))
		}

		return c.JSON(http.StatusOK, common.ResponseUser(http.StatusOK, "Success get all work_experience", res))
	}
}

func (tc *WorkExperienceController) GetByUid() echo.HandlerFunc {
	return func(c echo.Context) error {
		workExperienceUid := c.Param("work_experience_uid")
		userUid := middlewares.ExtractTokenUserUid(c)

		res, err := tc.repo.GetByUid(userUid, workExperienceUid)

		if err != nil {
			statusCode := http.StatusInternalServerError
			errorMessage := "There is some problem from the server"
			if err.Error() == "work_experience not found" {
				statusCode = http.StatusNotFound
				errorMessage = err.Error()
			}
			return c.JSON(statusCode, common.ResponseUser(http.StatusNotFound, errorMessage, nil))
		}

		return c.JSON(http.StatusOK, common.ResponseUser(http.StatusOK, "Success get work_experience by uid", res))
	}
}

func (tc *WorkExperienceController) Update() echo.HandlerFunc {
	return func(c echo.Context) error {
		var newWorkExperience = UpdateWorkExperienceRequestFormat{}
		workExperienceUid := c.Param("work_experience_uid")
		userUid := middlewares.ExtractTokenUserUid(c)
		c.Bind(&newWorkExperience)

		err := c.Validate(&newWorkExperience)
		if err != nil {
			return c.JSON(http.StatusBadRequest, common.ResponseUser(http.StatusBadRequest, "There is some problem from input", nil))
		}
		res, err := tc.repo.Update(workExperienceUid, entities.WorkExperience{
			UserUid:     userUid,
			CompanyName: newWorkExperience.CompanyName,
			Position:    newWorkExperience.Position,
			StartDate:   newWorkExperience.StartDate,
			EndDate:     newWorkExperience.EndDate,
		})

		if err != nil {
			statusCode := http.StatusInternalServerError
			errorMessage := "There is some problem from the server"
			if err.Error() == "work_experience not found" {
				statusCode = http.StatusNotFound
				errorMessage = err.Error()
			}
			return c.JSON(statusCode, common.ResponseUser(http.StatusNotFound, errorMessage, nil))
		}

		return c.JSON(http.StatusOK, common.ResponseUser(http.StatusOK, "Success update work_experience", res))
	}
}

func (tc *WorkExperienceController) Delete() echo.HandlerFunc {
	return func(c echo.Context) error {
		workExperienceUid := c.Param("work_experience_uid")
		userUid := middlewares.ExtractTokenUserUid(c)

		err := tc.repo.Delete(userUid, workExperienceUid)

		if err != nil {
			return c.JSON(http.StatusNotFound, common.ResponseUser(http.StatusNotFound, "not found", nil))
		}

		return c.JSON(http.StatusOK, common.ResponseUser(http.StatusOK, "Success delete work_experience", nil))
	}
}
