package database

import (
	"golang.org/x/crypto/bcrypt"
)

type Password interface {
	HashPassword(password string) (string, error)
	DoPasswordsMatch(hashedPassword, currPassword string) bool
}

// HashPassword takes a plain text password as input and returns the hashed
// password using the bcrypt algorithm. It uses bcrypt's minimum cost for
// hashing. If an error occurs during the hashing process, it returns an
// empty string and the error.
//
// Parameters:
//   - password: The plain text password to be hashed.
//
// Returns:
//   - A string representing the hashed password.
//   - An error if the hashing process fails.
func HashPassword(password string) (string, error) {
	// Convert password string to byte slice
	passwordBytes := []byte(password)

	// Hash password with Bcrypt's min cost
	hashedPasswordBytes, err := bcrypt.
		GenerateFromPassword(passwordBytes, bcrypt.MinCost)

	return string(hashedPasswordBytes), err
}

// DoPasswordsMatch compares a hashed password with a plain text password
// to check if they match.
//
// Parameters:
// - hashedPassword: the hashed password stored in the database.
// - currPassword: the plain text password to compare.
//
// Returns:
// - bool: true if the passwords match, false otherwise.
func DoPasswordsMatch(hashedPassword, currPassword string) bool {
	err := bcrypt.CompareHashAndPassword(
		[]byte(hashedPassword), []byte(currPassword))

	return err == nil
}
