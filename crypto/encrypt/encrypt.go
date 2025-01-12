/*
Copyright © 2021 Sniptt <support@sniptt.com>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package encrypt

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
)

func Bytes(bytes []byte) ([]byte, []byte, error) {
	// Should be 16 bytes (AES-128), 24 bytes (AES-192) or 32 bytes (AES-256)
	key := make([]byte, 32)
	_, err := rand.Read(key)
	if err != nil {
		return nil, nil, err
	}

	// Generate a new aes cipher using the key above
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, nil, err
	}

	// gcm or Galois/Counter Mode, is a mode of operation
	// for symmetric key cryptographic block ciphers
	// - https://en.wikipedia.org/wiki/Galois/Counter_Mode
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, nil, err
	}

	// Create a new byte array the size of the nonce,
	// populate with a cryptographically secure random sequence
	nonce := make([]byte, gcm.NonceSize())
	_, err = rand.Read(nonce)
	if err != nil {
		return nil, nil, err
	}

	// Encrypt and authenticate plaintext
	ciphertext := gcm.Seal(nonce, nonce, bytes, nil)

	return ciphertext, key, nil
}
