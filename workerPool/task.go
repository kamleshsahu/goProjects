package main

import (
	"fmt"
	"time"
)

type Task interface {
	Process()
	//GetId() int
}

type EmailProcessor struct {
	email, message string
}

func (e *EmailProcessor) Process() {
	fmt.Println("Message sent to ", e.email, " is : ", e.message)
	time.Sleep(3 * time.Second)
}

type MeetingScheduler struct {
	email string
}

func (m *MeetingScheduler) Process() {
	fmt.Println("Meeting scheduled with ", m.email)
	time.Sleep(3 * time.Second)
}
