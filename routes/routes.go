package routes

import (
	"go_modules/redis"
	"net/http"

	"github.com/gin-gonic/gin"
)

type MyData struct {
	Key string `json:"key"`
}

func Push(c *gin.Context) {

	var data MyData
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	redis.InsertIntoRedis("Key", data.Key)
	c.JSON(http.StatusOK,gin.H{"Message":"Data inserted into redis"})
}

func Home(c *gin.Context) {
	key := "Key"

	data := redis.GetFromRedis(key)

	c.JSON(http.StatusOK, gin.H{
		"Message": data,
	})

}
