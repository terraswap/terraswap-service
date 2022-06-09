package rdb

import (
	"testing"

	"github.com/delight-labs/terraswap-service/configs"
	"github.com/stretchr/testify/assert"
)

func TestFailedToLoadPostgreSQL(t *testing.T) {
	assert := assert.New(t)

	cfg := configs.RdbConfig{
		Host:     "1.1.1.1",
		Port:     "3306",
		Username: "1",
		Password: "2",
		Database: "test",
	}

	assert.Panics(func() {
		New(cfg)
	})
}
