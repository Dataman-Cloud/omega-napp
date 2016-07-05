package router

import (
	"net/http"
	"time"

	"github.com/Dataman-Cloud/omega-napp/api"
	"github.com/Dataman-Cloud/omega-napp/logger"
	log "github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
)

func Load(middleware ...gin.HandlerFunc) http.Handler {
	e := gin.New()
	e.Use(gin.Recovery(), logger.Ginrus(log.StandardLogger(), time.RFC3339, true))
	gin.SetMode(gin.ReleaseMode)
	//e.Use(middleware...)

	groupV4 := e.Group("/api/v4", middleware...)
	{
		groupV4.POST("/clusters/:cid/apps", api.DeployApp)
	}

	return e
}
