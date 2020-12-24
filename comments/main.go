package main

import (
	"github.com/micro/micro/v3/service"
	"github.com/micro/micro/v3/service/logger"

	"github.com/emadghaffari/micro-blog/comments/handler"
)

func main() {
	// Create service
	srv := service.New(
		service.Name("comments"),
		service.Version("latest"),
	)

	// Register handler
	srv.Handle(handler.NewComments())

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
