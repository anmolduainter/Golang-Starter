package main

import "fmt"

func main() {

	// Lets Do some Pointers Stuffs
	var a int = 20
	var b *int

	b = &a

	// Lets check by printing the values -----
	fmt.Println(" Address of a : " , &a)
	fmt.Println("Content of a  : " , a)
	fmt.Print("Content of a checking through b : " , *b)


	// -----------------------Array Of Pointers----------------------------------------
	fmt.Println("-----------------------------------")
	const MAX int = 3
	a1 := []int{01,20,30}
	for i:=0 ; i<MAX ; i++{
		fmt.Print(" -> " , a1[i])
	}

	var ptr [MAX]*int
	for i:= 0 ;i<MAX ; i++{
		ptr[i] = &a1[i]
	}
	fmt.Println("Printing the value ::: ")
	for i:=0;i<MAX;i++ {
		fmt.Println(" -> " , *ptr[i])
	}

}
