/*********************************************************************************
* @author   Freeman
* @version  0.0.1 
* @link     https://golang.org/pkg/
* @since    2019.07.28
* @Desc     THis lang is kind of difficult to get used to
*           Must use what is called in namespace and stated variables in a file, 
*	    otherwise error. 
***********************************************************************************/

package main


import (
	"os"
	"fmt"
	"time"
	"bufio"
	"syscall"
	"strconv"
	"math/rand"
)

// check to see if its digit
func isDigit(mess string) bool {
    _, err := strconv.ParseFloat(mess, 64)
    return err == nil
}

// random number range
func randInt(min int, max int) int {
    return min + rand.Intn(max-min)
}

func playgame() {

    scan := bufio.NewScanner(os.Stdin)
    fmt.Print("Want to play this game: (y or n) ? ")
    scan.Scan()
    play := scan.Text()

    for play == "y" || play == "yes"  {
        fmt.Print("Enter in a value inclusive (0 - 10) ? ")
        scan.Scan()
        temp := scan.Text()
        if isDigit(temp) == true  {
            num, err := strconv.ParseInt(temp, 10, 64)
            if err != nil {
                fmt.Println(err)
            }
            num1 := int(num)
            guessMyNum(num1)
            fmt.Println()
	    fmt.Print("Want to play: (y or n) ? ")
	    scan.Scan()
	    play := scan.Text()
	    if ( play == "n" || play == "no") {
		    fmt.Println("Exiting Game.....")
		    syscall.Exit(0)
            } else {
		    play = "y"
              }
         } else {
                fmt.Println("Exiting Game.....")
		syscall.Exit(0)
         }
    }
}

func guessMyNum(num int) {
    rand.Seed(time.Now().UnixNano())
    // rand.Seed(ns.Nanoseconds())
    pc := randInt(1, 11)

    if num < pc {
        fmt.Println("Your guess is lower than PC Number. ", "PC Number: ", pc)
    } else if num > pc {
        fmt.Println("Your guess is higher than PC Number. ", "PC Number: ", pc)
    } else if pc == num {
	fmt.Println("You have guessed correctly.")
    } else {
	fmt.Println("Error: something is incorrect.")
	syscall.Exit(1)
    }
}

func main() {
    playgame()
}
