package data

import (
	"log"
	"time"
)

//meetings: id, name, description, frequency_per_month, minutes, start_date, end_date
//meeting_users: meeting_id, user_id

type Meeting struct {
	ID                string `gorm:"primary_key"`
	Name              string
	Description       string `sql:"type:text"`
	Attendees         uint8
	Mins              uint8
	FrequencyPerMonth uint8
	StartDate         time.Time
	EndDate           time.Time
	CreatedAt         time.Time
	UpdatedAt         time.Time
}

func (mgr *manager) CreateMeetings(meetings []*Meeting) {
	errors := mgr.db.Create(meetings).GetErrors()
	if len(errors) > 0 {
		log.Fatalf("Error creating meeting: %v", errors[0])
	}
}

func (mgr *manager) CreateMeeting(meeting *Meeting) {
	errors := mgr.db.Set("gorm:insert_option", "ON DUPLICATE KEY UPDATE id=id").Create(meeting).GetErrors()
	if len(errors) > 0 {
		log.Fatalf("Error creating meeting: %v", errors[0])
	}
}

func (mgr *manager) ClearMeetings() {
	errors := mgr.db.Delete(Meeting{}).GetErrors()
	if len(errors) > 0 {
		log.Fatalf("Error clearing meetings: %v", errors[0])
	}
}
