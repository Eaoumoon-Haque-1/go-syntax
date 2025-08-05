package main
import "fmt"

var(
	a = 2
	b = 5
)
func add(c int, d int){
	sum:=c+d
	fmt.Println("Hello, this is your sum: ", sum)
}

func main(){
	p:=10
	q:=20
	add(p,a)
	add(p,q)
	if p>=10 {
		x:= 5
		mul:=x*p*a
		fmt.Println("Hello, this is your mul: ", mul)
	}
}
// All global variables and functions are part of global
// for a variable needed they will check with in their own scope and then in global scope
// for if block things are a bit different as the scope is within a scope first they will check in their own block than main and than global