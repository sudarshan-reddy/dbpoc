package main

import  (
	"taxdb"
	"fmt"
	"bufio"
	"os"
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
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter Tax Name: ")			
	testval.TaxName, _ = reader.ReadString('\n')
	fmt.Print("Enter Tax Environment: ")			
	testval.TaxEnv, _ = reader.ReadString('\n')
	fmt.Print("Enter Tax City: ")			
	testval.TaxCity, _ = reader.ReadString('\n')
	fmt.Print("Enter Tax State: ")			
	testval.TaxState, _ = reader.ReadString('\n')
	fmt.Print("Enter Tax Country: ")			
	testval.TaxCntry, _ = reader.ReadString('\n')
	fmt.Print("Enter Tax Rate: ")			
	fmt.Scanf("%b", &testval.TaxRate)
	testval.Dbcommit()
}
