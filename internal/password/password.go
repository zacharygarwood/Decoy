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
    iterations := 100000
    combined := []byte(masterPassword + accountId)

    initialKey := pbkdf2.Key(combined, combined, iterations, length, sha256.New)
    intermediate := sha256.Sum256(initialKey)
    finalKey := pbkdf2.Key(intermediate[:], combined, iterations, length, sha256.New)

    password := base64.URLEncoding.EncodeToString(finalKey)

    fmt.Println("Generated password: " + password)
    return password
}

