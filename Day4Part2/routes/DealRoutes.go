package routes

import (
	"SoftwareGoDay4/softcrypto"
	"SoftwareGoDay4/user"
	"fmt"
	"log"
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
	c.String(http.StatusOK, "\nuser created succesfully")
	//fmt.Println(user.UserDB[0])
}

func handle_signin(c *gin.Context) {
	var tmp user.User
	err := c.BindJSON(&tmp)
	if err != nil {
		return
	}
	for i, each := range user.UserDB {
		if tmp == each {
			tmp.Email = softcrypto.Encrypt(tmp.Email, user.APISECRET)
			tmp.Password = softcrypto.Encrypt(tmp.Password, user.APISECRET)
			c.SetCookie("Cookie", tmp.Email, 10000, "/", "locasthost", false, true)
			user.UserDB[i] = tmp
			c.String(http.StatusOK, "\nFOUND")
			return
		}
	}
	c.String(http.StatusBadRequest, "Couldn't find the user")
}

func handle_cookie(c *gin.Context) {
	str, err := c.Cookie("Cookie")
	if err == http.ErrNoCookie {
		c.String(http.StatusBadRequest, "Can't find the cookie")
		log.Fatal("non")
		return
	}
	for _, each := range user.UserDB {
		if str == each.Email {
			c.String(http.StatusOK, "Found cookie\n")
			c.String(http.StatusOK, softcrypto.Decrypt(each.Email, user.APISECRET))
			c.String(http.StatusOK, "\n")
			c.String(http.StatusOK, softcrypto.Decrypt(each.Password, user.APISECRET))
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
