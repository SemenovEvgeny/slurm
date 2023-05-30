package main

import (
	"fmt"
)

type HealthCheck struct {
	ServiceID int
	status    string
}

var h HealthCheck

const (
	passStatus = "pass"
	failStatus = "fail"
)

func GenerateCheck() []HealthCheck {

	healthChecks := make([]HealthCheck, 6)

	for i := 0; i < 6; i++ {
		if i%2 == 0 {
			healthChecks[i] = HealthCheck{i, passStatus}
		} else {
			healthChecks[i] = HealthCheck{i, failStatus}
		}
	}
	return healthChecks
}

func main() {

	fmt.Println("Тут будет выведен идентификатор")
	b := GenerateCheck()
	for _, i := range b {
		if i.status == "pass" {
			id := i.ServiceID
			fmt.Println(id)
		}

	}
}
