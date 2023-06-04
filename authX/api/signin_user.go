package api

import (
	"errors"
	"fmt"
	"strings"

	"github.com/pouyam79i/simple_quera/authX/config"
	mysql "github.com/pouyam79i/simple_quera/authX/db"
	"github.com/pouyam79i/simple_quera/authX/syserror"
)

func SingInUser(userInfo config.API_SIGNIN_REQUEST) (string, error) {

	if len(userInfo.Email) > 511 || len(userInfo.Password) > 511 {
		return "", errors.New(syserror.LongData)
	}

	userInfo.Email = strings.ToLower(userInfo.Email)

	if !strings.HasSuffix(userInfo.Email, "@gmail.com") || len(userInfo.Email) <= len("@gmail.com") {
		return "", errors.New(syserror.InvalidEmail)
	}

	db, err := mysql.DB()
	if err != nil {
		return "", err
	}

	// TODO: handle new user creation here!
	result, err := db.Query("SELECT * FROM "+mysql.Table()+" WHERE EMAIL = ?", userInfo.Email)
	// Handle errors properly
	if err != nil {
		fmt.Println("DB error: ", err.Error())
		return "", err
	}
	userData := config.USER_DATA{}
	if result.Next() {
		err = result.Scan(&userData.EMAIL, &userData.TOKEN, &userData.PWD)
		if err != nil {
			fmt.Println("DB error: ", err.Error())
			return "", err
		} else {
			fmt.Println("Received data: ", userData)
		}
	} else {
		fmt.Println("No user found with given information. Trying to create new user and sign in him")
		return CreateNewUser(userInfo)
	}

	if strings.Compare(userData.PWD, userInfo.Password) != 0 {
		return "", errors.New(syserror.InvalidPassword)
	}

	if strings.Compare(userData.EMAIL, userInfo.Email) == 0 {
		return userData.TOKEN, nil
	}

	return "", nil

}
