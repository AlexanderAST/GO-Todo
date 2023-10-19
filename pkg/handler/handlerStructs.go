package handler

type signInInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type deleteTodoInput struct {
	Title string `json:"title" binding:"required"`
}
