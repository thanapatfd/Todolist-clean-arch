package entity

import (
	"errors"
	"fmt"
)

type List struct {
	ID      int
	Name    string
	Status  string
	Details string
}

func (l *List) ChangeStatus(newStatus string) error {
	fmt.Printf("Current Status: %s, New Status: %s\n", l.Status, newStatus)

	if newStatus == "" {
		return errors.New("invalid status")
	}

	if l.Status == "Todo" && newStatus == "Doing" {
		l.Status = newStatus
		return nil
	}

	if l.Status == "Doing" && newStatus == "Done" {
		l.Status = newStatus
		return nil
	}

	if l.Status == "Done" && (newStatus == "Doing" || newStatus == "Todo") {
		return errors.New("cannot change status from Done")
	}

	if l.Status == "Done" && newStatus == "Done" {
		return nil
	}

	return errors.New("invalid status")
}
