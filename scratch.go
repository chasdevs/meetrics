package main

import (
	"github.com/chasdevs/meetrics/pkg/apis"
	"github.com/chasdevs/meetrics/pkg/conf"
	"github.com/chasdevs/meetrics/pkg/data"
	"github.com/chasdevs/meetrics/pkg/metrics"
	log "github.com/sirupsen/logrus"
	"strings"
	"time"
)

func init() {
	log.SetLevel(log.DebugLevel)
	data.Init()
}

func main() {
	eventChan := make(chan<- metrics.UserEvent)
	user := data.Mgr.GetUserById(82)
	metrics.CompileMetricsForUser(time.Date(2017, 11, 7, 0, 0, 0, 0, time.UTC), user, eventChan)
}

func searchUsers(name string) {
	adminApi := apis.Admin()
	res, err := adminApi.Users.List().Query("name:" + name).Domain(conf.GetString("google.domain")).Do()
	if err != nil {
		log.Fatalf("Error fetching users for organization: %v", err)
	}

	emails := make([]string, len(res.Users))
	for i, user := range res.Users {
		emails[i] = user.PrimaryEmail
	}
	log.Info("Matching emails: " + strings.Join(emails, ","))

}
