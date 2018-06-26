package main
import "fmt"


func PrintSomething(num1 int, str string){
	// Printing Values
	fmt.Println(num1)
	fmt.Println(str)
}

// Function to swap ------------------- Call By Value
func swap(num int , num1 int) (int,int){
		return num1, num
}

// Function to swap ------------------- Call By reference
func swapCallByRef(num *int , num1 *int){
	var temp int
	temp = *num
	*num = *num1
	*num1 = temp
}


// Main Function
func main(){

	// Saying Hello (first line)
	fmt.Println("Hello Bro !!!!")

	// Make Some Variables

	// Static Type
	var x float64
	x = 20.0
	fmt.Println(x)
//	fmt.Println("to check type %T" , x)

	// Dynamic Type
	y := 20
	fmt.Println(y)
//	fmt.Println("Type of y : %T" , y)


	// Mixed Variable Declaration
	var a , b, c = 3 , 4 , "foo"
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	//Constant
	const LENGTH int = 10
	const WIDTH int = 20

	area := LENGTH * WIDTH
	fmt.Println("Value of Area :" ,area)

	// Calling Function
	PrintSomething(2,"Hello")

	// Swap something
	x1 := 2
	y1 := 3
	fmt.Println("Before Swap x :" , x1)
	fmt.Println("Before swap y : " ,y1)
	x1 , y1 = swap(x1 , y1)
	fmt.Println("After Swap x :" , x1)
	fmt.Println("After swap y : " ,y1)


	/////////// Swap Something---->> Function Reference --------------------------
	swapCallByRef(&x1,&y1)
	fmt.Println("After Swap x :" , x1)
	fmt.Println("After swap y : " ,y1)


	//////////// Single Dimension Array ////////////////////////////
	var n [10]int

	// Input the value in array
	for i := 0 ; i<10 ; i++{
		n[i] = i + 100
	}

	// Output the value in array

	fmt.Println("Printing array now........")
	for i := 0;i<10;i++{
		fmt.Println(n[i])
	}


	

}
