package restAPI

import (
	"fmt"
	"github.com/ccchieh/ginHelper"
	"github.com/ccchieh/gospree/controller/restAPI/handler"
	"github.com/ccchieh/gospree/controller/restAPI/handler/user"
	"github.com/ccchieh/gospree/controller/restAPI/middleware"
	"github.com/ccchieh/gospree/core"
	_ "github.com/ccchieh/gospree/docs"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	"golang.org/x/crypto/bcrypt"
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
	Test()
	r := gin.New()
	r.Use(middleware.LoggerMiddleware(), gin.Recovery())
	//handler中自动建立路由
	ginHelper.Build(new(handler.Helper), r)
	//user的Group中自动建立路由
	ginHelper.Build(new(user.Helper), r.Group("/user"))

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

func Test() {
	passwordOK := "admin"
	passwordERR := "adminxx"

	hash, err := bcrypt.GenerateFromPassword([]byte(passwordOK), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println(err)
	}
	//fmt.Println(hash)

	encodePW := string(hash) // 保存在数据库的密码，虽然每次生成都不同，只需保存一份即可
	fmt.Println(encodePW)

	// 正确密码验证
	err = bcrypt.CompareHashAndPassword([]byte(encodePW), []byte(passwordOK))
	if err != nil {
		fmt.Println("pw wrong")
	} else {
		fmt.Println("pw ok")
	}

	// 错误密码验证
	err = bcrypt.CompareHashAndPassword([]byte(encodePW), []byte(passwordERR))
	if err != nil {
		fmt.Println("pw wrong")
	} else {
		fmt.Println("pw ok")
	}
}
