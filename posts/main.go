package main

import (
	"github.com/micro/micro/v3/service"
	"github.com/micro/micro/v3/service/logger"

	"github.com/emadghaffari/micro-blog/posts/handler"
)

func main() {
	// Create service
	srv := service.New(
		service.Name("posts"),
		service.Version("latest"),
	)

	srv.Handle(handler.NewPosts())
	// Register handler

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
