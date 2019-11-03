package apis

import (
	"github.com/chasdevs/meetrics/pkg/conf"
	"log"
	"path"
	"runtime"

	"google.golang.org/api/calendar/v3"

	"golang.org/x/net/context"
	"golang.org/x/oauth2/google"

	"io/ioutil"
	"net/http"

	"google.golang.org/api/admin/directory/v1"
)

// Calendar returns a client for the Google Calendar API. Provided email is the user who is being impersonated via the service account.
func Calendar(email string) calendar.Service {

	client := getClient(email)

	srv, err := calendar.New(client)
	if err != nil {
		log.Fatalf("Unable to retrieve calendar client: %v", err)
	}

	return *srv
}

// Admin returns a client for the Google Admin API
func Admin() admin.Service {
	client := getClient(conf.GetString("google.subject"))

	srv, err := admin.New(client)
	if err != nil {
		log.Fatalf("Unable to retrieve admin client: %v", err)
	}

	return *srv
}

func getClient(email string) *http.Client {
	return clientFromJwtConfig(email)
}

func clientFromJwtConfig(email string) *http.Client {

	// Your credentials should be obtained from the Google
	// Developer Console (https://console.developers.google.com).
	serviceAccountJSON, err := ioutil.ReadFile(serviceAccountFile())
	if err != nil {
		panic(err)
	}

	jwtConfig, _ := google.JWTConfigFromJSON(serviceAccountJSON, calendar.CalendarReadonlyScope, admin.AdminDirectoryUserReadonlyScope)
	jwtConfig.Subject = email
	return jwtConfig.Client(context.Background())
}


func serviceAccountFile() string {
	_, filename, _, _ := runtime.Caller(1)
	filepath := path.Join(path.Dir(filename), "../../config/service_account_key.json")
	return filepath
}