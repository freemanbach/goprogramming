/*********************************************************************************
* @author   Freeman
* @version  0.0.1 
* @link     https://golang.org/pkg/
* @since    2019.07.29
* @Desc     
*            
***********************************************************************************/

package main


import (
	"fmt"
)


func main() {

    // array aka tuple with pointers dereferenced
    var numbers = [6]int {1, 2, 3, 5}
    for i,x:= range numbers {
        fmt.Printf("value of x = %d at %d, %x\n", x, i, &numbers[i])
    }

    // pointers
    var a int = 19
    fmt.Printf("Pointer is: %x and value is: %d\n", &a, a)

    // slice aka list
    // var items := make([]int, 3, 5)

    for a := 0; a<10; a++ {
	fmt.Println("num: " , a)
    }

    for b := 0; b <= 10; b++ {
        if b % 5 == 0 {
            fmt.Println("Kool")
        } else if b %2 ==0 {
            fmt.Println("Kool, Beanie")
        }
    }
}
