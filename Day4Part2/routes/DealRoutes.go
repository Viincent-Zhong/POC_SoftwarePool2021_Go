package routes

import (
	"SoftwareGoDay4/softcrypto"
	"SoftwareGoDay4/user"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func handle_body(c *gin.Context) {
	var tmp []user.User
	err := c.BindJSON(&tmp)
	if err != nil {
		return
	}
	user.UserDB = append(user.UserDB, tmp...)
	fmt.Printf("%#v\n", tmp)
	c.JSON(http.StatusOK, tmp)
	//fmt.Println(user.UserDB[0])
}

func handle_signin(c *gin.Context) {
	var tmp user.User
	err := c.BindJSON(&tmp)
	if err != nil {
		return
	}
	for _, each := range user.UserDB {
		if tmp == each {
			tmp.Email = softcrypto.Encrypt(tmp.Email, user.APISECRET)
			tmp.Password = softcrypto.Encrypt(tmp.Password, user.APISECRET)
			c.Writer.Header().Add("Cookie", tmp.Email)
			each = tmp
			break
		}
	}
}

func handle_cookie(c *gin.Context) {
	str, err := c.Cookie("Cookie")
	if err == http.ErrNoCookie {
		return
	}
	for _, each := range user.UserDB {
		if str == each.Email {
			fmt.Printf(softcrypto.Decrypt(each.Email, user.APISECRET))
			fmt.Printf(softcrypto.Decrypt(each.Password, user.APISECRET))
		}
	}
}

func ApplyRoute(r *gin.Engine) {
	r.POST("/signup-session", handle_body)
	r.GET("/test", func(c *gin.Context) {
		fmt.Println(user.UserDB)
	})
	r.POST("/signin-session", handle_signin)
	r.GET("/me-session", handle_cookie)
}
