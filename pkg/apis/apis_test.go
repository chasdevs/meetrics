package apis_test

import (
	"github.com/chasdevs/meetrics/pkg/apis"
	"github.com/chasdevs/meetrics/pkg/conf"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestApis(t *testing.T) {
	adminApi := apis.Admin()
	domain := conf.GetString("google.domain")
	assert.Equal(t, "chas", domain)
	res, err := adminApi.Users.List().Domain(domain).Do()
	assert.Nil(t, err)
	assert.NotNil(t, res)
}