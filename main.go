package main

import "fmt"

func main()  {

	// letter and numbers AEIOUY10 have been removed
	// this prevents that bad or voulair words as well as names (companies
	// or persons) are generated

	letters := []string{"B","C","D","F","G","H","J","K","L","M","N","P","Q","R","S","T","V","W","X","Y","Z"}
	numbers := []int{2,3,4,5,6,7,8,9}

	for _, i := range(letters) {
		fmt.Println(i)
	}

	for _, t := range(numbers) {
		fmt.Println(t)
	}

}
