package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"simpleapi/forms"
)

//MakePurchase handle purchase create
func MakePurchase(c *gin.Context) {

	form := new(forms.Purchase)
	if err := c.Bind(form); err != nil {
		c.JSON(http.StatusBadRequest, forms.Response{"bind", "can't bind params"})
		return
	}

	// validate
	if errs := form.Validate(); errs != nil {
		c.JSON(http.StatusBadRequest, errs)
		return
	}

	// do
	creds, err := form.Do()
	if err != nil {
		c.JSON(http.StatusBadRequest, forms.Response{"do", err.Error()})
		return
	}

	// ok
	c.JSON(http.StatusOK, forms.Response{"credits_remain", creds})
	return
}
