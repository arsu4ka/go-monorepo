package app1

import (
	"context"
	"log"

	"github.com/arsu4ka/go-monorepo/pkg/tasks"
	"github.com/hibiken/asynq"
)

func handleUserJoinedTask(ctx context.Context, t *asynq.Task) error {
	log.Printf(" [x] Processing %s task. User data: %s", tasks.TypeUserJoined, string(t.Payload()))
	return nil
}
