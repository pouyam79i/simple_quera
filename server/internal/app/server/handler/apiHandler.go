package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/pouyam79i/simple_quera/server/internal/app/server/config"
	"github.com/pouyam79i/simple_quera/server/internal/app/server/handler/api"
)

// This is a simple test handler!
func HelloWorld(c echo.Context) error {
	return c.String(http.StatusOK, "Hello World from MY SERVER!!!!")
}

// TODO: complete sign in - make user login!
func SignIn(c echo.Context) error {
	c.Response().Header().Set("Access-Control-Allow-Origin", "*")

	data := config.SingInInfo{}
	err := json.NewDecoder(c.Request().Body).Decode(&data)
	if err != nil {
		resErr := &config.ClientMSG{
			Result: false,
			Token:  "",
			Info:   "Error While Parsing Json File:\n" + err.Error(),
		}
		return c.JSON(http.StatusBadRequest, resErr)
	} else {
		fmt.Println("Received Data", "\nEmail: ", data.Email, "\nPassword: ", data.Password)
	}

	// TODO: connect to authx
	token, err := api.Authx_SingIn(data)
	res := config.ClientMSG{}
	code := http.StatusOK
	if err != nil {
		res.Result = false
		res.Token = ""
		res.Info = err.Error()
		code = http.StatusForbidden
	} else {
		res.Result = true
		res.Token = token
		res.Info = "Successful"
	}

	return c.JSON(code, res)
}

func TestValidator(c echo.Context) error {
	data := config.JustToken{}
	err := json.NewDecoder(c.Request().Body).Decode(&data)
	if err != nil {
		resErr := &config.ClientMSG{
			Result: false,
			Token:  "",
			Info:   "Error While Parsing Json File:\n" + err.Error(),
		}
		return c.JSON(http.StatusBadRequest, resErr)
	} else {
		fmt.Println("Received Data from user", "\nGiven Token: ", data.Token)
	}

	isValid, err := api.Authx_Validate(data.Token)
	info := "successful"

	if err != nil {
		info = "Err: " + err.Error()
	}

	resErr := &config.ClientMSG{
		Result: isValid,
		Token:  "true means given token was valid",
		Info:   info,
	}
	return c.JSON(http.StatusOK, resErr)
}

// TODO: complete code uploading mechanism - test version
func UploadCode(c echo.Context) error {
	reqBody := config.ClientCode{}
	err := c.Bind(&reqBody)
	if err != nil {
		res := config.ClientMSG{
			Result: false,
			Info:   err.Error(),
		}
		return c.JSON(http.StatusBadRequest, res)
	}
	// TODO: save the file and tell the job scheduler!
	// api.SendCodeX(reqBody)
	api.SendDataToJobBuilder(reqBody)
	res := config.ClientMSG{
		Result: true,
		Info:   "We are sending your code to codex api. Please be patient, the result of your code will be sent to your email",
	}
	return c.JSON(http.StatusOK, res)
}

func Upload(c echo.Context) error {
	//------------
	// Read files
	//------------

	token := c.Request().Header.Get("X-PM-TOKEN")
	isValid, err := api.Authx_Validate(token)
	if err != nil {
		fmt.Println("Error while reading file: ", err.Error())
		res := config.ClientMSG{
			Result: false,
			Info:   "Authentication failed with Error",
		}
		return c.JSON(http.StatusUnauthorized, res)
	}
	if !isValid {
		fmt.Println("Unauthorized token")
		res := config.ClientMSG{
			Result: false,
			Info:   "Invalid Token",
		}
		return c.JSON(http.StatusUnauthorized, res)
	}

	// Multipart form
	form, err := c.MultipartForm()
	if err != nil {
		fmt.Println("Error while reading file: ", err.Error())
		res := config.ClientMSG{
			Result: false,
			Info:   "cannot read files",
		}
		return c.JSON(http.StatusBadRequest, res)
	}
	files := form.File["files"]

	// TODO: complete s3 uploader - cannot complete because downloading libs is forbidden
	// sess, err := session.NewSession(&aws.Config{
	// 	Credentials: credentials.NewStaticCredentials("<ACCESS_KEY>", "<SECRET_KEY>", ""),
	// 	Region:      aws.String("default"),
	// 	Endpoint:    aws.String("<ENDPOINT_URL>"),
	// })
	// uploader := s3manager.NewUploader(sess)
	// _, err = uploader.Upload(&s3manager.UploadInput{
	// 	Bucket: aws.String(bucket),

	// 	// Can also use the `filepath` standard library package to modify the
	// 	// filename as need for an S3 object key. Such as turning absolute path
	// 	// to a relative path.
	// 	Key: aws.String(filename),

	// 	// The file to be uploaded. io.ReadSeeker is preferred as the Uploader
	// 	// will be able to optimize memory when uploading large content. io.Reader
	// 	// is supported, but will require buffering of the reader's bytes for
	// 	// each part.
	// 	Body: file,
	// })

	for _, file := range files {
		fmt.Println("Received File name: ", file.Filename)
	}

	res := config.ClientMSG{
		Result: true,
		Info:   "Files Received Successfully",
	}

	// TODO: build a job and inform job builder

	return c.JSON(http.StatusOK, res)
}
