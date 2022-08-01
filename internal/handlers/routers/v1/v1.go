package v1

import (
	"go-poc/common/middleware"
	"go-poc/internal"

	"github.com/gin-gonic/gin"

	hc "go-poc/internal/handlers/health_check"
)

type (
	RoutePath struct {
		Method   string
		Group    string
		Pattern  string
		Handlers []gin.HandlerFunc
	}
)

func Paths() []RoutePath {
	return []RoutePath{
		{
			Method: "GET", Group: "", Pattern: "/", Handlers: []gin.HandlerFunc{
				middleware.Authentication(), hc.HealthCheckHandler(internal.NewContainer().HealthCheckUsecase),
			},
		},
		{
			Method: "GET", Group: "", Pattern: "/health-check", Handlers: []gin.HandlerFunc{
				middleware.Authentication(), hc.HealthCheckHandler(internal.NewContainer().HealthCheckUsecase),
			},
		},
		{
			Method: "GET", Group: "users", Pattern: "/list", Handlers: []gin.HandlerFunc{
				middleware.Authentication(), hc.HealthCheckHandler(internal.NewContainer().HealthCheckUsecase),
			},
		},
		{
			Method: "GET", Group: "users", Pattern: "/detail/:id", Handlers: []gin.HandlerFunc{
				middleware.Authentication(), hc.HealthCheckHandler(internal.NewContainer().HealthCheckUsecase),
			},
		},
	}
}
