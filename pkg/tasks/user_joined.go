package tasks

import (
	"encoding/json"

	"github.com/hibiken/asynq"
)

const (
	TypeUserJoined = "user:joined"
)

type userJoinedPayload struct {
	UserId string `json:"userId"`
}

func ParseUserJoined(t *asynq.Task) (userJoinedPayload, error) {
	var p userJoinedPayload
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		return p, err
	}

	return p, nil
}

func NewUserJoined(userId string) (*asynq.Task, error) {
	payload, err := json.Marshal(userJoinedPayload{UserId: userId})
	if err != nil {
		return nil, err
	}

	return asynq.NewTask(TypeUserJoined, payload), nil
}
