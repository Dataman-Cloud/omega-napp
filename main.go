package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/Dataman-Cloud/omega-napp/config"
	"github.com/Dataman-Cloud/omega-napp/logger"
	"github.com/Dataman-Cloud/omega-napp/router"
	"github.com/Dataman-Cloud/omega-napp/router/middleware"
	"github.com/Dataman-Cloud/omega-napp/store"
	log "github.com/Sirupsen/logrus"
	_ "github.com/gin-gonic/gin"
)

func main() {
	logger.Init()
	log.Info("omega-napp starting ......")
	config.InitConfig("deploy/env")

	store.InitDB()
	store.UpgradeDB()
	store.InitCache()

	h := router.Load(
		middleware.Authenticate,
	)
	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", config.Pairs().Port),
		Handler:        h,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}
