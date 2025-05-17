package test

import (
	"learning-dependency-injection-golang/di"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConnection(t *testing.T) {
	connection, cleanup := di.InitializedConnection("Database")
	assert.NotNil(t, connection)
	cleanup()
}
