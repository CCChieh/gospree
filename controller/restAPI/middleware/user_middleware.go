package middleware

import (
	"github.com/ccchieh/gospree/core"
	"github.com/ccchieh/gospree/model"
	"github.com/ccchieh/gospree/util/ret"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func TokenMiddleware() gin.HandlerFunc {
	//下面代码暂时不用，是用来对当前身份的认证
	//var authFunc []func(tokenID uint, userID uint) bool
	//selfAuth := func(tokenID uint, userID uint) bool { return tokenID == userID }
	//for _, value := range identities {
	//	switch value {
	//	case "self":
	//		authFunc=append(authFunc, selfAuth)
	//	}
	//}

	return func(c *gin.Context) {
		auth := c.GetHeader("authorization")

		if auth == "" {
			ret.Result(c, http.StatusUnauthorized, nil, core.ErrValidation)
			c.Abort()
			return
		}
		tokens := strings.Split(auth, " ")
		if len(tokens) < 2 {
			ret.Result(c, http.StatusUnauthorized, nil, core.ErrValidation)
			c.Abort()
			return
		}
		token := tokens[1]
		user := new(model.User)
		if !user.CheckToken(token) {
			ret.Result(c, http.StatusUnauthorized, nil, core.ErrValidation)
			c.Abort()
			return
		}
		c.Set("user", user)
		c.Next()
	}
}
