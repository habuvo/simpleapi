package lib

import "database/sql"

const (
	companyInit = "CREATE TABLE IF NOT EXISTS company (" +
		" id MEDIUMINT NOT NULL AUTO_INCREMENT , " +
		" name CHAR(30) NOT NULL UNIQUE," +
		" reg_code CHAR(30), " +
		" PRIMARY KEY (id)" +
		" ) ENGINE=InnoDB"
	contractInit = "CREATE TABLE IF NOT EXISTS contract (" +
		" id MEDIUMINT NOT NULL AUTO_INCREMENT, " +
		" seller_id MEDIUMINT ," +
		" client_id MEDIUMINT REFERENCES company(id)," +
		" contract_number CHAR(30), " +
		" signed TIMESTAMP, " +
		" valid_till TIMESTAMP, " +
		" credits_init SMALLINT, " +
		" credits_remain SMALLINT," +
		" CONSTRAINT seller FOREIGN KEY (seller_id) REFERENCES company(id) ON DELETE CASCADE, " +
		" CONSTRAINT client FOREIGN KEY (client_id) REFERENCES company(id) ON DELETE CASCADE, " +
		"  PRIMARY KEY (id)" +
		" ) ENGINE InnoDB"
	purchaseInit = "CREATE TABLE IF NOT EXISTS purchase (" +
		"  id MEDIUMINT NOT NULL AUTO_INCREMENT , " +
		"  affecting_contract MEDIUMINT REFERENCES contract(id), " +
		"  purchase_time TIMESTAMP," +
		"  credits_spent SMALLINT, " +
		"  FOREIGN KEY (affecting_contract) REFERENCES contract(id) ON DELETE CASCADE," +
		" PRIMARY KEY (id)" +
		" ) ENGINE InnoDB"
)

//InitDB initializes DB with project structure
func InitDB(db *sql.DB) error {
	if _, err := db.Exec(companyInit); err != nil {
		return err
	}
	if _, err := db.Exec(contractInit); err != nil {
		return err
	}
	_, err := db.Exec(purchaseInit)
	return err
}
