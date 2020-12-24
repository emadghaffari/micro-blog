package main

import (
	"github.com/micro/micro/v3/service"
	"github.com/micro/micro/v3/service/logger"

	"github.com/emadghaffari/micro-blog/posts/handler"
	pb "github.com/emadghaffari/micro-blog/posts/proto"
)

func main() {
	// Create service
	srv := service.New(
		service.Name("posts"),
		service.Version("latest"),
	)

	// Register handler
	pb.RegisterPostsHandler(srv.Server(), new(handler.Posts))

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
