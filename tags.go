package main

import "github.com/gin-gonic/gin"
import "net/http"

type tag struct {
	ID          int      `json:"id"`
	Syllabus    syllabus `json:"syllabus"`
	DisplayName string   `json:"display_name"`
}

func BuildTagRoutes(router *gin.Engine) {
	router.GET("/tags/get/all", GetTags)
}

func GetTags(ret *gin.Context) {
	ret.JSON(http.StatusOK, gin.H{
		"message": "It works!",
	})
}
func GetTagById(ret *gin.Context) {
	ret.JSON(http.StatusOK, gin.H{
		"message": "It works!",
	})
}
