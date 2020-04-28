package controllers

import (
	"fmt"
	"net/http"
	"time"

	"common_dashboard_backend/apps/Dashboard/interactors"
	"common_dashboard_backend/apps/Dashboard/interfaces"

	. "common_dashboard_backend/common/logger"

	// "github.com/xeipuuv/gojsonschema"

	"github.com/gin-gonic/contrib/gzip"
	"github.com/gin-gonic/gin"
)

var DashboardUseCase interfaces.DashboardUseCases
var secret = "developmentSecretIsNotSoSecret"

// var MediaPath string

const schema = `
{
	"type": "object",
	"properties": {
	  "x":      { "type": "number" },
	  "y":     { "type": "number" },
	  "width":   { "type": "number" },
	  "height": { "type": "number" }
	},
	"required": ["x", "y","width","height"]
  }
`

func StartApplicationBackend() {
	// Creates a gin router with default middleware:
	// logger and recovery (crash-free) middleware
	//router := gin.LoggerWithWriter()
	router := gin.New()

	router.Use(gin.Recovery(), Logger(), Headers())
	router.Use(gzip.Gzip(gzip.DefaultCompression))

	//gin.SetMode(gin.DebugMode)

	public, private := InitRoutesGroups(router)
	InitRoutes(public, private)

	DashboardUseCase = &interactors.DashboardInteractor{}

	// MediaPath = mediaPath
	//Compare schema and data
	// loader := gojsonschema.NewStringLoader(`{"type": "string"}`)

	// Listen and server on 0.0.0.0:8200
	err := http.ListenAndServe(":8300", router)
	if err != nil {
		// LogError(err)
		fmt.Println(err)
	}
	LogInfo("Server is running on :8300 port")
}

func InitRoutesGroups(router *gin.Engine) (public, private *gin.RouterGroup) {
	public = router.Group("/")
	private = router.Group("/api/")
	//test = router.Group("/api-test/")

	//private.Use(accessTypeAvailabilityCheck())
	//private.Use(RoleCheck())
	return
}

func Headers() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Run this on all requests
		// Should be moved to a proper middleware
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type,Token,Authorization,X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "DELETE,GET,HEAD,POST,PUT,OPTIONS,TRACE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
		}

		c.Next()
	}
}

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()

		// Set example variable
		// c.Set("example", "12345")
		// before request

		c.Next()

		// after request
		latency := time.Since(t)
		path := c.Request.URL.Path
		method := c.Request.Method

		// access the status we are sending
		status := c.Writer.Status()
		LogInfo(status, method, path, latency)
	}
}
