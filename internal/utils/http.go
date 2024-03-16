package utils

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetBody binds the JSON body of the request to a struct and returns it
func GetBody[ValidatorStruct any](ctx *gin.Context) ValidatorStruct {
	var body ValidatorStruct
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	return body
}

// GetParam binds the URI parameters of the request to a struct and returns it
func GetParam[ValidatorStruct any](ctx *gin.Context) ValidatorStruct {
	var param ValidatorStruct
	if err := ctx.ShouldBindUri(&param); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	return param
}

// GetQueryInt gets an integer query parameter from the request URL
func GetQueryInt(ctx *gin.Context, key string, defaultValue int) int {
	valueStr := ctx.DefaultQuery(key, strconv.Itoa(defaultValue))
	value, err := strconv.Atoi(valueStr)
	if err != nil {
		value = defaultValue
	}
	return value
}

// GetQueryString gets a string query parameter from the request URL
func GetQueryString(ctx *gin.Context, key string, defaultValue string) string {
	return ctx.DefaultQuery(key, defaultValue)
}
