package api

import (
	"github.com/gin-gonic/gin"
)

func CreateAPI() (*gin.Engine, error) {
	api := gin.New()

	return api, nil
}
