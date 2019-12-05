package restAPI

import (
	"fmt"
	"github.com/ccchieh/gospree/controller/restAPI/handler"
	"github.com/ccchieh/gospree/controller/restAPI/handler/user"
	"github.com/ccchieh/gospree/controller/restAPI/middleware"
	"github.com/ccchieh/gospree/core"
	_ "github.com/ccchieh/gospree/docs"
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
	//handler中自动建立路由
	handler.Build(r)
	//user的Group中自动建立路由
	user.Build(r.Group("/user"))

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
