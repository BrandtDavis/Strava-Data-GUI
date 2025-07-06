package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type activity struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
	Date string `json:"date"`
}

var activities = []activity{
	{ID: "1", Name: "Easy Run", Type: "Run", Date: "2025-06-12"},
}

func getTodos(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, activities)
}

func main() {
	router := gin.Default()
	router.GET("/activities", getTodos)
	router.Run("localhost:8080")
}
