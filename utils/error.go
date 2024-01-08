package utils

import (
	"github.com/gin-gonic/gin"
)

func CheckIfErrorContextJson(errorText string, error any, errorStatus int, context *gin.Context) bool {
	if error != nil {
		context.JSON(errorStatus, gin.H{"message": errorText})
		return true
	}

	return false
}
