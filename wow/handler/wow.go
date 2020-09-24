package handler

import (
	"context"

	log "github.com/micro/micro/v3/service/logger"

	wow "wow/proto"
)

type Wow struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *Wow) Call(ctx context.Context, req *wow.Request, rsp *wow.Response) error {
	log.Info("Received Wow.Call request")
	rsp.Msg = "Hello " + req.Name
	return nil
}

// Stream is a server side stream handler called via client.Stream or the generated client code
func (e *Wow) Stream(ctx context.Context, req *wow.StreamingRequest, stream wow.Wow_StreamStream) error {
	log.Infof("Received Wow.Stream request with count: %d", req.Count)

	for i := 0; i < int(req.Count); i++ {
		log.Infof("Responding: %d", i)
		if err := stream.Send(&wow.StreamingResponse{
			Count: int64(i),
		}); err != nil {
			return err
		}
	}

	return nil
}

// PingPong is a bidirectional stream handler called via client.Stream or the generated client code
func (e *Wow) PingPong(ctx context.Context, stream wow.Wow_PingPongStream) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		log.Infof("Got ping %v", req.Stroke)
		if err := stream.Send(&wow.Pong{Stroke: req.Stroke}); err != nil {
			return err
		}
	}
}
