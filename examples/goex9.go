/*********************************************************************************
* @author   Freeman
* @version  0.0.1 
* @link     https://golang.org/pkg/
* @since    2019.07.30
* @Desc     
*           
***********************************************************************************/

package main

import (
	"fmt"
)

func main() {

    // insert individual key-value into a map
    email := make(map[string] string)
    email["flo"] = "flo@google.com"
    email["john"] = "john@google.com"
    fmt.Println(email)
    fmt.Println(email["john"])

    // delete and print
    delete(email, "john")
    fmt.Println(email)

    // add multiple keys-values into a map
    address := map[string] string {
	"hob" : "hob@google.com",
	"ace" : "ace@google.com"}
    fmt.Println(address)
}
