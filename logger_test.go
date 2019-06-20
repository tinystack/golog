package golog

import "testing"

func TestLogger(t *testing.T) {
	log := New(NewFileHandler("/tmp/test.log", ChannelDaily))
	log.Info("Hello, this is info log")
}
