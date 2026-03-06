package security

import (
    "crypto/rand"
    "encoding/base64"
    "errors"

    "golang.org/x/crypto/argon2"
)

const (
    memory      = 64 * 1024
    iterations  = 3
    parallelism = 2
    saltLength  = 16
    keyLength   = 32
)

func GeneratePasswordHash(password string) (string, error) {
    salt := make([]byte, saltLength)
    _, err := rand.Read(salt)
    if err != nil {
        return "", err
    }

    hash := argon2.IDKey([]byte(password), salt, iterations, memory, parallelism, keyLength)

    encoded := base64.RawStdEncoding.EncodeToString(append(salt, hash...))
    return encoded, nil
}

func ComparePasswordHash(password, encoded string) (bool, error) {
    data, err := base64.RawStdEncoding.DecodeString(encoded)
    if err != nil {
        return false, err
    }

    if len(data) < saltLength {
        return false, errors.New("invalid encoded hash")
    }

    salt := data[:saltLength]
    hash := data[saltLength:]

    newHash := argon2.IDKey([]byte(password), salt, iterations, memory, parallelism, keyLength)

    return string(newHash) == string(hash), nil
}
