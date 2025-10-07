package crypto

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"github.com/go-jose/go-jose/v4"
)

var key []byte

func init() {
	keyStr := os.Getenv("ENCRYPTION_KEY")
	if keyStr == "" {
		panic("ENCRYPTION_KEY is not set")
	}
	key = []byte(keyStr)

	// AES accepts key sizes: 16, 24, or 32 bytes
	if len(key) != 16 && len(key) != 24 && len(key) != 32 {
		panic("ENCRYPTION_KEY must be 16, 24, or 32 bytes long")
	}
}

func EncryptJSON(data interface{}) (string, error) {
	jsonBytes, err := json.Marshal(data)
	if err != nil {
		return "", err
	}

	return Encrypt(string(jsonBytes))
}

func DecryptJSON(ciphertext string, v interface{}) error {
	plaintext, err := Decrypt(ciphertext)
	if err != nil {
		return err
	}

	return json.Unmarshal([]byte(plaintext), v)
}

// Encrypt performs AES-GCM encryption and returns base64 ciphertext
func Encrypt(plaintext string) (string, error) {
	// Algorithm: direct key (dir) + AES256-GCM for content encryption (A256GCM)
	var enc jose.ContentEncryption

	switch len(key) {
	case 16:
		enc = jose.A128GCM
	case 24:
		enc = jose.A192GCM
	case 32:
		enc = jose.A256GCM
	default:
		return "", errors.New("invalid key length for encryption")
	}

	encrypter, err := jose.NewEncrypter(enc, jose.Recipient{
		Algorithm: jose.DIRECT,
		Key:       key,
	}, nil)

	if err != nil {
		return "", err
	}

	jweObj, err := encrypter.Encrypt([]byte(plaintext))
	if err != nil {
		return "", err
	}

	return jweObj.CompactSerialize()
}

// Decrypt performs AES-GCM decryption and returns the plaintext string
func Decrypt(jweString string) (string, error) {
	fmt.Println("Decrypting JWE:", jweString) // Debugging line
	var encryptedString interface{}
	if err := json.Unmarshal([]byte(jweString), &encryptedString); err == nil {
		return "", errors.New("input appears to be JSON, expected JWE compact serialization")
	}

	jweObj, err := jose.ParseEncrypted(jweString, []jose.KeyAlgorithm{jose.DIRECT},
		[]jose.ContentEncryption{jose.A256GCM})
	if err != nil {
		return "", err
	}

	decrypted, err := jweObj.Decrypt(key)
	if err != nil {
		return "", err
	}

	return string(decrypted), nil
}
