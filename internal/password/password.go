package password

import (
    "golang.org/x/crypto/pbkdf2"
    "encoding/base64"
    "crypto/sha256"
    "fmt"
)

func Generate(
    masterPassword string,
    accountId string,
    length int,
) string {
    iterations := 4096
    key := pbkdf2.Key([]byte(masterPassword), []byte(accountId), iterations, length, sha256.New)
    password := base64.URLEncoding.EncodeToString(key)

    fmt.Println("Generated password: " + password)
    return password
}

