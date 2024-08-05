package test

import "github.com/mateothegreat/go-multilog/multilog"

func Setup() {
	multilog.RegisterLogger(multilog.LogMethod("console"), multilog.NewConsoleLogger(&multilog.NewConsoleLoggerArgs{
		Level:  multilog.DEBUG,
		Format: multilog.FormatText,
	}))
}

func GetDockerHost() string {
	return "unix:///Users/matthewdavis/.docker/run/docker.sock"
}
