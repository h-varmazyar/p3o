package utils

import "github.com/gin-gonic/gin"

func FetchUserId(ctx *gin.Context) uint{
	return ctx.GetUint("user_id")
}