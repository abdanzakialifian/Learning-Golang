package test

import (
	"fmt"
	"testing"

	"github.com/sirupsen/logrus"
)

type SampleHook struct{}

func (hook *SampleHook) Levels() []logrus.Level {
	return []logrus.Level{logrus.WarnLevel, logrus.ErrorLevel}
}

func (hhok *SampleHook) Fire(entry *logrus.Entry) error {
	fmt.Println("Sample Hook", entry.Level, entry.Message)
	return nil
}

func TestHook(t *testing.T) {
	logger := logrus.New()
	logger.AddHook(&SampleHook{})

	logger.Info("This is Log Info")
	logger.Warn("This is Log Warning")
	logger.Error("This is Log Error")
}
