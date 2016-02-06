package main

import  (
	"taxdb"
	"fmt"
)

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

func getFloat(name string) (float32 , bool) {
	fmt.Println("Please enter", name)
	var input float32
	_, err := fmt.Scanf("%b", &input)
	if err != nil {
		return 0.0 , false	
	}else{
		return input , true
	}

}

func main() {
	var testval taxdb.TaxData
	fmt.Println("Enter Tax Name")			
	fmt.Scanf("%s", testval.TaxName)
	fmt.Println("Enter Tax Environment")			
	fmt.Scanf("%s", testval.TaxEnv)
	fmt.Println("Enter Tax City")			
	fmt.Scanf("%s", testval.TaxCity)
	fmt.Println("Enter Tax State")			
	fmt.Scanf("%s", testval.TaxState)
	fmt.Println("Enter Tax Country")			
	fmt.Scanf("%s", testval.TaxCntry)
	fmt.Println("Enter Tax Rate")			
	fmt.Scanf("%b", testval.TaxRate)

	testval.Dbcommit()
}
