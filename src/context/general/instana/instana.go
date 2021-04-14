package instana

import (
	instana "github.com/instana/go-sensor"
	"github.com/sirupsen/logrus"
)

type Instana struct {
	Logger *logrus.Logger
}

func NewInstana(logger *logrus.Logger) Instana {
	return Instana{
		Logger: logger,
	}
}

func (it *Instana) InitMetrics() {
	instana.InitSensor(&instana.Options{
		EnableAutoProfile: true,
	})
	instana.SetLogger(it.Logger)
}
