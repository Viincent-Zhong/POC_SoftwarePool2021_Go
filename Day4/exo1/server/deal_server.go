package server

import "github.com/gin-gonic/gin"

type Server struct {
	App *gin.Engine
}

func NewServer() Server {
	router := Server{gin.Default()}
	return router
}
