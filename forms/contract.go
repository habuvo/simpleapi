package forms

import (
	"gopkg.in/go-playground/validator.v8"
	"time"
)

// Contract data struct
type Contract struct {
	ID             int64     `form:"id" json:"id"`
	Seller         int64     `form:"seller" json:"seller"`
	Client         int64     `form:"client" json:"client"`
	ContractNumber string    `form:"contract_number" json:"contract_number"`
	DateSigned     time.Time `form:"date_signed" json:"date_signed"`
	ValidTill      time.Time `form:"valid_till" json:"valid_till"`
	CreditsInit    uint      `form:"credits_init" json:"credits_init"`
}

//Validate check values in request
func (f *Contract) Validate(action string) (errs []*Response) {

	config := &validator.Config{TagName: "validate"}
	validate := validator.New(config)
	if action == "create" {
		err := validate.Field(f.Seller, "required")
		if err != nil {
			errs = append(errs, &Response{"seller", "wrong seller id"})
		}

		err = validate.Field(f.Client, "required")
		if err != nil {
			errs = append(errs, &Response{"client", "wrong client id"})
		}

		err = validate.Field(f.DateSigned, "required")
		if err != nil {
			errs = append(errs, &Response{"date_signed", "wrong signed date"})
		}

		err = validate.Field(f.ValidTill, "required")
		if err != nil {
			errs = append(errs, &Response{"date_till", "wrong revoked data "})
		}

		err = validate.Field(f.CreditsInit, "required")
		if err != nil {
			errs = append(errs, &Response{"credits", "wrong credits amount"})
		}

	} else {
		err := validate.Field(f.ID, "required")
		if err != nil {
			errs = append(errs, &Response{"id", "wrong id"})
		}
	}
	return errs
}
