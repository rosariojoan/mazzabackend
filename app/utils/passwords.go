package utils

import "golang.org/x/crypto/bcrypt"

/* Convert a value to a password hash. Return an empty string if an error occurs. */
func HashPwd(value *string) (string, error) {
	pwdHash, err := bcrypt.GenerateFromPassword([]byte(*value), 10)
	if err != nil {
		return "", nil
	}
	return string(pwdHash), nil
}

/* Check if the readable password is equal to the hash password. Return error is false. */
func CompareHashAndPassword(pwdHash string, pwdStr string) error {
	return bcrypt.CompareHashAndPassword([]byte(pwdHash), []byte(pwdStr))
}