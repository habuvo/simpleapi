package forms

import (
	"fmt"
	"gopkg.in/go-playground/validator.v8"
	"simpleapi/lib"
	"simpleapi/models"
	"strconv"
	"time"
)

// Purchase data struct
type Purchase struct {
	Contract int64 `form:"contract" json:"contract"`
	Credits  uint  `form:"credits" json:"credits"`
}

//Validate check request values
func (f *Purchase) Validate() (errs []*Response) {

	config := &validator.Config{TagName: "validate"}
	validate := validator.New(config)

	err := validate.Field(f.Contract, "required")
	if err != nil {
		errs = append(errs, &Response{"contract", "wrong contract id"})
	}

	err = validate.Field(f.Credits, "required")
	if err != nil {
		errs = append(errs, &Response{"credits", "wrong credits amount"})
	}

	return errs
}

//Do process purchase request
func (f *Purchase) Do() (creds string, err error) {
	cont, err := models.GetContractByID(f.Contract)
	if err != nil {
		err = fmt.Errorf("server error")
		return
	}

	if cont.CreditsRemain < f.Credits {
		err = fmt.Errorf("not enough credits")
		return
	}

	if cont.ValidTill.Before(time.Now()) {
		err = fmt.Errorf("contract unactive")
		return
	}

	tx, err := lib.Env.DB.Begin()
	if err != nil {
		err = fmt.Errorf("server error")
		return
	}

	//decrese credits
	_, err = tx.Exec("UPDATE contract SET credits_remain = ? WHERE id = ?", cont.CreditsRemain-f.Credits, f.Contract)
	if err != nil {
		_ = tx.Rollback()
		err = fmt.Errorf("server error")
		return
	}

	//add purchase
	_, err = tx.Exec("INSERT INTO purchase (affecting_contract,credits_spent) VALUES (?,?)", f.Contract, cont.CreditsRemain-f.Credits)
	if err != nil {
		_ = tx.Rollback()
		err = fmt.Errorf("server error")
		return
	}

	if err = tx.Commit(); err != nil {
		err = fmt.Errorf("server error")
		return
	}

	creds = strconv.Itoa(int(cont.CreditsRemain - f.Credits))
	return
}
