package taxdb

import (
	"fmt"
	"database/sql"
	_ "github.com/lib/pq"
)

const (
	DB_USER = "postgres"
	DB_PASSWD = "test1234"
	DB_NAME = "test"
)

func checkErr(err error) {
	if err != nil {
		panic(err)	
	}
}

type TaxData struct {
	TaxName string
	TaxEnv string
	TaxCity string
	TaxState string
	TaxCntry string
	TaxRate float32
}


func (tx *TaxData) Dbcommit() bool {
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
	DB_USER, DB_PASSWD, DB_NAME)
	db , err := sql.Open("postgres", dbinfo)
	checkErr(err)
	defer db.Close()
	if db.Ping() == nil {
		db.Query(`CREATE TABLE IF NOT EXISTS taxData( 
			id integer, 
			taxName varchar(20), 
			taxEnv varchar(10),
			taxCity varchar(20), 
			taxState varchar(8), 
			taxCntry varchar(15), 
			taxRate decimal)`)
		checkErr(err)
		insertQuery := fmt.Sprintf(`INSERT INTO taxData (taxName, taxEnv, taxCity,
		taxState, taxCntry, taxRate)
		VALUES ('%s', '%s', '%s', '%s', '%s', '%f')`,tx.TaxName,tx.TaxEnv,tx.TaxCity,tx.TaxState,tx.TaxCntry,tx.TaxRate)
		value , err := db.Query(insertQuery)
		fmt.Println(value, err)
		return true
	}
	return false

}

