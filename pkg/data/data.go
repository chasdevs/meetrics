package data

import (
	"github.com/chasdevs/meetrics/pkg/conf"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

type Manager interface {
	CreateMeeting(meeting *Meeting)
	CreateMeetings(meetings []*Meeting)
	ClearMeetings()

	CreateUserMeetingMins(date time.Time, user User, meetingMins map[string]uint)
	ClearUserMeetingMins()

	AddAllUsers(users []*User)
	GetAllUsers() []User
	GetUserByEmail(email string) User
	GetUserById(id int) User
}

type manager struct {
	db *gorm.DB
}

var Mgr Manager

// When lowercase, init() will run during package initialization
func Init() {
	config := conf.MysqlConfig()
	db, err := gorm.Open("mysql", config.FormatDSN())
	if err != nil {
		panic("failed to connect database")
	}
	Mgr = &manager{db: db}
}
