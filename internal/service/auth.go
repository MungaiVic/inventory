package service

import "golang.org/x/crypto/bcrypt"

// HashPassword generates a password hash for storage in the DB.
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

/*
ValidatePassword will be used to check if user has entered correct password before allowing jwt token generation.

# It will take in the hashed password from the DB and a user-supplied password  and return true if they match

Example:

	ValidatePassword(`'hashedPassword'`, `userPassword`)
*/
func ValidatePassword(hash, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))

	return err == nil
}

func Login() {
	// login logic will come here
}
