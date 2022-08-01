package healthcheck

import (
	"context"
	"go-poc/internal/entities"
)

type interactor struct {
}

// NewUsecase --
func NewUsecase() Inport {
	return interactor{}
}

func (i interactor) Execute(ctx context.Context) InportResponse {
	return InportResponse{
		Data: entities.HealthChekResponse{
			Name:    "Service Poc",
			Version: "0.0.1",
		},
	}
}
