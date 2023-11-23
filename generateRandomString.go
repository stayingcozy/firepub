package main

import (
	"math/rand"
	"time"
)

func generateRandomString(length int) string {
	// Seed
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	// Define the characters from which to generate the string
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	// Create a byte slice to store the random string
	randomString := make([]byte, length)

	// Populate the byte slice with random characters
	for i := 0; i < length; i++ {
		randomString[i] = charset[r.Intn(len(charset))]
	}

	// Convert the byte slice to a string
	return string(randomString)
}