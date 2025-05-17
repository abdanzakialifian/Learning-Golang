package test

import (
	"learning-dependency-injection-golang/di"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSimpleError(t *testing.T) {
	simpleService, err := di.InitializedService(true)
	assert.Nil(t, simpleService)
	assert.Equal(t, "failed create service", err.Error())
}

func TestSimpleNotError(t *testing.T) {
	simpleService, err := di.InitializedService(false)
	assert.Nil(t, err)
	assert.NotNil(t, simpleService)
}
