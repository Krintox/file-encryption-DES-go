package main

import (
	"bytes"
	"crypto/cipher"
	"crypto/des"
	"encoding/base64"
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/encrypt", encryptHandler)
	http.ListenAndServe(":8080", addCorsHeader(http.DefaultServeMux))
}

func encryptHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Parse input file
	file, _, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Failed to parse file", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	// Generate a 24-byte key for Triple DES
	key := []byte("123456789012345678901234")

	// Verify the length of the key
	if len(key) != 24 {
		http.Error(w, "Invalid key length", http.StatusInternalServerError)
		return
	}

	// Create Triple DES cipher block
	block, err := des.NewTripleDESCipher(key)
	if err != nil {
		http.Error(w, "Failed to create cipher block", http.StatusInternalServerError)
		return
	}

	// Create a temporary buffer to hold the encrypted data
	var encryptedData []byte
	buf := make([]byte, block.BlockSize())

	// Encrypt the data from the input file
	mode := cipher.NewCBCEncrypter(block, key[:8])
	for {
		n, err := file.Read(buf)
		if err == io.EOF {
			break
		}
		if err != nil {
			http.Error(w, "Error reading file", http.StatusInternalServerError)
			return
		}
		// PKCS5Padding for the last block
		if n < len(buf) {
			buf = PKCS5Padding(buf[:n], block.BlockSize())
		}
		mode.CryptBlocks(buf, buf)
		encryptedData = append(encryptedData, buf...)
	}

	// Encode the encrypted data in base64 and send it as response
	encodedText := base64.StdEncoding.EncodeToString(encryptedData)
	w.Write([]byte(encodedText))
}

func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

// addCorsHeader adds necessary CORS headers to the given HTTP handler.
func addCorsHeader(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, Authorization")
		// Stop here if its Preflighted OPTIONS request
		if r.Method == "OPTIONS" {
			return
		}
		// Lets Gorilla work
		h.ServeHTTP(w, r)
	})
}
