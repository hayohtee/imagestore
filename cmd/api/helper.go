package main

import (
	"crypto/rand"
	"errors"
	"fmt"
	"mime"
	"path/filepath"
	"strings"
)

func generateRandomString(n int) (string, error) {
	randBytes := make([]byte, n)
	_, err := rand.Read(randBytes)
	if err != nil {
		return "", err
	}
	// Convert bytes to an hexadecimal string
	return fmt.Sprintf("%x", randBytes), nil
}

func getFileExtension(filename, contentType string) string {
	// If contentType is available, derive extension from it
	if contentType != "" {
		exts, err := mime.ExtensionsByType(contentType)
		if err == nil && len(exts) > 0 {
			return exts[0]
		}
	}
	// If the contentType is not specified, fallback to using the filename
	return strings.ToLower(filepath.Ext(filename))
}

func generateUniqueFilename(filename, contentType string) (string, error) {
	randomString, err := generateRandomString(16)
	if err != nil {
		return "", err
	}

	fileExtension := getFileExtension(filename, contentType)
	if fileExtension == "" {
		return "", errors.New("unable to generate file extension")
	}

	uniqueFileName := randomString + fileExtension
	return uniqueFileName, nil
}