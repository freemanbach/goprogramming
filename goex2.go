package main

/**
* @author  Freeman
* @version 0.0.1 
* @link https://golang.org/pkg/
* @since   2018.12.11
*/


import (
  "bufio"
  "fmt"
  "os"
)


func userInput1() {
    var age int
    fmt.Print("Enter a Number: ")
    fmt.Scanln(&age)
    fmt.Println(age)
}


func userInput2() {
    var word string
    reader := bufio.NewReader(os.Stdin)
    fmt.Print("What is your name? ")
    word, _ = reader.ReadString('\n')
    fmt.Println(word)
}


func main() {

    userInput1()
    userInput2()
}
