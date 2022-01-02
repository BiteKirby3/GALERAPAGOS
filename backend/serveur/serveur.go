package main

import (
	"encoding/json"
	"fmt"
	"gitlab.utc.fr/ia04_group/galerapagos_ia04/backend/Agents"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"sync"
	"time"

	"golang.org/x/net/websocket"
)

type Player struct {
	Id           int
	Name         string
	Role         string
	Selfishness  int
	Intelligence int
}

type FormPlayers struct {
	FromPage string
	Players  []Player
}

type Serveur struct {
	nbPlayers   int
	nbTurns     int
	listPlayers []Agents.Joueur
}

var messages []Agents.Message
var l sync.Mutex

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
			//responseMap := make(map[string]string)
			var responseMap map[string]interface{}
			err := json.Unmarshal([]byte(reply), &responseMap)
			if err != nil {
				panic(err)
			}

			if responseMap["fromPage"] == "home" { //If the request is sended from the homePage :  we should store the nbPlayers and nbTurns value.
				s.nbPlayers, _ = strconv.Atoi(responseMap["nbPlayers"].(string))
				s.nbTurns, _ = strconv.Atoi(responseMap["nbTurns"].(string))
			} else if responseMap["fromPage"] == "players_connexion" { //If the request is sended to server when the user arrives at the Players page : send the nbPlayers as response
				if err = websocket.Message.Send(ws, "{ \"nbPlayers\": "+strconv.Itoa(s.nbPlayers)+"}"); err != nil {
					fmt.Println("Can't send" + reply)
					break
				}
			} else if responseMap["fromPage"] == "players" { //If the request is sended to server when the user clicks the Suivant Button in the Players page : we should collect the players settings from the request
				var responsePlayers FormPlayers
				err := json.Unmarshal([]byte(reply), &responsePlayers)
				if err != nil {
					panic(err)
				}
				for _, val := range responsePlayers.Players {
					var j Agents.Joueur
					if val.Role == "fisherman" {
						j = Agents.NewJoueur(val.Id, val.Selfishness, val.Intelligence, true, false, val.Name)
					} else if val.Role == "handyman" {
						j = Agents.NewJoueur(val.Id, val.Selfishness, val.Intelligence, false, true, val.Name)
					} else {
						j = Agents.NewJoueur(val.Id, val.Selfishness, val.Intelligence, false, false, val.Name)
					}
					s.listPlayers = append(s.listPlayers, j)
				}
				fmt.Println(responsePlayers)

			} else if responseMap["fromPage"] == "simulation_connexion" { //If the request is sended to server when the user arrives at the Simulation page : start the simulation of the Galerapagos Game and send any messages that we want to display in the front
				countMessageSend := 0
				messages = make([]Agents.Message, 0)
				go Simulation(s.listPlayers, s.nbTurns, s.nbPlayers)
				for {
					if countMessageSend < len(messages) {
						sendMessage(messages[countMessageSend], ws, err)
						countMessageSend++
					}
				}
			}
		}
	}
}

func addMessage(m Agents.Message) {
	l.Lock()
	messages = append(messages, m)
	l.Unlock()
}

func sendMessage(m Agents.Message, ws *websocket.Conn, err error) {
	fmt.Println("Send message" + m.Description)
	marshalledMessage, errMarshal := json.Marshal(m)
	if errMarshal != nil {
		fmt.Println("Can't marshal message")
	}
	if err = websocket.Message.Send(ws, string(marshalledMessage)); err != nil {
		fmt.Println("Can't send" + m.Description)
	}
}

func InitCompteur(plateauInitial Agents.Jeu, joueurs []Agents.Joueur, log string) {
	message := Agents.Message{joueurs, plateauInitial, log, "INIT"}
	addMessage(message)
}

