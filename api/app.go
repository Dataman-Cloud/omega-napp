package api

import (
	log "github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
)

func DeployApp(c *gin.Context) {
	log.Info("-------: ", c.MustGet("uid"))
	c.JSON(200, gin.H{
		"message": "Deplay app",
	})
}
