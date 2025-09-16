package main

import (
	"bufio"
	"crypto/rand"
	"crypto/sha256"
	"os"
	"slices"
	"strings"

	// "crypto/sha512"
	"encoding/base64"
	"fmt"
	"io"
)


func main(){
	enteredPassword := "password123"
	// hashedPass256Byte := sha256.Sum256([]byte(enteredPassword))
	// hashedPass512Byte := sha512.Sum512([]byte(enteredPassword))
	// hashedPassHex := fmt.Sprintf("%x\n",hashedPass256Byte)

	// fmt.Printf("Hexadecimal format: %s\n 256 Bytes format: %v\n 512 Bytes format: %v\n",hashedPassHex,hashedPass256Byte, hashedPass512Byte)
	salt , _ := generateSaltUsingBufio()
	fmt.Println("salt:", salt)
	saltInBase64 := base64.StdEncoding.EncodeToString(salt)
	fmt.Println("Salt in base64 format is:", saltInBase64)
	saltInBytes, _ := base64.StdEncoding.DecodeString(saltInBase64)
	fmt.Println("Salt in bytes format is:", saltInBytes)

	signUpHashedPassword := hashSaltedPassword(salt, enteredPassword)
	fmt.Println("The hashed salted Pass is:", signUpHashedPassword)

	// Verify
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter your enteredPassword:")
	loginPass, _ := reader.ReadString('\n')
	loginPass = strings.TrimSpace(loginPass)
	loginHashedPass := hashSaltedPassword(saltInBytes, loginPass)

	if slices.Equal(loginHashedPass, signUpHashedPassword) {
		fmt.Println("You logged in successfully!")
	} else {
		fmt.Println("The enteredPassword is incorrect! Try again!")
	}
}

func generateSaltUsingBufio () ([]byte, error) {
	reader := bufio.NewReader(rand.Reader)
	salt := make([]byte, 16)
	_, err := reader.Read(salt)
	if err != nil {
		fmt.Errorf("Error in reading :%v", err)
	}
	return salt, nil	
}

func generateSaltUsingIO () ([]byte, error) {
	salt := make([]byte, 16)
	_ , err := io.ReadFull(rand.Reader, salt)
	if err != nil {
		fmt.Errorf("Error in reading :%v", err)
	}
	return salt, nil	
}

func hashSaltedPassword (salt []byte, enteredPassword string) []byte {
	saltedPass := append(salt,[]byte(enteredPassword)...)
	hashedPassword := sha256.Sum256(saltedPass)
	return []byte(base64.StdEncoding.EncodeToString(hashedPassword[:]))
}