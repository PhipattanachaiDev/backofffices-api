package v2

import (
	response "ezview.asia/ezview-web/ezview-lite-back-office/types/responses"
	"github.com/gin-gonic/gin"
)

var db = make(map[string]string)

func Ping(c *gin.Context) {
	response.OK(c, "pong")
}
