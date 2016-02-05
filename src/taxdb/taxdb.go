package taxdb

import (
	"fmt"
	"database/sql"
	_ "github.com/lib/pq"
	"strings"
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

func getString(name string) (string, bool) {
	fmt.Println("Please enter" , name)	
	var input string
	_, err := fmt.Scanf("%s", &input)
	if err != nil {
		return "" , false	
	}else{
		return input , true
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
	if db.Ping() == nil {
		value, err := db.Query(`CREATE TABLE IF NOT EXISTS taxData( 
			id integer, 
			taxName varchar(20), 
			taxEnv varchar(10),
			taxCity varchar(20), 
			taxState varchar(8), 
			taxCntry varchar(15), 
			taxRate decimal)`)
		checkErr(err)
		if value == nil {
			fmt.Println("Value is nil")
		}
		return true
	}
	return false

}

func (tax *TaxData) EnterTaxDetails () (state bool) {
	defer func(){
		if r := recover(); r != nil {
			state = false	
		}
	} ()

	taxvals := []string{"Name,String",
						"Environment,String",
						"City,String",
						"State,String",
						"Country,String",
						"Rate,Float32"}

	for _, vals := range taxvals {
		result := strings.Split(vals, ",")
		if result[1] == "String" {
			input, err := getString(result[0])
			if err == false {
				fmt.Println(input)
			}
		}
	}
	
	state = true
	return state
								 
}
