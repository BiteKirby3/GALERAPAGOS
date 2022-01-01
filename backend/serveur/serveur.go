package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"golang.org/x/net/websocket"
)

type Serveur struct {
	nbPlayers int
	nbTurns   int
}

func NewServeur() *Serveur {
	return &Serveur{nbPlayers: 3, nbTurns: 2}
}

func (s *Serveur) React(ws *websocket.Conn) {
	var err error

	for {
		var reply string
		//Receive message from front or Send message to front
		err = websocket.Message.Receive(ws, &reply)
		if err != nil {
			fmt.Println("Can't receive")
		} else {
			fmt.Println("Received back from client: " + reply)

			//convert JSON string response to map
			responseMap := make(map[string]string)
			err := json.Unmarshal([]byte(reply), &responseMap)
			if err != nil {
				panic(err)
			}

			if responseMap["fromPage"] == "home" { //If the request is sended from the homePage :  we should store the nbPlayers and nbTurns value.
				s.nbPlayers, _ = strconv.Atoi(responseMap["nbPlayers"])
				s.nbTurns, _ = strconv.Atoi(responseMap["nbTurns"])
			} else if responseMap["fromPage"] == "players_connexion" { //If the request is sended to server when the user arrives at the Players page : send the nbPlayers as response
				if err = websocket.Message.Send(ws, "{ \"nbPlayers\": "+strconv.Itoa(s.nbPlayers)+"}"); err != nil {
					fmt.Println("Can't send" + reply)
					break
				}
			} else if responseMap["fromPage"] == "players" { //If the request is sended to server when the user clicks the Suivant Button in the Players page : we should collect the players settings from the request
				// To Do

			} else if responseMap["fromPage"] == "simulation_connexion" { //If the request is sended to server when the user arrives at the Simulation page : start the simulation of the Galerapagos Game and send any messages that we want to display in the front
				//The real part (simulation of the agents) starts from here
				//To send a message to front, use this :
				/*
					if err = websocket.Message.Send(ws, reply); err != nil {
					fmt.Println("Can't send" + reply)
					break
				*/
			}
		}
	}
}

func main() {
	fmt.Println("start server")
	s := NewServeur()
	http.Handle("/", websocket.Handler(s.React))

	if err := http.ListenAndServe(":5000", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
