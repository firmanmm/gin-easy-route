package gineasyroute

import (
	"github.com/gin-gonic/gin"
)

type middleware struct {
	url     string
	handler gin.HandlerFunc
}
