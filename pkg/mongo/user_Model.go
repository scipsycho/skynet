package mongo

import (
	root "skynet/pkg"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type userModel struct {
	Identifier string

	UserName     string
	PasswordHash string
	Salt         string
}

func newUserModel(u *root.User) (*userModel, error) {
	user := userModel{UserName: u.UserName}
	err := user.setSaltedPassword(u.Password)
	if err != nil {
		return nil, err
	}

	user.Identifier, err = generateUniqueIdentifier(user.UserName, user.PasswordHash, user.Salt)

	return &user, err
}

func (u *userModel) comparePassword(password string) error {
	incoming := []byte(password + u.Salt)
	existing := []byte(u.PasswordHash)
	err := bcrypt.CompareHashAndPassword(existing, incoming)

	return err
}

func (u *userModel) setSaltedPassword(password string) error {
	salt := uuid.New().String()
	passwordBytes := []byte(password + salt)
	hash, err := bcrypt.GenerateFromPassword(passwordBytes, bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	u.PasswordHash = string(hash[:])
	u.Salt = salt

	return nil
}

func generateUniqueIdentifier(username, passwordHash, salt string) (string, error) {
	randomIdentifier := []byte(username + passwordHash + salt)
	hash, err := bcrypt.GenerateFromPassword(randomIdentifier, bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hash[:]), nil
}
