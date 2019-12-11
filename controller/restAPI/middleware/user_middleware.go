package middleware

import (
	"github.com/ccchieh/gospree/model"
	"github.com/ccchieh/gospree/util/ret"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func TokenMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		//token := util.GetToken(r)
		//if token == "" {
		//	jsonapi.Result(w, http.StatusUnauthorized, nil, elog.NewErrUnauthorized(r.URL.Path))
		//	return
		//}
		//user := new(User)
		//if !user.CheckToken(token) {
		//	jsonapi.Result(w, http.StatusUnauthorized, nil, elog.NewErrUnauthorized(r.URL.Path))
		//	return
		//}
		//handler.ServeHTTP(w, r)
		auth := c.GetHeader("authorization")

		if auth == "" {
			ret.Result(c, http.StatusUnauthorized, nil, ret.ErrValidation)
			c.Abort()
			return
		}
		tokens := strings.Split(auth, " ")
		if len(tokens) < 2 {
			ret.Result(c, http.StatusUnauthorized, nil, ret.ErrValidation)
			c.Abort()
			return
		}
		token := tokens[1]
		user := new(model.User)
		if !user.CheckToken(token) {
			ret.Result(c, http.StatusUnauthorized, nil, ret.ErrValidation)
			c.Abort()
			return
		}
		c.Set("user", user)
		c.Next()
	}
}
