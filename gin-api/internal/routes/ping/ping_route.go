package ping

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddPingRoute(api *gin.Engine) error {
	routerGroup := api.Group("/ping")

	routerGroup.GET("", handlePing)

	return nil
}

func handlePing(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}
