package services

import "golang.org/x/crypto/bcrypt"

// NormalizeStringHash func for a returning the users input as a byte slice.
func NormalizeStringHash(p string) []byte {
	return []byte(p)
}

func GenerateStringHash(p string) string {
	// Normalize password from string to []byte.
	bytePwd := NormalizeStringHash(p)

	// MinCost is just an integer constant provided by the bcrypt package
	// along with DefaultCost & MaxCost. The cost can be any value
	// you want provided it isn't lower than the MinCost (4).
	hash, err := bcrypt.GenerateFromPassword(bytePwd, bcrypt.MinCost)
	if err != nil {
		return err.Error()
	}

	// GenerateFromPassword returns a byte slice so we need to
	// convert the bytes to a string and return it.
	return string(hash)
}

// CompareStringHash func for a comparing password.
func CompareStringHash(hashedStr, inputStr string) bool {
	// Since we'll be getting the hashed password from the DB it will be a string,
	// so we'll need to convert it to a byte slice.
	byteHash := NormalizeStringHash(hashedStr)
	byteInput := NormalizeStringHash(inputStr)

	// Return result.
	if err := bcrypt.CompareHashAndPassword(byteHash, byteInput); err != nil {
		return false
	}

	return true
}
