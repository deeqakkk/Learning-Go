package main
import "fmt"

func main(){

	var username string="deepak"
	fmt.Println(username)
	fmt.Printf("variable is of type: %T \n", username)

	var isLoggedIn bool= true
	fmt.Println(isLoggedIn)
	fmt.Printf("isLoggedIn is of type: %T \n", isLoggedIn)

	var smallFloat float32= 233.222232323232
	fmt.Println(smallFloat)
	fmt.Printf("isLoggedIn is of type: %T \n", smallFloat)

	var largeFloat float64= 233.222232323232
	fmt.Println(largeFloat)
	fmt.Printf("isLoggedIn is of type: %T \n", largeFloat)

	//default values and some aliases
	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("Variable is of type: %T \n", anotherVariable);

	var stringCheck string
	fmt.Println(stringCheck)
	fmt.Printf("Variable is of type: %T \n", stringCheck)


	userName := "Deepak"
	fmt.Println(userName)
	fmt.Printf("Variable is of type: %T \n", userName)

	

}