package api

import (
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"fmt"
	"strings"

	"github.com/pouyam79i/simple_quera/authX/config"
	mysql "github.com/pouyam79i/simple_quera/authX/db"
	"github.com/pouyam79i/simple_quera/authX/syserror"
)

func CreateNewUser(newUser config.API_SIGNIN_REQUEST) (string, error) {

	if len(newUser.Email) > 511 || len(newUser.Password) > 511 {
		return "", errors.New(syserror.LongData)
	}

	newUser.Email = strings.ToLower(newUser.Email)

	if !strings.HasSuffix(newUser.Email, "@gmail.com") || len(newUser.Email) <= len("@gmail.com") {
		return "", errors.New(syserror.InvalidEmail)
	}

	// TODO: also avoid invalid character.
	if len(newUser.Password) < 8 || strings.ContainsAny(newUser.Password, "\n "){
		return "", errors.New(syserror.ShortPassword)
	}

	db, err := mysql.DB()

	if err != nil {
		return "", err
	}

	hash := sha256.New()
	hash.Write(hash.Sum([]byte(newUser.Email + "+" + newUser.Password)))
	token := base64.URLEncoding.EncodeToString(hash.Sum(nil))
	
	// Injecting new user
	q, err := db.Query("INSERT INTO "+mysql.Table()+" VALUES (?, ?, ?)", newUser.Email, token, newUser.Password)

	if err != nil {
		fmt.Println("Unexpected Error: ", err.Error())
		return "", err
	}

	q.Close()

	fmt.Println("New user created:", "\nEmail: ", newUser.Email, "\nPassword: ", newUser.Password, "\nToken: ", token)

	return token, nil
}
