package test

import (
	"learning-dependency-injection-golang/di"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSimpleDatabase(t *testing.T) {
	databaseRepository := di.InitializedDatabaseRepository()
	assert.Equal(t, databaseRepository.DatabasePostgreSQL.Name, "PostgreSQL")
	assert.Equal(t, databaseRepository.DatabaseMongoDB.Name, "MongoDB")
}
