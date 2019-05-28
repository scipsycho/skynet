package crypto

import (
	"errors"
	"strings"

	"crypto/rand"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"math/big"
)

type Crypto struct{}

var deliminator = "||"

func (c *Crypto) Generate(s string) (string, error) {
	salt := uuid.New().String()
	saltedBytes := []byte(s + salt)
	hashedBytes, err := bcrypt.GenerateFromPassword(saltedBytes, bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	hash := string(hashedBytes[:])
	return hash + deliminator + salt, nil
}

func (c *Crypto) Compare(hash string, s string) error {
	parts := strings.Split(hash, deliminator)
	if len(parts) != 2 {
		return errors.New("Invalid hash, must have 2 parts")
	}

	incoming := []byte(s + parts[1])
	existing := []byte(parts[0])
	return bcrypt.CompareHashAndPassword(existing, incoming)
}

func (c *Crypto) GenerateRandomASCIIString(length int) (string, error) {
	result := ""
	for {
		if len(result) >= length {
			return result, nil
		}
		num, err := rand.Int(rand.Reader, big.NewInt(int64(127)))
		if err != nil {
			return "", err
		}
		n := num.Int64()
		// Make sure that the number/byte/letter is inside
		// the range of printable ASCII characters (excluding space and DEL)
		if n > 32 && n < 127 {
			result += string(n)
		}
	}
}
