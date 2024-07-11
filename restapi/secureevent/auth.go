package secureevent

import (
	"fmt"
	"net/http"

	"example.com/first-app/secure"
	"github.com/gin-gonic/gin"
)

func Authenticate(context *gin.Context) {
	token := context.Request.Header.Get("Authorization")
	if token == "" {
		fmt.Print(token)
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "not 1authorized"})
		return
	}
	authid, err := secure.Verifytoken(token)
	if err != nil {
		fmt.Print(err)
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "not authorized"})
		return
	}
	context.Set("authid", authid)
	context.Next()
}
