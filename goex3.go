package main

/**
* @author  Freeman
* @version 0.0.1 
* @link https://golang.org/pkg/
* @since   2018.12.12
*/

import "fmt"

func main() {

    // varible declaration
    var grade float32

    // prompt a user for input
    fmt.Print("Enter a grade in value: ")
    fmt.Scanln(&grade)

    // sorting out the grades
    if grade >=90 {
        fmt.Println("A")
    } else if grade >= 80 {
	fmt.Println("B")
    } else if grade >= 70 {
	fmt.Println("C")
    } else if grade >= 60 {
	fmt.Println("D")
    } else {
	fmt.Println("F")
    }
}
