package test

import (
	"testing"

	"github.com/sirupsen/logrus"
)

func TestSingleton(t *testing.T) {
	logrus.Info("Hello Logrus!")
	logrus.Warn("Hello Logrus!")
	logrus.Error("Hello Logrus!")

	logrus.SetFormatter(&logrus.JSONFormatter{})

	logrus.Info("Hello Logrus!")
	logrus.Warn("Hello Logrus!")
	logrus.Error("Hello Logrus!")
}
