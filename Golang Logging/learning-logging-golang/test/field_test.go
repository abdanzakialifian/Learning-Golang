package test

import (
	"os"
	"testing"

	"github.com/sirupsen/logrus"
)

func TestField(t *testing.T) {
	logger := logrus.New()

	file, err := os.OpenFile("application3.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		panic(err)
	}
	logger.SetOutput(file)

	logger.WithField("username", "abdanzakialifian").Info("This is Log Info")
	logger.WithField("username", "abdanzakialifian").Warn("This is Log Warning")
	logger.WithField("username", "abdanzakialifian").Error("This is Log Error")

	logger.WithField("username", "abdanzakialifian").WithField("name", "Abdan Zaki Alifian").Info("This is Log Info")
	logger.WithField("username", "abdanzakialifian").WithField("name", "Abdan Zaki Alifian").Warn("This is Log Warning")
	logger.WithField("username", "abdanzakialifian").WithField("name", "Abdan Zaki").Error("This is Log Error")
}

func TestFields(t *testing.T) {
	logger := logrus.New()

	file, err := os.OpenFile("application4.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		panic(err)
	}
	logger.SetOutput(file)

	logger.WithField("username", "abdanzakialifian").Info("This is Log Info")
	logger.WithField("username", "abdanzakialifian").Warn("This is Log Warning")
	logger.WithField("username", "abdanzakialifian").Error("This is Log Error")

	withFields := logrus.Fields{
		"username": "abdanzakialifian",
		"name":     "Abdan Zaki Alifian",
	}

	logger.WithFields(withFields).Info("This is Log Info")
	logger.WithFields(withFields).Warn("This is Log Warning")
	logger.WithFields(withFields).Error("This is Log Error")
}
