/**
* @author  Freeman
* @version 0.0.1 
* @link https://golang.org/pkg/
* @since   2019.07.28
* @Desc   
*/


package main

import (
	"fmt"
	"math"
)

func main() {
	var aString string = "Hello Planet"
	var test bool = true
	var aNum int = 90
	var sizeA float32 = 2.3456
	var sizeB float64 = 2.3456
	var sizeC complex64 = 3.1415926

	var age int32 = 38
	fmt.Println(age)

	fmt.Println(aString)
	fmt.Println(test)
	fmt.Println(aNum)
	fmt.Println(sizeA)
	fmt.Println(sizeB)
	fmt.Println(sizeC)

	fmt.Printf("%T\n", age)

	fmt.Println(math.Floor(19.98876))
	fmt.Println(math.Ceil(sizeB))
	fmt.Println(math.Sqrt(90))
}
