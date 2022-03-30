package language

import (
	"cv_app/delivery/controllers/common"
	"cv_app/entities"
	"cv_app/middlewares"
	"cv_app/repository/language"
	"net/http"

	"github.com/labstack/echo/v4"
)

type LanguageController struct {
	repo language.Language
}

func New(repository language.Language) *LanguageController {
	return &LanguageController{
		repo: repository,
	}
}
func (tc *LanguageController) Create() echo.HandlerFunc {
	return func(c echo.Context) error {
		language := LanguageRequestFormat{}
		userUid := middlewares.ExtractTokenUserUid(c)
		c.Bind(&language)

		err := c.Validate(&language)
		if err != nil {
			return c.JSON(http.StatusBadRequest, common.ResponseUser(http.StatusBadRequest, "There is some problem from input", nil))
		}

		res, err := tc.repo.Create(entities.Language{
			UserUid:         userUid,
			LanguageName:    language.LanguageName,
			LanguageMeasure: language.LanguageMeasure,
		})

		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.ResponseUser(http.StatusInternalServerError, "There is some error on server", nil))
		}

		return c.JSON(http.StatusCreated, common.ResponseUser(http.StatusCreated, "Success create language", res))
	}
}

func (tc *LanguageController) Get() echo.HandlerFunc {
	return func(c echo.Context) error {
		userUid := middlewares.ExtractTokenUserUid(c)

		res, err := tc.repo.Get(userUid)

		if err != nil {
			statusCode := http.StatusInternalServerError
			errorMessage := "There is some problem from the server"
			if err.Error() == "language is empty" {
				statusCode = http.StatusOK
				errorMessage = err.Error()
			}
			return c.JSON(statusCode, common.ResponseUser(statusCode, errorMessage, nil))
		}

		return c.JSON(http.StatusOK, common.ResponseUser(http.StatusOK, "Success get all language", res))
	}
}

func (tc *LanguageController) GetByUid() echo.HandlerFunc {
	return func(c echo.Context) error {
		languageUid := c.Param("language_uid")
		userUid := middlewares.ExtractTokenUserUid(c)

		res, err := tc.repo.GetByUid(userUid, languageUid)

		if err != nil {
			statusCode := http.StatusInternalServerError
			errorMessage := "There is some problem from the server"
			if err.Error() == "language not found" {
				statusCode = http.StatusNotFound
				errorMessage = err.Error()
			}
			return c.JSON(statusCode, common.ResponseUser(http.StatusNotFound, errorMessage, nil))
		}

		return c.JSON(http.StatusOK, common.ResponseUser(http.StatusOK, "Success get language by uid", res))
	}
}

func (tc *LanguageController) Update() echo.HandlerFunc {
	return func(c echo.Context) error {
		var newLanguage = UpdateLanguageRequestFormat{}
		languageUid := c.Param("language_uid")
		userUid := middlewares.ExtractTokenUserUid(c)
		c.Bind(&newLanguage)

		err := c.Validate(&newLanguage)
		if err != nil {
			return c.JSON(http.StatusBadRequest, common.ResponseUser(http.StatusBadRequest, "There is some problem from input", nil))
		}
		res, err := tc.repo.Update(languageUid, entities.Language{
			UserUid:         userUid,
			LanguageName:    newLanguage.LanguageName,
			LanguageMeasure: newLanguage.LanguageMeasure,
		})

		if err != nil {
			statusCode := http.StatusInternalServerError
			errorMessage := "There is some problem from the server"
			if err.Error() == "language not found" {
				statusCode = http.StatusNotFound
				errorMessage = err.Error()
			}
			return c.JSON(statusCode, common.ResponseUser(http.StatusNotFound, errorMessage, nil))
		}

		return c.JSON(http.StatusOK, common.ResponseUser(http.StatusOK, "Success update language", res))
	}
}

func (tc *LanguageController) Delete() echo.HandlerFunc {
	return func(c echo.Context) error {
		languageUid := c.Param("language_uid")
		userUid := middlewares.ExtractTokenUserUid(c)

		err := tc.repo.Delete(userUid, languageUid)

		if err != nil {
			return c.JSON(http.StatusNotFound, common.ResponseUser(http.StatusNotFound, "not found", nil))
		}

		return c.JSON(http.StatusOK, common.ResponseUser(http.StatusOK, "Success delete language", nil))
	}
}
