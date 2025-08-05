package main

import "fmt"

func main(){

	// operators && || < <= > => ! ==
	//age :=20
	//sex := "male"
	//isPretty:=false

	// if age > 18 {
	// 	fmt.Println("You can marry...")
	// } else if age < 18{
	// 	fmt.Println("You cant marry")
	// } else {
	// 	fmt.Println("You are a teenager...")
	// }

	/* if age==20 && sex ==" male" {
		fmt.Println("You can marry..!!!")
	} else { //no extra line before else. otherwise error
		fmt.Println("You cant Marry...")
	}*/
	/*if !isPretty {
		fmt.Println("MashAllah")
	} else {
		fmt.Println("You are also pretty")
	}*/
	/////Switch Case
	a:=3
	switch a {
	case 1:
		fmt.Println("a is 1")
	case 2,3:
		fmt.Println("a is either 2 or 3")
	default:
		fmt.Println("a is neither 1, 2 or 3")
	}
}
