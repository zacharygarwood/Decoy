package main

import (
    "golang.org/x/crypto/pbkdf2"
    "encoding/base64"
    "crypto/sha256"
    "fmt"
)

func main() {
    // TODO: 
    // - Create separate module for Password
    //   - Will have a password struct with value, websiteId, length
    //   - Will have generate() function that does what generatePassword does
    // - Find CLI library
    //   - Allow the user to enter all the information in
    //   - Allow user to copy terminal output
    //   - Allow user to choose different types of password (standard, long, short, symbols

    masterPassword := "I'm a decoy password!"
    websiteId := "amazon.com"
    passwordLength := 32

    fmt.Println(generatePassword(masterPassword, websiteId, passwordLength))
}

func generatePassword(
    master string,
    website string,
    length int,
) string {
    iterations := 4096

    key := pbkdf2.Key([]byte(master), []byte(website), iterations, length, sha256.New)
    password := base64.URLEncoding.EncodeToString(key)

    return password
}
