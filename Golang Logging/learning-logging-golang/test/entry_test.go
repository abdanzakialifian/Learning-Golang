package test

import (
	"testing"

	"github.com/sirupsen/logrus"
)

func TestEntry(t *testing.T) {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})

	entry := logrus.NewEntry(logger)
	entry.Info("Hello Logger With Manual Entry!")
	entry.WithField("username", "abdanzakialifian").Info("Hello Logger With Manual Entry!")
}
