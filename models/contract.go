package models

import (
	"simpleapi/lib"
	"time"
)

//Contract is data type for contract
type Contract struct {
	ID             int64
	SellerID       int64
	ClientID       int64
	ContractNumber string
	Signed         time.Time
	ValidTill      time.Time
	CreditsInit    uint
	CreditsRemain  uint
}

//GetContractByID returns contract by item id
func GetContractByID(id int64) (cont Contract, err error) {

	err = lib.Env.DB.QueryRow("SELECT * FROM contract WHERE ID = ?", id).
		Scan(&cont.ID, &cont.SellerID, &cont.ClientID, &cont.ContractNumber, &cont.Signed, &cont.ValidTill, &cont.CreditsInit, &cont.CreditsRemain)
	return
}

//Create creates contract item
func (c *Contract) Create() error {
	resp, err := lib.Env.DB.Exec("INSERT INTO contract (seller_id,client_id,contract_number,signed,valid_till,credits_init,credits_remain) "+
		" VALUES (?,?,?,?,?,?,?)", c.SellerID, c.ClientID, c.ContractNumber, c.Signed, c.ValidTill, c.CreditsInit, c.CreditsInit)
	if resp != nil {
		c.ID, _ = resp.LastInsertId()
	}
	return err
}

//Update updates contract items values
func (c *Contract) Update() error {
	_, err := lib.Env.DB.Exec("UPDATE contract set contract_number = ?, credits_init = ?, credits_remain = ? WHERE id = ?", c.ContractNumber, c.CreditsInit, c.CreditsRemain, c.ID)
	return err
}

//ContractDeleteByID delete contract item by item id
func ContractDeleteByID(id int64) error {
	_, err := lib.Env.DB.Exec("DELETE FROM contract WHERE id = ?", id)
	return err
}
