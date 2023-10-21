package handler

import "github.com/gin-gonic/gin"

type Error struct {
	Msg string `json:"msg"`
}

func ErrorHandleResponse(ctx *gin.Context, statusCode int, msg string) {
	ctx.AbortWithStatusJSON(statusCode, Error{Msg: msg})
}
