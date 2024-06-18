package app1

import (
	"log"
	"runtime"

	"github.com/arsu4ka/go-monorepo/pkg/config"
	"github.com/arsu4ka/go-monorepo/pkg/tasks"
	"github.com/hibiken/asynq"
)

type Server struct {
	mux *asynq.ServeMux
	srv *asynq.Server
}

func NewServer(redisConfig config.Redis) *Server {
	srv := asynq.NewServer(
		asynq.RedisClientOpt{Addr: redisConfig.GetAddress()},
		asynq.Config{Concurrency: runtime.NumCPU()},
	)
	return &Server{
		mux: nil,
		srv: srv,
	}
}

func (s *Server) setup() error {
	// create new ServeMux instance
	s.mux = asynq.NewServeMux()

	// bind routes
	s.mux.HandleFunc(tasks.TypeUserJoined, handleUserJoinedTask)

	return nil
}

func (s *Server) Start() error {
	if s.mux == nil {
		if err := s.setup(); err != nil {
			return err
		}
	}

	log.Print("Starting the server")
	return s.srv.Run(s.mux)
}
