package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"simpleapi/forms"
	"simpleapi/models"
)

//HandleCompany process company requests
func HandleCompany(c *gin.Context) {
	action, ok := c.Params.Get("action")
	if !ok || len(action) == 0 {
		c.JSON(http.StatusBadRequest, forms.Response{"action", "no action to do"})
		return
	}

	form := new(forms.Company)
	if err := c.Bind(form); err != nil {
		c.JSON(http.StatusBadRequest, forms.Response{"bind", "can't bind params"})
		return
	}

	// validate
	if errs := form.Validate(action); errs != nil {
		c.JSON(http.StatusBadRequest, errs)
		return
	}

	switch action {

	case "create":
		company := models.Company{
			Name:    form.Name,
			RegCode: form.RegCode,
		}

		if err := company.Create(); err != nil {
			c.JSON(http.StatusInternalServerError, forms.Response{"error", err.Error()})
		} else {
			c.JSON(http.StatusOK, company)
		}

	case "read":
		if form.ID <= 0 {
			c.JSON(http.StatusBadRequest, forms.Response{"id", "wrong id"})
			return
		}
		company, err := models.CompanyReadByID(form.ID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, forms.Response{"error", err.Error()})
		} else {
			c.JSON(http.StatusOK, company)
		}

	case "update":
		company := models.Company{
			Name:    form.Name,
			RegCode: form.RegCode,
		}

		if err := company.Update(); err != nil {
			c.JSON(http.StatusInternalServerError, forms.Response{"error", err.Error()})
		} else {
			c.JSON(http.StatusOK, company)
		}
	case "delete":
		if form.ID <= 0 {
			c.JSON(http.StatusBadRequest, forms.Response{"id", "wrong id"})
			return
		}
		err := models.CompanyDeleteByID(form.ID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, forms.Response{"error", err.Error()})
		} else {
			c.JSON(http.StatusOK, models.Company{ID: form.ID})
		}
	default:
		c.JSON(http.StatusBadRequest, forms.Response{"action", "wrong action"})
	}
	return
}
