package main

import (
	"encoding/base64"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/jshinonome/geek"
)

var q = geek.QProcess{Port: 1800}

func main() {
	q.Dial()
	r := gin.Default()
	r.GET("/hosts", getHosts)
	r.Run(":8080")
}

func getHosts(c *gin.Context) {
	h := make([]host, 0)
	credential, err := base64.StdEncoding.DecodeString(strings.Split(c.Request.Header["Authorization"][0], " ")[1])
	if err != nil {
		c.JSON(400, "Invalid Credential")
		return
	}

	credentials := strings.Split(string(credential), ":")
	user := credentials[0]
	password := credentials[1]
	if q.IsConnected() {
		q.Sync(&h, []string{".discovery.getHosts", user, password})
	}
	c.JSON(200, h)
}

type host struct {
	Host  string `json:"host" k:"host"`
	Port  int64  `json:"port" k:"port"`
	Label string `json:"label" k:"label"`
}
