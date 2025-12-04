package main

import "fmt"

func main(){
	// var a [3]int = [3]int{1,2,3}

	// fmt.Println(a)

	// a[0] = 4

	// fmt.Println(a)


	// call all index array 

	// var a [3]int = [3]int{1,2,3}

	// for i,v := range a{
	// 	fmt.Println(a[i],v)
		
	// }



	// index and value

	// var a [3]int = [3]int{1,2,3}

	// for i, v := range a{
	// 	fmt.Println(i, v)
		
	// }


	
	// one parameter
	// var a [3]int = [3]int{1,2,3}

	// for _,v := range a{
	// 	fmt.Println(v)
		
	// }



	// Slice
	// b := []int{1,2,3}
	// fmt.Println(b)

	// b := make([]int, 0)

	// b = append(b, 1,2,3,4)

	// fmt.Println(b)

	// b[2] = 7

	// b = append(b, 5,6)

	// fmt.Println(b)




	// if using array value is copy but slice is pointer 

// 	a := [3]int{1,2,3}

// 	fmt.Println("Outside")
// 	for i := range a {
// 		fmt.Printf("%v \n", &a[i])
// 	}

// 	double(a)

// 	fmt.Println(a)

// }

// 	func double (nums [3]int){
// 		fmt.Println("Inside")
// 		for i := range nums{
// 			fmt.Printf("%v \n", &nums[i])
// 			nums[i] *= 2
// 		}



	a := []int{1,2,3}
	b := double(a)


	fmt.Println(a)
	fmt.Println(b)

}


func double (nums []int) []int {
		newNums := make([]int, len(nums))

		for i := range nums{
			newNums[i] = nums[i] * 2
		}

		return  newNums
	}