package main

import  (
	"taxdb"
	"fmt"
	"bufio"
	"os"
	"io"
	"net/http"
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

func getInputFromCmdLine(db *taxdb.TaxData) bool {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter Tax Name: ")			
	db.TaxName, _ = reader.ReadString('\n')
	fmt.Print("Enter Tax Environment: ")			
	db.TaxEnv, _ = reader.ReadString('\n')
	fmt.Print("Enter Tax City: ")			
	db.TaxCity, _ = reader.ReadString('\n')
	fmt.Print("Enter Tax State: ")			
	db.TaxState, _ = reader.ReadString('\n')
	fmt.Print("Enter Tax Country: ")			
	db.TaxCntry, _ = reader.ReadString('\n')
	fmt.Print("Enter Tax Rate: ")			
	fmt.Scanf("%b", &db.TaxRate)
	return true

}

func ipform(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w , "Testing")
}

func getInputFromServer(db *taxdb.TaxData) bool {

	fmt.Println("Server online @ http://localhost:8001")
	http.HandleFunc("/", ipform)	
	http.ListenAndServe(":8001", nil)
	return true
}

func main() {
	var testval taxdb.TaxData
	//getInputFromCmdLine(&testval) 
	getInputFromServer(&testval)
	testval.Dbcommit()

}
