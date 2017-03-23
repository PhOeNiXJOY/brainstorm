package main

import (
	"flag"
	"gopkg.in/gin-gonic/gin.v1"
	// "log"
	// "net/http"
	"fmt"
	"github.com/night-codes/go-gin-ttpl"
	// "gopkg.in/night-codes/types.v1"
	"os"
	"text/template"
	"time"
)

var (
	hub = newHub()
)

func hubMain(c *gin.Context) {
	serveWs(hub, c.Writer, c.Request)
}

func main() {
	flag.Parse()
	go hub.run()
	// go test()
	r := gin.Default()

	ttpl.Use(r, "templates/main/*", template.FuncMap{"dot": dot})
	r.Static("files", "./files")
	// r.Use(Middle)
	r.GET("/", Page)
	r.GET("/ws", hubMain)
	r.GET("/startgame", startGame)

	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = ":47000"
	}
	panic(r.Run(port))
}

func startGame(c *gin.Context) {

	room := newRoom(hub.clients)
	fmt.Println("ready to start")
	go room.InitRoom()

	time.Sleep(time.Second * 1)
	fmt.Println("1")
	message := "Добро пожаловить в игру"
	room.Send <- []byte(message)
	fmt.Println("2")
	time.Sleep(time.Second * 2)
	room.Send <- []byte("Первый вопрос будет таков:")
	time.Sleep(time.Second * 2)
	room.Send <- []byte("Какого хуя Цырин еще не уволен?")
	// room.StartRoom()

}
