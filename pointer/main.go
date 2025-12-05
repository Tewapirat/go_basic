package main

import "fmt"

func main() {

	// b := new(int)
	// c := new(*int)

	// a := 1
	// b = &a
	// c = &b

	// fmt.Printf("address a %v\n", &a)
	// fmt.Printf("address b %v\n", &b)
	// fmt.Printf("address c %v\n", &c)

	// fmt.Println()

	// fmt.Printf("value a %v\n", a)
	// fmt.Printf("value b %v\n", b)
	// fmt.Printf("value c %v\n", c)

	// fmt.Println()

	// fmt.Printf("*b %v\n", *b)
	// fmt.Printf("*c %v\n", *c)
	// fmt.Printf("**c %v\n", **c)


	// pass by pointer
	a := 2
	var n *int = &a

	double(n)

	fmt.Printf("address a %v\n",&a)
	fmt.Printf("value a %v\n",a)

	fmt.Println()

	fmt.Printf("address n %v\n",&n)
	fmt.Printf("address a in n %v\n",n)
	fmt.Printf("value n %v\n",*n)


	// pass by reference
	m := 2

	double(&m)
	fmt.Printf("value m %v\n",m)


}

func double(n *int){
	*n *=2
}