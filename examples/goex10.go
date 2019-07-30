/*********************************************************************************
* @author   Freeman
* @version  0.0.1 
* @link     https://golang.org/pkg/
* @since    2019.07.29
* @Desc     
*            
***********************************************************************************/

package main

// kind of fun to play with cryptography
import (
//    "crypto/elliptic"
//    "crypto/aes"
//    "crypto/des"
//    "crypto/dsa"
    "crypto/sha512"
//    "encoding/base64"
    "fmt"
)

func encrypt_sha512() {
    mess := "Hello this is planet Earth."
    digest := sha512.New()
    digest.Write([] byte(mess))
    bs := digest.Sum(nil)
    fmt.Println(mess)
    fmt.Printf("%x\n", bs)
}

func main() {
    encrypt_sha512()
}