func Simulation(joueurs []Agents.Joueur, nbTours int, nbJoueurs int) {
	//INITIALISATION
	rand.Seed(time.Now().UnixNano())
	plateau := Agents.InitJeu(nbJoueurs, nbTours)
	//LANCEMENT DU JEU
	nbJoueurVivants := nbJoueurs

	for _, val := range joueurs {
		val.Prefs = Agents.MakePrefs(val, joueurs)
	}
	InitCompteur(plateau, joueurs, "Le jeu commence")
	/*
		DEROULEMENT DE LA PARTIE
	*/
	premierjoueur := joueurs[0] //on détermine le premier joueur
	// Le jeu continue tant que pas d'ouragan et au moins 1 joueurs sont en vie
	for plateau.TourActuel != plateau.NbTour && nbJoueurVivants >= 1 {
		var profile [][]int
		//Tirage de la carte Météo et incrémentation d'un tour
		plateau = Agents.NewDay(plateau)
		fmt.Println("_______Tour : ", plateau.TourActuel)
		fmt.Println("Météo : ", plateau.Meteo)
		//Changement du premier joueur
		if plateau.TourActuel == 1 {
			premierjoueur = joueurs[0]
		} else {
			premierjoueur = Agents.AuTourDe(joueurs, premierjoueur)
		}
		fmt.Println("Pour ce tour c'est ", premierjoueur.Nom, " qui commence")

		joueurJouant := premierjoueur
		//Action des joueurs
		for i := 1; i <= nbJoueurVivants; i++ {
			fmt.Println(joueurJouant.Nom, ", c'est à toi !")
			profile = append(profile, joueurJouant.Prefs)
			plateau = Agents.Joue(plateau, joueurJouant, nbJoueurVivants)
			joueurJouant = Agents.AuTourDe(joueurs, joueurJouant)
		}

		//Vote
		if (plateau.StockEau > 0) && (plateau.StockNourriture > 0) {
			if plateau.StockEau < nbJoueurVivants {
				nbDePersonneATue := nbJoueurVivants - plateau.StockEau
				IDMort := Agents.Vote(profile, nbDePersonneATue)
				for i := 1; i <= nbJoueurVivants; i++ {
					for _, val := range IDMort {
						if joueurs[i].ID == val {
							joueurs[i].EstMort = true
							fmt.Println(joueurs[i].Nom + " a été tué")
							nbJoueurVivants--
						}
					}
				}
			}
			plateau.StockEau = plateau.StockEau - nbJoueurVivants

			if plateau.StockNourriture < nbJoueurVivants {
				nbDePersonnesATue := nbJoueurVivants - plateau.StockNourriture
				IDMort := Agents.Vote(profile, nbDePersonnesATue)
				for i := 1; i <= nbJoueurVivants; i++ {
					for _, val := range IDMort {
						if joueurs[i].ID == val {
							joueurs[i].EstMort = true
							fmt.Println(joueurs[i].Nom + " a été tué")
							nbJoueurVivants--
						}
					}
				}
			}

			plateau.StockNourriture = plateau.StockNourriture - nbJoueurVivants

			if (plateau.PlaceRadeau < nbJoueurVivants) && (plateau.Meteo == 4) {
				nbDePersonneATue := nbJoueurVivants - plateau.PlaceRadeau
				IDMort := Agents.Vote(profile, nbDePersonneATue)
				for i := 1; i <= nbJoueurVivants; i++ {
					for _, val := range IDMort {
						if joueurs[i].ID == val {
							joueurs[i].EstMort = true
							fmt.Println(joueurs[i].Nom + " a été tué")
							nbJoueurVivants--
						}
					}
				}
			}

			//Survie des naufragés
			if plateau.Meteo == 4 {
				fmt.Println("Un ouragan a ravagé l'île...")
				for _, val := range joueurs {
					if !val.EstMort {
						fmt.Println(val.Nom + " a survécu")
					}
				}
			}
		} else {
			fmt.Println("Les joueurs restants sont mort de faim et/ou de soif")
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
