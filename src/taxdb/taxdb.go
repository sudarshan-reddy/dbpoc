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
	taxName string
	taxEnv string
	taxCity string
	taxState string
	taxCntry string
	taxRate float32
}

func (tx *TaxData) Dbcommit() bool {
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
	DB_USER, DB_PASSWD, DB_NAME)
	db , err := sql.Open("postgres", dbinfo)
	checkErr(err)
	defer db.Close()
	fmt.Println("Connected")
	return true
}

func (tax *TaxData) EnterTaxDetails (ipTaxName string,
									 ipTaxEnv string,
									 ipTaxCity string,
									 ipTaxState string,
									 ipTaxCntry string,
									 ipTaxRate float32) (state bool) {
	defer func(){
		if r := recover(); r != nil {
			state = false	
		}
	} ()
	tax.taxName    =  ipTaxName 
	tax.taxEnv     =  ipTaxEnv 
	tax.taxCity    =  ipTaxCity 
	tax.taxState   =  ipTaxState
	tax.taxCntry   =  ipTaxCntry
	tax.taxRate    =  ipTaxRate 
	state = true
	return state
								 
}

func (tax TaxData) commit () bool {
	return true
}
