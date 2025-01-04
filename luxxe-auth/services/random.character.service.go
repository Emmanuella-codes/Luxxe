package services

import (
	"math/rand"
	"time"
)

var CapitalLettersAlphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
var SmallLettersAlphabet = "abcdefghijklmnopqrstuvwxyz"
var NumbersAlphabet = "0123456789"
var SpecialCharactersAlphabet = "~!@#$%^&*()_+=/.,<>;:"

func GetRandomCharacters(count int, alphabet ...string) string {
	// Seed the random number generator
	seed := time.Now().UnixNano()
	rng := rand.New(rand.NewSource(seed))

	var charSet string
	if len(alphabet) == 0 {
		// Define the character set
		charSet = CapitalLettersAlphabet + SmallLettersAlphabet + NumbersAlphabet
	} else {
		charSet = alphabet[0]
	}

	// Generate the random string
	randomString := make([]byte, count)
	for i := range randomString {
		randomString[i] = charSet[rng.Intn(len(charSet))]
	}
	return string(randomString)
}
