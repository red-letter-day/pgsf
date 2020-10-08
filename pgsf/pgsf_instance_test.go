package pgsf_test

import (
	"github.com/red-letter-day/pgsf/pgsf"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewPgsfInstance(t *testing.T) {

	configuration := pgsf.Configuration{"testServer", 10, 8083, "/game"}
	instance := pgsf.NewPgsfInstance(configuration)

	assert.Equal(t, configuration, instance.GetConfiguration(), "expected configuration equality")
}
