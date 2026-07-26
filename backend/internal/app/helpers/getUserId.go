package helpers

import "github.com/gin-gonic/gin"

func GetUserId(ctx *gin.Context) int {
	userID := ctx.MustGet("id").(int)
	return userID
}
