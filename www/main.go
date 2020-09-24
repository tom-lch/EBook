package main

import (
	"www/handler"
	pb "www/proto"

	"github.com/micro/micro/v3/service"
	"github.com/micro/micro/v3/service/logger"
)

func main() {
	// Create service
	srv := service.New(
		service.Name("www"),
		service.Version("latest"),
	)

	// Register handler
	pb.RegisterWwwHandler(srv.Server(), new(handler.Www))

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
