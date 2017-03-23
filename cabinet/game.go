package main

import (
	"fmt"
	// "time"
)

type (
	RoomStruct struct {
		Players map[*Client]bool
		Send    chan []byte
		Leave   chan *Client
	}
)

func newRoom(players map[*Client]bool) RoomStruct {
	send := make(chan []byte)
	leave := make(chan *Client)
	room := RoomStruct{Players: players, Send: send, Leave: leave}
	return room
}

func (r *RoomStruct) InitRoom() {
	fmt.Println("init")

	for {
		select {
		case message := <-r.Send:
			fmt.Println("work?")
			for client := range r.Players {
				client.send <- message
			}

		case client := <-r.Leave:
			if _, ok := r.Players[client]; ok {
				delete(r.Players, client)
				close(client.send)
			}
		}
	}

	fmt.Println("wery good")
}

// func (r *RoomStruct) StartRoom() {

// 	fmt.Println("1")
// 	time.Sleep(time.Second * 1)
// 	message := "Добро пожаловить в игру"
// 	r.Send <- []byte(message)
// 	fmt.Println("2")
// 	time.Sleep(time.Second * 2)
// 	r.Send <- []byte("Первый вопрос будет таков:")
// 	time.Sleep(time.Second * 2)
// 	r.Send <- []byte("Какого хуя Цырин еще не уволен?")

// }
