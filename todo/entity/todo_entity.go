package entity

import (
	"errors"
)

type List struct {
	ID      int
	Name    string
	Status  string
	Details string
}

func (l *List) ChangeStatus(newStatus string) error {
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
