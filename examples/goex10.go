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
    "crypto/sha1"
    "crypto/sha256"
    "crypto/sha512"
   b64 "encoding/base64"
    "fmt"
)

func hash_sha512() {
    mess := "Hello this is planet Earth."
    digest := sha512.New()
    digest.Write([] byte(mess))
    bs := digest.Sum(nil)
    fmt.Println(mess)
    fmt.Printf("%x\n", bs)
}

func hash_sha1() {
    mess := "Hello this is planet Earth."
    digest := sha1.New()
    digest.Write([] byte(mess))
    bs := digest.Sum(nil)
    fmt.Println(mess)
    fmt.Printf("%x\n", bs)
}

func hash_sha256(){
    mess := "Hello this is planet Earth."
    digest := sha256.New()
    digest.Write([] byte(mess))
    bs := digest.Sum(nil)
    fmt.Println(mess)
    fmt.Printf("%x\n", bs)
}

func b64Encode() {
    data := "Hello this is planet Earth."
    enc := b64.StdEncoding.EncodeToString([]byte(data))
    fmt.Println(data)
    fmt.Println(enc)
}

func main() {
    hash_sha1()
    hash_sha256()
    hash_sha512()
    b64Encode()
}
