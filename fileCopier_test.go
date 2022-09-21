package main

import (
	"fmt"
	"testing"
)

func TestParseWork(t *testing.T) {
	fds, auto := ParseTask()
	fmt.Printf("%+v %v", fds, auto)
}

func TestVerifyWork(t *testing.T) {
	fds, _ := ParseTask()
	for _, fd := range *fds {
		ok := VerifyTask(&fd)
		if !ok {
			fmt.Println(fd, false)
		}
	}
}

func TestProcessWork(t *testing.T) {
	fds, _ := ParseTask()
	for _, fd := range *fds {
		ProcessTask(&fd)
	}
}

func TestProcessTrash(t *testing.T) {
	fds, _ := ParseTask()
	for _, fd := range *fds {
		ProcessTrash(&fd)
	}
}
