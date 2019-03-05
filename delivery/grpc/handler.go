package grpc

import (
	"context"

	"google.golang.org/grpc"

	"github.com/rizkix/wired/controller"
	pb "github.com/rizkix/wired/proto"
)

type Handler struct {
	Controller controller.Controller
	Instance   *grpc.Server
}

func (h *Handler) GetResourceId(ctx context.Context, in *pb.Request) (*pb.Response, error) {
	a := h.Controller.Get("a")
	return &pb.Response{ResourceId: a.ResourceID}, nil
}

func New(c controller.Controller) Handler {
	s := grpc.NewServer()
	handler := Handler{Controller: c, Instance: s}

	pb.RegisterDataServer(s, &handler)

	return handler
}
