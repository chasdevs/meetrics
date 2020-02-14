package util_test

import (
	"github.com/chasdevs/meetrics/pkg/data"
	"github.com/chasdevs/meetrics/pkg/metrics"
	"github.com/chasdevs/meetrics/pkg/util"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestDisk(t *testing.T) {

	data.Init()
	user := data.Mgr.GetUserById(16)

	events := metrics.GetEventsForUser(time.Now(), user)

 	err := util.SaveEvents(events)
 	assert.Nil(t, err)
}
