package gineasyroute

import (
	"github.com/gin-gonic/gin"
)

type route struct {
	url     string
	handler map[string]gin.HandlerFunc
}
