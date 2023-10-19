package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"todo"
)

func (h *Handler) todoCreate(c *gin.Context) {
	var input todo.Todo

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	id, err := h.services.Todo.CreateTodo(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"ok": id,
	})
}

func (h *Handler) allContent(c *gin.Context) {
	content, _ := h.services.AllContent()
	c.JSON(http.StatusOK, map[string]interface{}{
		"content": content,
	})
}

func (h *Handler) deleteTodo(c *gin.Context) {
	var input deleteTodoInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	str, _ := h.services.DeleteTodo(input.Title)
	c.JSON(http.StatusOK, map[string]interface{}{
		"status": str,
	})
}
func (h *Handler) updateTodo(c *gin.Context) {
	var input todo.Todo
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	str, err := h.services.UpdateTodo(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"status": str,
	})
}
