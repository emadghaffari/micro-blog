package main

import (
	"github.com/micro/micro/v3/service"
	"github.com/micro/micro/v3/service/logger"

	"github.com/emadghaffari/micro-blog/tags/handler"
	pb "github.com/emadghaffari/micro-blog/tags/proto"
)

func main() {
	// Create service
	srv := service.New(
		service.Name("tags"),
		service.Version("latest"),
	)

	// Register handler
	pb.RegisterTagsHandler(srv.Server(), new(handler.Tags))

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
