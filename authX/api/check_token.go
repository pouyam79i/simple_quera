package api

import (
	"errors"
	"fmt"
	"strings"

	mysql "github.com/pouyam79i/simple_quera/authX/db"
)

func CheckToken(token string) (bool, error) {

	db, err := mysql.DB()
	if err != nil {
		return false, err
	}

	// TODO: handle errors properly
	result, err := db.Query("SELECT token FROM "+mysql.Table()+" WHERE TOKEN = ?", token)
	// Handle errors properly
	if err != nil {
		fmt.Println("DB error: ", err.Error())
		return false, err
	}

	var userToken string = ""

	if result.Next() {
		err = result.Scan(&userToken)
		if err != nil {
			fmt.Println("DB error: ", err.Error())
			return false, err
		} else {
			fmt.Println("Received data: ", userToken)
		}
	} else {
		fmt.Println("No user found with given information. Trying to create new user and sign in him")
		return false, errors.New("invalid token")
	}

	if strings.Compare(userToken, token) != 0 {
		return false, errors.New("server side error! failed to query to db properly")
	}

	return true, nil

}
