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

func main() {
    var fruits[2] string
    fruits[0] = "apple"
    fruits[1] = "orange"
    fmt.Println(fruits)

    fruitarr := [2] string {
		"Coco Cake",
		"Vanilla Cake"}
    fmt.Println(fruitarr)

    fslice := [] string {"c","b","a"}
    fmt.Println(len(fslice))
    fmt.Println(fslice[0:2])
}
