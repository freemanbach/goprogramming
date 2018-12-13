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
        fmt.Println(grade , " is an A")
    } else if grade >= 80 {
	fmt.Println(grade, " is a B")
    } else if grade >= 70 {
	fmt.Println(grade, " is a C")
    } else if grade >= 60 {
	fmt.Println(grade, " is a D")
    } else {
	fmt.Println(grade, " is a F")
    }
}
