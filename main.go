package main

import (
	"fmt"
	"flag"
	"crypto/rand"
	"math/big"
	"strings"
)

const (
	letters      = "ABCDEFGHIJKLMNOPQRSTUVWXYabcdefghijklmnopqrstuvwxyz"
    numbers      = "0123456789"
    specialChars = "!@#$%^&*()-_=+[]{}|;:',.<>?/`~"
)

func main() {

	passLength := flag.Int("length", 12, "Length of the password")
	includeNumbers := flag.Bool("numbers", true, "Include numbers in the password")
	includeSpecial := flag.Bool("special", true, "Include special characters in the password")

	flag.Parse()

	fmt.Println("Generating password with the following settings:")
	fmt.Printf("Length: %d\n", *passLength)
	fmt.Printf("Include Numbers: %t\n", *includeNumbers)
	fmt.Printf("Include Special Characters: %t\n", *includeSpecial)

	var password strings.Builder
	password.Grow(*passLength)

	halfLength := *passLength / 2

	// alternate between character, number, and special characters 
	// first index is character, second is number, third is special character, and so on
	for i := 0; i < *passLength; i++ {

		var char byte
		var err error

		if(i%3 == 0) {
			char, err = getRandomChar(letters)
		} else if(i%3 == 1 && *includeNumbers && i < halfLength) {
			char, err = getRandomNumbers()
		} else if(i%3 == 2 && *includeSpecial && i < halfLength) {
			char, err = getRandomSpecialChar()
		} else  {
			char, err = getRandomChar(letters)
		}

		if err != nil {
			fmt.Printf("Error generating character: %v\n", err)
			password.WriteByte(letters[0])
			continue
		}

		password.WriteByte(char)
	}

	fmt.Printf("Generated Password: %s\n", password.String())  // Add .String()

}

func getRandomChar(charset string) (byte, error) {
	randomNum, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
	return charset[randomNum.Int64()], err
}


func getRandomNumbers() (byte, error) {
	randomNum, err := rand.Int(rand.Reader, big.NewInt(int64(len(numbers))))
	return numbers[randomNum.Int64()], err
} 

func getRandomSpecialChar() (byte, error) {
	randomNum, err := rand.Int(rand.Reader, big.NewInt(int64(len(specialChars))))
	return specialChars[randomNum.Int64()], err
} 