package main

import (
	"fmt"
	"time"
)

func main(){
	fmt.Println("Welcome to time study of GoLang")

	presentTime := time.Now()
	fmt.Println(presentTime)

	fmt.Println(presentTime.Date())
	fmt.Println(presentTime.Day())
	fmt.Println(presentTime.GoString())
	fmt.Println(presentTime.Weekday())
	fmt.Println(presentTime.Format("01-02-2006 Monday"))

	// following line gives weird results
	fmt.Println(presentTime.Format("01-02-2022"))


	createdDate := time.Date(2020, time.December, 10,9,9,0,0,time.Local)
	fmt.Println(createdDate)

	// if we pass something other than Monday dayName gets distorted
	fmt.Println(createdDate.Format("01-02-2006 Monday"))
}