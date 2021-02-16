package routes

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func world(c *gin.Context) {
	/*c.JSON(http.StatusOK, gin.H{
	"test": "aled"})*/
	c.String(http.StatusOK, fmt.Sprintf("%v tfk %v pk", "test", "aled"))
	c.String(http.StatusOK, "sqoijqsifjpsfjsmdowlighmldighmldqifghmlidqfhgpqdifghpodifgh")
	/*c.Query()
	c.Param()
	c.Header()
	c.Cookie()*/
	//c.Request.Body()
}

func handle_query(c *gin.Context) {
	id, erri := c.GetQuery("id")
	name, errn := c.GetQuery("name")
	if erri == false || errn == false {
		return
	}
	c.String(http.StatusOK, fmt.Sprintf("id %v, name %v", id, name))
}

func handle_param(c *gin.Context) {
	id := c.Param("id")
	c.String(http.StatusOK, id)
}

func handle_header(c *gin.Context) {
	str := c.GetHeader("str")
	c.String(http.StatusOK, str)
}

func handle_cookie(c *gin.Context) {
	str, err := c.Cookie("str")
	if err == http.ErrNoCookie {
		return
	}
	c.String(http.StatusOK, str)
}

func handle_body(c *gin.Context) {
	//str := c.Request.Body()
}

func ApplyRoutes(r *gin.Engine) {
	r.GET("/hello", world)
	r.GET("/repeat-my-query", handle_query)
	r.GET("/repeat-my-param/:message", handle_param)
	r.GET("/repeat-my-header", handle_header)
	r.GET("/repeat-my-cookie", handle_cookie)
	//r.GET("/hello", )
}
