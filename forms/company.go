package forms

import "gopkg.in/go-playground/validator.v8"

//Company data struct
type Company struct {
	ID      int64  `form:"id" json:"id"`
	Name    string `form:"name" json:"name"`
	RegCode string `form:"reg_code" json:"reg_code"`
}

//Validate check form values
func (f *Company) Validate(action string) (errs []*Response) {

	config := &validator.Config{TagName: "validate"}
	validate := validator.New(config)
	if action == "create" {
		err := validate.Field(f.Name, "required")
		if err != nil {
			errs = append(errs, &Response{"name", "wrong company name"})
		}

	} else {
		err := validate.Field(f.ID, "required")
		if err != nil {
			errs = append(errs, &Response{"id", "wrong id"})
		}
	}
	return errs
}
