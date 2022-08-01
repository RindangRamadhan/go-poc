package healthcheck

import (
	"go-poc/common/response"

	"github.com/gin-gonic/gin"

	hc "go-poc/internal/usecases/health_check"
)

func HealthCheckHandler(inport hc.Inport) gin.HandlerFunc {
	return func(c *gin.Context) {
		resp := inport.Execute(c.Copy().Request.Context())
		response.WriteSuccess(c, "Success health check", resp.Data, nil)
	}
}
