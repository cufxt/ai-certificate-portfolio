package handler

import (
	"context"

	"connectrpc.com/connect"
	healthv1 "github.com/cufxt/ai-certificate-portfolio/backend/gen/health/v1"
)

type HealthHandler struct{}

func NewHealthHandler() *HealthHandler {
	return &HealthHandler{}
}

func (h *HealthHandler) Ping(
	ctx context.Context,
	req *connect.Request[healthv1.PingRequest],
) (*connect.Response[healthv1.PingResponse], error) {
	return connect.NewResponse(&healthv1.PingResponse{Status: "ok"}), nil
}
