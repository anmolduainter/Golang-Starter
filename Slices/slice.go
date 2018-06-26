package main
import "fmt"

// Function to print slice
func printSlice(x []int){
	fmt.Println("len -> ",len(x)," Capacity -> ",cap(x) ," Slice-> " , x)
}

// Main Function
func main(){

	// Make an array
	arr := []int{1,2,3,4,4,5,5,6,6}

	// Take a look at slice
	printSlice(arr)

	// Lets do something amazing....
	fmt.Println(" 1 and 4 indexes between elements : " , arr[1:(4 + 1)])
	fmt.Println("0 and 3 indexes between elements : " , arr[:(3+1)])

	// Making of slice (type , length , capacity)
	numbers1 := make([]int , 0 , 5)
	printSlice(numbers1)

	// append function

	// Adding 2 in the slice and printing it
	printSlice(append(numbers1 , 2))
	//
	//// Copy Function
	//var numbers []int
	//numbers = append()
}
