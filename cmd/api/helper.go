package main

import (
	"crypto/rand"
)

func generateRandomBytes(n int) ([]byte, error) {
	randBytes := make([]byte, n)
	_, err := rand.Read(randBytes)
	if err != nil {
		return nil, err
	}
	return randBytes, nil
}