package test

import (
	"testing"

	"github.com/sirupsen/logrus"
)

func TestLevel(t *testing.T) {
	logger := logrus.New()
	logger.Trace("This is Log Trace")
	logger.Debug("This is Log Debug")
	logger.Info("This is Log Info")
	logger.Warn("This is Log Warning")
	logger.Error("This is Log Error")
}

func TestChangeDefaultLogLevel(t *testing.T) {
	logger := logrus.New()
	logger.SetLevel(logrus.TraceLevel)
	logger.Trace("This is Log Trace")
	logger.Debug("This is Log Debug")
	logger.Info("This is Log Info")
	logger.Warn("This is Log Warning")
	logger.Error("This is Log Error")
}
