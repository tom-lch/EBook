package main

import (
	"wow/handler"
	pb "wow/proto"

	"github.com/micro/micro/v3/service"
	"github.com/micro/micro/v3/service/logger"
)

func main() {
	// Create service
	srv := service.New(
		service.Name("wow"),
		service.Version("latest"),
	)

	// Register handler
	pb.RegisterWowHandler(srv.Server(), new(handler.Wow))

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
