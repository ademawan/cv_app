package education

import (
	"cv_app/delivery/controllers/common"
	"cv_app/entities"
	"cv_app/middlewares"
	"cv_app/repository/education"
	"net/http"

	"github.com/labstack/echo/v4"
)

type EducationController struct {
	repo education.Education
}

func New(repository education.Education) *EducationController {
	return &EducationController{
		repo: repository,
	}
}
func (tc *EducationController) Create() echo.HandlerFunc {
	return func(c echo.Context) error {
		education := EducationRequestFormat{}
		userUid := middlewares.ExtractTokenUserUid(c)
		c.Bind(&education)

		err := c.Validate(&education)
		if err != nil {
			return c.JSON(http.StatusBadRequest, common.ResponseUser(http.StatusBadRequest, "There is some problem from input", nil))
		}

		res, err := tc.repo.Create(entities.Education{
			UserUid:   userUid,
			Academy:   education.Academy,
			Major:     education.Major,
			StartDate: education.StartDate,
			EndDate:   education.EndDate,
		})

		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.ResponseUser(http.StatusInternalServerError, "There is some error on server", nil))
		}

		return c.JSON(http.StatusCreated, common.ResponseUser(http.StatusCreated, "Success create education", res))
	}
}

func (tc *EducationController) Get() echo.HandlerFunc {
	return func(c echo.Context) error {
		userUid := middlewares.ExtractTokenUserUid(c)

		res, err := tc.repo.Get(userUid)

		if err != nil {
			statusCode := http.StatusInternalServerError
			errorMessage := "There is some problem from the server"
			if err.Error() == "education is empty" {
				statusCode = http.StatusOK
				errorMessage = err.Error()
			}
			return c.JSON(statusCode, common.ResponseUser(statusCode, errorMessage, nil))
		}

		return c.JSON(http.StatusOK, common.ResponseUser(http.StatusOK, "Success get all education", res))
	}
}

func (tc *EducationController) GetByUid() echo.HandlerFunc {
	return func(c echo.Context) error {
		educationUid := c.Param("education_uid")
		userUid := middlewares.ExtractTokenUserUid(c)

		res, err := tc.repo.GetByUid(userUid, educationUid)

		if err != nil {
			statusCode := http.StatusInternalServerError
			errorMessage := "There is some problem from the server"
			if err.Error() == "education not found" {
				statusCode = http.StatusNotFound
				errorMessage = err.Error()
			}
			return c.JSON(statusCode, common.ResponseUser(http.StatusNotFound, errorMessage, nil))
		}

		return c.JSON(http.StatusOK, common.ResponseUser(http.StatusOK, "Success get education by uid", res))
	}
}

func (tc *EducationController) Update() echo.HandlerFunc {
	return func(c echo.Context) error {
		var newEducation = UpdateEducationRequestFormat{}
		educationUid := c.Param("education_uid")
		userUid := middlewares.ExtractTokenUserUid(c)
		c.Bind(&newEducation)

		err := c.Validate(&newEducation)
		if err != nil {
			return c.JSON(http.StatusBadRequest, common.ResponseUser(http.StatusBadRequest, "There is some problem from input", nil))
		}
		res, err := tc.repo.Update(educationUid, entities.Education{
			UserUid:   userUid,
			Academy:   newEducation.Academy,
			Major:     newEducation.Major,
			StartDate: newEducation.StartDate,
			EndDate:   newEducation.EndDate,
		})

		if err != nil {
			statusCode := http.StatusInternalServerError
			errorMessage := "There is some problem from the server"
			if err.Error() == "education not found" {
				statusCode = http.StatusNotFound
				errorMessage = err.Error()
			}
			return c.JSON(statusCode, common.ResponseUser(http.StatusNotFound, errorMessage, nil))
		}

		return c.JSON(http.StatusOK, common.ResponseUser(http.StatusOK, "Success update education", res))
	}
}

func (tc *EducationController) Delete() echo.HandlerFunc {
	return func(c echo.Context) error {
		educationUid := c.Param("education_uid")
		userUid := middlewares.ExtractTokenUserUid(c)

		err := tc.repo.Delete(userUid, educationUid)

		if err != nil {
			return c.JSON(http.StatusNotFound, common.ResponseUser(http.StatusNotFound, "not found", nil))
		}

		return c.JSON(http.StatusOK, common.ResponseUser(http.StatusOK, "Success delete education", nil))
	}
}
