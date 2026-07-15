package types

import (
	"fmt"
	"time"
)

type Task struct {
	Id          int
	Description string
	Status      string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type Tasks []Task

func (t Task) String() string {
	return fmt.Sprintf("Id: %d Description: %s Status: %s CreatedAt: %s UpdatedAt: %s",
		t.Id, t.Description, t.Status,
		t.CreatedAt.Format("2001-02-04 15:04:05"),
		t.UpdatedAt.Format("2001-02-04 15:04:05"))
}
