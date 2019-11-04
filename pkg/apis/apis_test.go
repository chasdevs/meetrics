package apis_test

import (
	"github.com/chasdevs/meetrics/pkg/apis"
	"github.com/chasdevs/meetrics/pkg/conf"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAdminApi(t *testing.T) {
	adminApi := apis.Admin()
	res, err := adminApi.Users.List().Domain(conf.GetString("google.domain")).Do()
	assert.Nil(t, err)
	assert.NotNil(t, res)
	assert.GreaterOrEqual(t, len(res.Users), 1)
}

func TestCalendarApi(t *testing.T) {
	subject := conf.GetString("google.subject")
	calendarApi := apis.Calendar(subject)
	res, err := calendarApi.Events.List(subject).Do()
	assert.Nil(t, err)
	assert.NotNil(t, res)
	assert.GreaterOrEqual(t, len(res.Items), 1)
}