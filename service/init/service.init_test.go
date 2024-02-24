package service

import (
	"testing"
)

func TestDoWork(t *testing.T) {
	serv := NewInitService("abc")

	var done = make(chan struct{})
	go DoWork(serv, done)

	<-done
	//time.Sleep(time.Second * 20)
}
