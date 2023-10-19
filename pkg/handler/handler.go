package handler

import (
	"github.com/gin-gonic/gin"
	"todo/pkg/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}
func (h *Handler) InitRoutes() *gin.Engine {
	routers := gin.New()

	auth := routers.Group("/auth")
	{
		auth.POST("/sign-up", h.register)
		auth.POST("/sign-in", h.authorization)
		auth.POST("/start", h.start)
	}

	todo := routers.Group("/todo")
	{
		todo.POST("/todo-create", h.todoCreate)
		todo.GET("/todo-allContent", h.allContent)
		todo.DELETE("/todo-delete", h.deleteTodo)
		todo.POST("/todo-update", h.updateTodo)
	}

	return routers
}
