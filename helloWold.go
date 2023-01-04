package main
import ("fmt")

const (
	A int=1
	B string="Deepak Verma"
	C bool= true
)
func testFun(){
	var student1 string
	student1 = "kanak";
	fmt.Println(student1);
}

func main() {
	fmt.Println("-------Chapter 1---------")
	fmt.Println("Hello Mr. Deepak")
	var firstVariable  = 1;
	x := 2;
	fmt.Println(firstVariable);
	fmt.Println(x)
	fmt.Println("--------Chapter 2--------")
	testFun();

	// chapter 3
	fmt.Println("-------Chapter 3--------")
	var a,b,c,d ,e int =1,3,5,4,5
	fmt.Println(a,b,c,d,e);

	fmt.Println("------Chapter 4---------")
	// A="hello"
	fmt.Println(A,B,C);
	fmt.Printf("DeepakVerm%%")
}