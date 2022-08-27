package main

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

var inMemoryStore = make(map[string]string)
var redirectURL = "http://0.0.0.0:9000"

func setupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/vote", func(c *gin.Context) {
		payload := gin.H{}
		rogueA := "You're Hacked! Your GitOps Supply Chain is compromised :D:D:D"
		payload["Dear User"] = rogueA
		c.JSON(http.StatusOK, payload)
	})

	r.POST("/vote", func(c *gin.Context) {
		buf := make([]byte, 1024)
		num, _ := c.Request.Body.Read(buf)
		reqBody := buf[0:num]
		temp := map[string]string{}
		json.Unmarshal(reqBody, &temp)
		c.JSON(http.StatusOK, reqBody)
		voter_id := temp["voter_id"]
		vote := temp["vote"]
		inMemoryStore[voter_id] = vote
	})
	return r
}

func main() {
	r := setupRouter()
	r.Run(":9000")
}
