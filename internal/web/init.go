package web

import (
	"github.com/marek5050/kube-elk/internal/httplog"
	"github.com/sirupsen/logrus"
)

var log logrus.Logger

func init() {
	cfg := httplog.Config{
		MinLevel: logrus.InfoLevel,
	}
	h := httplog.NewHook(cfg, "http://org1.log.example.com/key/1233")
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetLevel(logrus.InfoLevel)
	logrus.AddHook(h)

}
