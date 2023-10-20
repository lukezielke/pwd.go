package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"
)

func main() {
	//all characters that can be in the password
	var symbols = [81]string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z", "A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z", "~", ":", "+", "[", "@", "^", "{", "%", "(", "-", "*", "|", ",", "&", "<", "`", "}", ".", "_", "=", "]", "!", ">", ";", "?", "#", "$", ")", "/"}
	passwords := []string{}
	password := ""
	//seed for the random function
	r := rand.New(rand.NewSource(time.Now().Unix()))

	//optional command line arguments for length, amount and output file
	lengthPtr := flag.Int("length", 12, "[OPTIONAL] Desired length of the password")
	amountPtr := flag.Int("amount", 1, "[OPTIONAL] Desired amount of passwords to be generated")
	outputPtr := flag.String("output", "", "[OPTIONAL] Path of a file to write the passwords to")

	flag.Parse()

	//desired amount of passwords with specified length is generated here
	for i := 0; i < *amountPtr; i++ {
		password = ""
		for j := 0; j < *lengthPtr; j++ {
			password = password + symbols[r.Intn(len(symbols))]
		}
		passwords = append(passwords, password)
	}

	//check if the passwords should be written to an output file
	if *outputPtr != "" {
		f, err := os.OpenFile(*outputPtr, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

		if err != nil {
			log.Fatal(err)
		} else {
			for _, line := range passwords {
				_, err := f.WriteString(line + "\n")
				if err != nil {
					log.Fatal(err)
				}
			}

		}
		defer f.Close()

	} else {
		for _, line := range passwords {
			fmt.Println(line)
		}
	}
}
