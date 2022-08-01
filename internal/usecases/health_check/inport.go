package healthcheck

import (
	"context"
	"go-poc/internal/entities"
)

type Inport interface {
	Execute(context.Context) InportResponse
}

type InportResponse struct {
	Data entities.HealthChekResponse
}
