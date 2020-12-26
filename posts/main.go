package main

import (
	"github.com/micro/micro/service"
	"github.com/micro/micro/v3/service/logger"

	"github.com/emadghaffari/micro-blog/posts/handler"
	tags "github.com/emadghaffari/micro-blog/tags/proto"
)

func main() {
	// Create service
	srv := service.New(
		service.Name("posts"),
		service.Version("latest"),
	)

	srv.Handle(handler.NewPosts(
		tags.NewTagsService("tags", srv.Client()),
	))
	// Register handler

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
