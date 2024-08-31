package cli

import (
    "bufio"
    "fmt"
    "os"
    "syscall"

    "golang.org/x/term"
    "decoy/internal/password"
)

func Run() (string, error) {
    masterPassword, err := getMasterPassword()
    if err != nil {
        return "", err
    }

    accountId, err := getAccountId()
    if err != nil {
        return "", err
    }

    passwordLength := 32
    return password.Generate(string(masterPassword), accountId, passwordLength), nil
}

func getMasterPassword() ([]byte, error) {
    fmt.Print("Enter master password: ")
    return term.ReadPassword(int(syscall.Stdin))
}

func getAccountId() (string, error) {
    reader := bufio.NewReader(os.Stdin)
    fmt.Print("\nEnter account identifier: ")
    return reader.ReadString('\n')
}

