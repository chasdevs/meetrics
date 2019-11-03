package setup

import (
	"github.com/chasdevs/meetrics/pkg/apis"
	"github.com/chasdevs/meetrics/pkg/conf"
	"github.com/chasdevs/meetrics/pkg/data"
	"github.com/chasdevs/meetrics/pkg/setup/userdepartments"
	"log"
)

func PopulateUsersFromApi() {
	// fetch all users from the google org and store in users table
	// - Filter emails somehow? Blacklist email addresses which are not people?
	// - Do not overwrite emails if they already exist (keep the same id)
	adminApi := apis.Admin()
	res, err := adminApi.Users.List().Domain(conf.GetString("google.domain")).Do()
	if err != nil {
		log.Fatalf("Error fetching users for organization: %v", err)
	}

	users := make([]*data.User, len(res.Users))
	for i, user := range res.Users {
		users[i] = &data.User{Email: user.PrimaryEmail, Name: user.Name.FullName}
	}

	addUsers(users)
}

func PopulateUsersFromCsv() {
	users := userdepartments.GetUsersFromFile()
	addUsers(users)
}

func addUsers(users []*data.User) {
	data.Init()
	data.Mgr.AddAllUsers(users)
}
