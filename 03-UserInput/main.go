package main

import (
	"bufio"
	"fmt"
	"os"
)

func main(){
fmt.Println("Welcome to User Input")

reader := bufio.NewReader(os.Stdin)
fmt.Println("Enter the rating of our website: ")

//comma ok || err ok
// _ here means that we are ignoring errors
input, _ := reader.ReadString('\n')
fmt.Println("Thanks for rating,", input)
fmt.Printf("Type of this rating is %T",input)

}