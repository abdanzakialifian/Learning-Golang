package test

import (
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func TestViper(t *testing.T) {
	var config = viper.New()
	assert.NotNil(t, config)
}

func TestJSON(t *testing.T) {
	config := viper.New()
	config.SetConfigName("config")
	config.SetConfigType("json")
	config.AddConfigPath("../")

	// read config
	err := config.ReadInConfig()
	assert.Nil(t, err)

	assert.Equal(t, "learning-viper-golang", config.GetString("app.name"))
	assert.Equal(t, "Abdan Zaki Alifian", config.GetString("app.author"))
	assert.Equal(t, "localhost", config.GetString("database.host"))
	assert.Equal(t, 3306, config.GetInt("database.port"))
	assert.True(t, config.GetBool("database.show_sql"))
}

func TestYAML(t *testing.T) {
	config := viper.New()
	config.SetConfigName("config")
	config.SetConfigType("yaml")
	config.AddConfigPath("../")

	// read config
	err := config.ReadInConfig()
	assert.Nil(t, err)

	assert.Equal(t, "learning-viper-golang", config.GetString("app.name"))
	assert.Equal(t, "Abdan Zaki Alifian", config.GetString("app.author"))
	assert.Equal(t, "localhost", config.GetString("database.host"))
	assert.Equal(t, 3306, config.GetInt("database.port"))
	assert.True(t, config.GetBool("database.show_sql"))
}

func TestENV(t *testing.T) {
	config := viper.New()
	config.SetConfigFile("config.env")
	config.AddConfigPath(".")

	// read config
	err := config.ReadInConfig()
	assert.Nil(t, err)

	assert.Equal(t, "learning-viper-golang", config.GetString("APP_NAME"))
	assert.Equal(t, "Abdan Zaki Alifian", config.GetString("APP_AUTHOR"))
	assert.Equal(t, "localhost", config.GetString("DATABASE_HOST"))
	assert.Equal(t, 3306, config.GetInt("DATABASE_PORT"))
	assert.True(t, config.GetBool("DATABASE_SHOW_SQL"))
}
