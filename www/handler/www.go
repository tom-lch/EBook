package handler

import (
	"context"

	log "github.com/micro/micro/v3/service/logger"

	www "www/proto"
)

type Www struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *Www) Call(ctx context.Context, req *www.Request, rsp *www.Response) error {
	log.Info("Received Www.Call request")
	rsp.Msg = "Hello " + req.Name
	return nil
}

// Stream is a server side stream handler called via client.Stream or the generated client code
func (e *Www) Stream(ctx context.Context, req *www.StreamingRequest, stream www.Www_StreamStream) error {
	log.Infof("Received Www.Stream request with count: %d", req.Count)

	for i := 0; i < int(req.Count); i++ {
		log.Infof("Responding: %d", i)
		if err := stream.Send(&www.StreamingResponse{
			Count: int64(i),
		}); err != nil {
			return err
		}
	}

	return nil
}

// PingPong is a bidirectional stream handler called via client.Stream or the generated client code
func (e *Www) PingPong(ctx context.Context, stream www.Www_PingPongStream) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		log.Infof("Got ping %v", req.Stroke)
		if err := stream.Send(&www.Pong{Stroke: req.Stroke}); err != nil {
			return err
		}
	}
}
