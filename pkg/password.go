package password

import "golang.org/x/crypto/bcrypt"

// Function to generate a hash using string
func HashPassword(password string) (string, error) {
    bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
    return string(bytes), err
}