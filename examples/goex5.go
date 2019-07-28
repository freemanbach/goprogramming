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
)

func getsum(a int, b int) int {
	return a + b
}

func greet(name string) string {

	return "Hello, " + name
}

func main() {
	fmt.Println(greet("YourName"))
	fmt.Println(getsum(3,5))
}
