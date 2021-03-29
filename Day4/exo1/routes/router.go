package routes

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func world(c *gin.Context) {
	c.String(http.StatusOK, fmt.Sprintf("%v tfk %v pk", "test", "aled"))
	c.String(http.StatusOK, "sqoijqsifjpsfjsmdowlighmldighmldqifghmlidqfhgpqdifghpodifgh")
}

func handle_query(c *gin.Context) {
	id, erri := c.GetQuery("id")
	name, errn := c.GetQuery("name")
	//str := c.Request.URL.Query()

	if erri == false || errn == false {
		c.String(http.StatusBadRequest, "Pas bon")
		return
	}
	/*test, err := json.Marshal(str)
	if err != nil {
		c.String(http.StatusBadRequest, "pas bon")
		return
	}
	c.JSON(http.StatusOK, test)*/
	var s_query struct {
		Id   string
		Name string
	}
	s_query.Id = id
	s_query.Name = name
	c.JSON(http.StatusOK, s_query)
	//c.String(http.StatusOK, fmt.Sprintf("id: %v, name: %v", id, name))
}

func handle_param(c *gin.Context) {
	name := c.Params.ByName("message")
	if name == "" {
		c.String(http.StatusBadRequest, "Pas bon")
		return
	}
	c.String(http.StatusOK, name)
}

func handle_header(c *gin.Context) {
	str := c.GetHeader("param")
	c.String(http.StatusOK, str)
}

func handle_cookie(c *gin.Context) {
	name, err := c.Cookie("name")
	if err == http.ErrNoCookie {
		return
	}
	c.String(http.StatusOK, name)
}

func handle_body(c *gin.Context) {
	json := make([]byte, c.Request.ContentLength)
	_, err := c.Request.Body.Read(json)
	if err != nil {
		return
	}
	c.String(http.StatusOK, string(json))
}

func ApplyRoutes(r *gin.Engine) {
	str := os.Getenv("HELLO_MESSAGE")
	if str == "" {
		log.Fatalf("no message defined")
	}
	r.GET(str, world)
	r.GET("/repeat-my-query", handle_query)
	r.GET("/repeat-my-param/:message", handle_param)
	r.GET("/repeat-my-header", handle_header)
	r.GET("/repeat-my-cookie", handle_cookie)
	r.POST("/repeat-my-body", handle_body)
}
