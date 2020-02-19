package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/tamada/morningglory/common"
	"github.com/tamada/morningglory/points"
	"github.com/tamada/morningglory/users"
)

func init() {
	common.InitDatastore()
}

func registerUser(context *gin.Context) {
	var err = users.RegisterUser(context)
	if err != nil {
		context.JSON(400, gin.H{"message": err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "success"})
}

func updateKeyPhrase(context *gin.Context) {
	if err := users.UpdateKeyPhrase(context); err != nil {
		context.JSON(400, gin.H{"message": err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "success"})
}

func deleteUser(context *gin.Context) {
	if err := users.DeleteUser(context); err != nil {
		context.JSON(400, gin.H{"message": err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "success"})
}

func findUser(context *gin.Context) {
	var _, err = users.FindUser(context)
	if err != nil {
		context.JSON(400, gin.H{"message": err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("%s: found", context.Param("userName"))})
}

func registerPoint(context *gin.Context) {
	if err := points.RegisterPoints(context); err != nil {
		context.JSON(400, gin.H{"message": err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "success"})
}

func historyOfPoints(context *gin.Context) {

}

func goMain(args []string) int {
	var router = gin.Default()
	router.GET("/v1/users/:userName", findUser)
	router.POST("/v1/users/:userName", registerUser)
	router.PUT("/v1/users/:userName", updateKeyPhrase)
	router.DELETE("/v1/users/:userName", deleteUser)

	router.POST("/v1/users/:userName/points", registerPoint)
	router.GET("/v1/uesrs/:userName/points", historyOfPoints)

	router.Run()
	return 0
}

func main() {
	var status = goMain(os.Args)
	os.Exit(status)
}
