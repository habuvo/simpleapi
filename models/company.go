package models

import "simpleapi/lib"

//Company name
//Company registration code (alphanumeric, optional)

//Company implement model for companies
type Company struct {
	ID      int64
	Name    string
	RegCode string
}

//Create creates company item
func (c *Company) Create() error {
	resp, err := lib.Env.DB.Exec("INSERT INTO company (name,reg_code) "+
		" VALUES (?,?)", c.Name, c.RegCode)
	if resp != nil {
		c.ID, _ = resp.LastInsertId()
	}
	return err
}

//CompanyReadByID returns company item by ID
func CompanyReadByID(id int64) (company *Company, err error) {
	err = lib.Env.DB.QueryRow("SELECT * from company WHERE id = ?", id).Scan(company)
	return
}

//Update updates company item values
func (c *Company) Update() error {
	_, err := lib.Env.DB.Exec("UPDATE company SET name = ?, reg_code = ?", c.Name, c.RegCode)
	return err
}

//CompanyDeleteByID deletes company item by item id
func CompanyDeleteByID(id int64) error {
	_, err := lib.Env.DB.Exec("DELETE FROM company WHERE id = ?", id)
	return err
}
