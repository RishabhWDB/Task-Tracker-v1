package types

import (
	"fmt"
	"strings"
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
	return fmt.Sprintf("%-4d %-25s %-14s %-18s %-18s",
		t.Id, t.Description, t.Status,
		t.CreatedAt.Format("3:04 pm 2 Jan 06"),
		t.UpdatedAt.Format("3:04 pm 2 Jan 06"))
}

func (ts Tasks) String() string {
	var b strings.Builder
	_, _ = fmt.Fprintf(&b, "%-4s %-25s %-14s %-18s %-18s\n", "ID", "DESCRIPTION", "STATUS", "CREATED", "UPDATED")
	b.WriteString(strings.Repeat("-", 80) + "\n")
	for _, t := range ts {
		b.WriteString(t.String() + "\n")
	}
	return b.String()
}
