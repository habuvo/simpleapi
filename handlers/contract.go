package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"simpleapi/forms"
	"simpleapi/models"
)

//HandleContract process contract requests
func HandleContract(c *gin.Context) {

	action, ok := c.Params.Get("action")
	if !ok || len(action) == 0 {
		c.JSON(http.StatusBadRequest, forms.Response{"action", "no action to do"})
		return
	}

	form := new(forms.Contract)
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
		contract := models.Contract{
			SellerID:       form.Seller,
			ClientID:       form.Client,
			ContractNumber: form.ContractNumber,
			Signed:         form.DateSigned,
			ValidTill:      form.ValidTill,
			CreditsInit:    form.CreditsInit,
		}

		if err := contract.Create(); err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
		} else {
			c.JSON(http.StatusOK, contract)
		}

	case "read":
		if form.ID <= 0 {
			c.JSON(http.StatusBadRequest, forms.Response{"id", "wrong id"})
			return
		}
		contract, err := models.GetContractByID(form.ID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
		} else {
			c.JSON(http.StatusOK, contract)
		}

	case "update":
		contract := models.Contract{
			SellerID:       form.Seller,
			ClientID:       form.Client,
			ContractNumber: form.ContractNumber,
			Signed:         form.DateSigned,
			ValidTill:      form.ValidTill,
			CreditsInit:    form.CreditsInit,
		}

		if err := contract.Update(); err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
		} else {
			c.JSON(http.StatusOK, contract)
		}
	case "delete":
		if form.ID <= 0 {
			c.JSON(http.StatusBadRequest, forms.Response{"id", "wrong id"})
			return
		}
		err := models.ContractDeleteByID(form.ID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, forms.Response{"error", err.Error()})
		} else {
			c.JSON(http.StatusOK, models.Contract{ID: form.ID})
		}
	default:
		c.JSON(http.StatusBadRequest, forms.Response{"action", "wrong action"})
	}
	return
}
