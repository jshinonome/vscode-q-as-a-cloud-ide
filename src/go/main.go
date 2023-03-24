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
	hosts := make([]host, 0)
	auth := c.Request.Header["Authorization"]
	user := ""
	password := ""
	if len(auth) > 0 {
		credential, err := base64.StdEncoding.DecodeString(strings.Split(auth[0], " ")[1])
		if err != nil {
			c.String(400, "Invalid Credential")
			return
		}
		credentials := strings.Split(string(credential), ":")
		user = credentials[0]
		password = credentials[1]
	}

	if q.IsConnected() {
		err := q.Sync(&hosts, []string{".discovery.getHosts", user, password})
		if err != nil {
			c.String(400, err.Error())
		}
	}
	c.JSON(200, hosts)
}

type host struct {
	Host  string `json:"host" k:"host"`
	Port  int64  `json:"port" k:"port"`
	Label string `json:"label" k:"label"`
}
