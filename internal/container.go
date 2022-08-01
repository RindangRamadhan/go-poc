package internal

import (
	hc "go-poc/internal/usecases/health_check"
)

type Container struct {
	HealthCheckUsecase hc.Inport
}

func NewContainer() *Container {
	return &Container{
		HealthCheckUsecase: hc.NewUsecase(),
	}
}
