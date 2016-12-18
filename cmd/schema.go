package cmd

import (
	"fmt"
	"strconv"
	"time"
)

type Unixtime struct {
	time.Time
}

func IntToUnixtime(timestamp int) Unixtime {
	return Unixtime{time.Unix(int64(timestamp), 0)}
}

func (t *Unixtime) MarshalJSON() ([]byte, error) {
	timestamp := fmt.Sprint(t.Unix())
	return []byte(timestamp), nil
}

func (t *Unixtime) UnmarshalJSON(b []byte) error {
	timestamp, err := strconv.Atoi(string(b))
	if err != nil {
		return err
	}
	t.Time = time.Unix(int64(timestamp), 0)

	return nil
}

func (t Unixtime) String() string {
	return t.Format(time.RFC3339)
}

type AppStack struct {
	ID         int      `json:"id"`
	Name       string   `json:"name"`
	InsertedAt Unixtime `json:"inserted_at"`
	UpdatedAt  Unixtime `json:"updated_at"`
}

type AppStackShowRequest struct {
	ID int `json:"id"`
}

type AppStackShowResponse struct {
	AppStack AppStack `json:"app_stack"`
}
