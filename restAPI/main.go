package restAPI

import (
	"fmt"
	"github.com/ccchieh/gospree/core"
	_ "github.com/ccchieh/gospree/docs"
	"github.com/ccchieh/gospree/restAPI/handler"
	"github.com/ccchieh/gospree/restAPI/middleware"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	"net/http"
	"time"
)

type RestServer struct {
	port int
}

func (rest *RestServer) Start(port int) {
	if port < 0 || port > 65535 {
		core.Log.Error("Port range error")
		return
	}

	//gin.SetMode(gin.ReleaseMode)

	r := gin.New()
	r.Use(middleware.LoggerMiddleware(), gin.Recovery())

	r.GET("/HelloHandler", handler.HelloHandler)

	userG := r.Group("/user")
	{
		userG.POST("", handler.CreateUserHandler)
		userG.GET("/info", handler.GetUserInfoHandler)
	}

	r.GET("/swagger/*any", ginSwagger.DisablingWrapHandler(swaggerFiles.Handler, "RELEASE"))

	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", port),
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	if err := s.ListenAndServe(); err != nil {
		core.Log.Err(err)
	}
}
