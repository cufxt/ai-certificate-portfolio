package handler_test

import (
	"context"
	"testing"

	"connectrpc.com/connect"
	"github.com/stretchr/testify/assert"

	healthv1 "github.com/cufxt/ai-certificate-portfolio/backend/gen/health/v1"
	"github.com/cufxt/ai-certificate-portfolio/backend/internal/handler"
)

func TestHealthHandler_Ping(t *testing.T) {
	h := handler.NewHealthHandler()

	resp, err := h.Ping(context.Background(), connect.NewRequest(&healthv1.PingRequest{}))

	assert.NoError(t, err)
	assert.Equal(t, "ok", resp.Msg.Status)
}
