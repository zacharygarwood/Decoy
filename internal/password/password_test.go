package password

import (
    "testing"
)

func TestGenerateConsistency(t *testing.T) {
    masterPassword := "password123"
    accountId := "example.com"
    length := 32

    result1 := Generate(masterPassword, accountId, length)
    result2 := Generate(masterPassword, accountId, length)

    if result1 != result2 {
        t.Errorf("Generate() returned different results for the same input")
    }
}

func TestGenerateDifferentPasswords(t *testing.T) {
    accountId := "example.com"
    length := 32

    result1 := Generate("password123", accountId, length)
    result2 := Generate("123password", accountId, length)

    if result1 == result2 {
        t.Errorf("Generate() returned the same results for different passwords")
    }
}

func TestGenerateDifferentAccountIds(t *testing.T) {
    masterPassword := "password123"
    length := 32

    result1 := Generate(masterPassword, "example.com", length)
    result2 := Generate(masterPassword, "test.com", length)

    if result1 == result2 {
        t.Errorf("Generate() returned the same results for different account identifiers")
    }
}
