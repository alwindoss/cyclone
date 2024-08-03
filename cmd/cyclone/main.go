package main

import (
	"net/http"

	"github.com/8ideaz/cyclone/internal/testcase"

	"github.com/gin-gonic/gin"
)

var store = testcase.NewTestCaseStore()

func main() {
	router := gin.Default()

	router.POST("/testcases", func(c *gin.Context) {
		var tc testcase.TestCase
		if err := c.ShouldBindJSON(&tc); err == nil {
			store.AddTestCase(&tc)
			c.JSON(http.StatusOK, gin.H{"status": "added"})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
	})

	router.GET("/testcases", func(c *gin.Context) {
		c.JSON(http.StatusOK, store.ListTestCases())
	})

	router.Run(":8080")
}
