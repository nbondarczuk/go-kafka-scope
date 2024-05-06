package admin

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ReadCluster(c *gin.Context) {
	var rval interface{}
	r := map[string]interface{}{
		"Status": "Ok",
		"Object": rval,
	}
	c.JSON(http.StatusOK, r)
	return
}

func ReadConfig(c *gin.Context) {
	var rval interface{}
	r := map[string]interface{}{
		"Status": "Ok",
		"Object": rval,
	}
	c.JSON(http.StatusOK, r)
	return
}

func ReadTopic(c *gin.Context) {
	var rval interface{}
	r := map[string]interface{}{
		"Status": "Ok",
		"Object": rval,
	}
	c.JSON(http.StatusOK, r)
	return
}

func ReadConsumerGroup(c *gin.Context) {
	var rval interface{}
	r := map[string]interface{}{
		"Status": "Ok",
		"Object": rval,
	}
	c.JSON(http.StatusOK, r)
	return
}
