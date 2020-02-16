package data

import (
	log "github.com/sirupsen/logrus"
	"time"
)

// Data

func (mgr *manager) CreateUserMeetingMins(date time.Time, user User, meetingMins map[string]uint) {
	ormObj := UserMeetingMins{
		Date:      date,
		UserID:    user.ID,
		Mins0:     meetingMins["mins0"],
		Mins1:     meetingMins["mins1"],
		Mins2Plus: meetingMins["mins2Plus"],
		Cranktime: 	   meetingMins["crank"],
		Deadtime: 	   meetingMins["dead"],
	}

	errors := mgr.db.Create(&ormObj).GetErrors()
	if len(errors) > 0 {
		log.Error("Error adding user meeting mins: %v", errors[0])
	}
}

func (mgr *manager) ClearUserMeetingMins() {
	errors := mgr.db.Delete(UserMeetingMins{}).GetErrors()
	if len(errors) > 0 {
		log.Fatalf("Error clearing user meeting mins: %v", errors[0])
	}
}

type UserMeetingMins struct {
	Date      time.Time `gorm:"primary_key" sql:"type:date"`
	UserID    uint      `gorm:"primary_key" sql:"type:int unsigned"`
	Mins0     uint
	Mins1     uint
	Mins2Plus uint
	Cranktime uint
	Deadtime  uint
	CreatedAt time.Time
}
