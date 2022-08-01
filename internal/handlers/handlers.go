package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	v1 "go-poc/internal/handlers/routers/v1"
)

type (
	Router struct {
		router *gin.Engine
	}

	RoutePath struct {
		Method   string
		Group    string
		Pattern  string
		Handlers []gin.HandlerFunc
	}
)

func NewRouter(router *gin.Engine) *Router {
	return &Router{
		router: router,
	}
}

func (_h *Router) RegisterRouter() {
	// API V1
	apiv1 := _h.router.Group("/api/v1")
	for _, v := range v1.Paths() {
		apiGroup := apiv1.Group(v.Group)

		switch v.Method {
		case http.MethodGet:
			apiGroup.GET(v.Pattern, v.Handlers...)
		}
	}

}
