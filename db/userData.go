package db

import (
	"encoding/json"
	"log"
	"os"

	"github.com/abhijitWakchaure/mock-api-server-go/mylogger"
	"github.com/abhijitWakchaure/mock-api-server-go/user"
	"github.com/gobuffalo/packr/v2"
)

// MockUserData ...
var MockUserData []user.User

func init() {
	mylogger.InfoLog("Initializing database...")
	box := packr.New("myBox", ".")
	userBytes, err := box.Find("userData.json")
	if err != nil {
		mylogger.ErrorLog("Unable to open user data from json file: ", err.Error())
		os.Exit(1)
	}
	// jsonFile, err := os.Open("userData.json")
	// if err != nil {
	// 	mylogger.ErrorLog("Unable to open user data from json file: ", err.Error())
	//   os.Exit(1)
	// }
	// userBytes, _ := ioutil.ReadAll(jsonFile)
	err = json.Unmarshal([]byte(userBytes), &MockUserData)
	if err != nil {
		log.Fatalf("Unable to initialize users data: %v", err)
	}
	mylogger.InfoLog("Loaded %d users in the database", len(MockUserData))
}
