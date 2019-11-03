package apis

import (
	"github.com/chasdevs/meetrics/pkg/conf"
	"log"

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
	client := getClient(conf.GetString("google.adminEmail"))

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
	serviceAccountJSONFile := "config/service_account_key.json"
	serviceAccountJSON, _ := ioutil.ReadFile(serviceAccountJSONFile)
	conf, _ := google.JWTConfigFromJSON(serviceAccountJSON, calendar.CalendarReadonlyScope, admin.AdminDirectoryUserReadonlyScope)
	conf.Subject = email

	return conf.Client(context.Background())
}
