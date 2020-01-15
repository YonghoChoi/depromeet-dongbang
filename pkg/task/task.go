package task

import (
	"fmt"
	"strconv"
	"time"
)

type Deadline struct {
	time.Time
}

func (d Deadline) MarshalJSON() ([]byte, error) {
	return strconv.AppendInt(nil, d.Unix(), 10), nil
}

func (d *Deadline) UnmarshalJSON(data []byte) error {
	unix, err := strconv.ParseInt(string(data), 10, 64)
	if err != nil {
		return err
	}
	d.Time = time.Unix(unix, 0)
	return nil
}

type status int

const (
	UNKNOWN status = iota
	TODO
	DONE
)

type Task struct {
	Title    string
	Status   status
	Deadline *Deadline
	Priority int
}

func (t Task) String() string {
	check := "v"
	if t.Status != DONE {
		check = " "
	}

	return fmt.Sprintf("[%s] %s %s", check, t.Title, t.Deadline)
}
