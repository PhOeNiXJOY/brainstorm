package main

import (
	// "fmt"
	// "github.com/gin-gonic/gin"
	// "github.com/kennygrant/sanitize"
	// "github.com/mirrr/mgo-wrapper"
	// "golang.org/x/exp/utf8string"
	"gopkg.in/gin-gonic/gin.v1"
	// "gopkg.in/mgo.v2"
	// "gopkg.in/mgo.v2/bson"
	// "gopkg.in/mirrr/types.v1"
	// "strings"
	// "time"
)

func Page(c *gin.Context) {

	settings := gin.H{
		"title": "Тестовый вебсокет",
	}

	// key, ex := c.Get("user")
	// if ex {
	// 	settings = gin.H{
	// 		"title":      "Новости",
	// 		"active":     obj{"news": true},
	// 		"allSliders": getAllNews(),
	// 		"user":       key.(Us),
	// 		"rubrics":    rubrics.GetArr(),
	// 	}
	// }
	c.HTML(200, "index.html", settings)
}
