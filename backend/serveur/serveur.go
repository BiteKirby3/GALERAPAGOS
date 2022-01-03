package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"sync"
	"time"

	"gitlab.utc.fr/ia04_group/galerapagos_ia04/backend/Agents"

	"golang.org/x/net/websocket"
)

var durationSleep = 2 * time.Second

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
					} else {
						time.Sleep(100 * time.Millisecond)
						var desc []Agents.Joueur
						var jeu Agents.Jeu
						m := Agents.Message{desc, jeu, "Message to keep connection open", "empty"}
						sendMessage(m, ws, err)
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
	marshalledMessage, errMarshal := json.Marshal(m)
	if errMarshal != nil {
		fmt.Println("Can't marshal message")
	}
	if err = websocket.Message.Send(ws, string(marshalledMessage)); err != nil {
		fmt.Println("Can't send" + m.Description)
	}
}

func InitCompteur(plateauInitial Agents.Jeu, joueurs []Agents.Joueur) {
	message := Agents.Message{joueurs, plateauInitial, "Le jeu commence", "gameStart"}
	addMessage(message)
	time.Sleep(durationSleep)
}

func FirstPlayer(plateauInitial Agents.Jeu, joueurs []Agents.Joueur, i int) {
	message := Agents.Message{joueurs, plateauInitial, joueurs[i].Nom + " commence la partie", "firstPlayer"}
	addMessage(message)
	time.Sleep(durationSleep)
}

func InitMeteo(plateauInitial Agents.Jeu, joueurs []Agents.Joueur, i int) {
	var message Agents.Message
	if i == 4 {
		message = Agents.Message{joueurs, plateauInitial, "Météo : Ouragan (eau : 4)", "meteo"}
	} else if i == 3 {
		message = Agents.Message{joueurs, plateauInitial, "Météo : Averse (eau : 3)", "meteo"}
	} else if i == 2 {
		message = Agents.Message{joueurs, plateauInitial, "Météo : Pluvieux (eau : 2)", "meteo"}
	} else if i == 1 {
		message = Agents.Message{joueurs, plateauInitial, "Météo : Nuageux (eau : 1)", "meteo"}
	} else {
		message = Agents.Message{joueurs, plateauInitial, "Météo : Sécheresse (eau : 0)", "meteo"}
	}
	addMessage(message)
	time.Sleep(durationSleep)
}

func ActionPlayer(plateauInitial Agents.Jeu, joueurs []Agents.Joueur, id int, typeAction int, nb int) {
	nbString := strconv.Itoa(nb)
	s := joueurs[id].Nom + " a récupéré " + nbString
	var message Agents.Message
	if typeAction == 0 {
		message = Agents.Message{joueurs, plateauInitial, s + " gourdes d'eau.", "action"}
	} else if typeAction == 1 {
		message = Agents.Message{joueurs, plateauInitial, s + " poissons.", "action"}
	} else {
		message = Agents.Message{joueurs, plateauInitial, s + " planches de bois.", "action"}
	}
	addMessage(message)
	time.Sleep(durationSleep)
}

func NextPlayer(plateauInitial Agents.Jeu, joueurs []Agents.Joueur, i int) {
	message := Agents.Message{joueurs, plateauInitial, "c'est à " + joueurs[i].Nom + " de jouer", "nextPlayer"}
	addMessage(message)
	time.Sleep(durationSleep)
}

func DeathPlayer(plateauInitial Agents.Jeu, joueurs []Agents.Joueur, i int) {
	message := Agents.Message{joueurs, plateauInitial, joueurs[i].Nom + " est mort", "death"}
	addMessage(message)
	time.Sleep(durationSleep)
}

func SurvivePlayer(plateauInitial Agents.Jeu, joueurs []Agents.Joueur, i int) {
	message := Agents.Message{joueurs, plateauInitial, joueurs[i].Nom + " a survécu", "alive"}
	addMessage(message)
	time.Sleep(durationSleep)
}

func GameEnd(plateauInitial Agents.Jeu, joueurs []Agents.Joueur, win bool) {
	var message Agents.Message
	if !win {
		message = Agents.Message{joueurs, plateauInitial, "C'est la fin du jeu! Les joueurs restants sont mort de faim et/ou de soif", "gameEnd"}
	} else {
		message = Agents.Message{joueurs, plateauInitial, "C'est la fin du jeu! ", "gameEnd"}
	}
	addMessage(message)
	time.Sleep(durationSleep)
}

func RoundEnd(plateauInitial Agents.Jeu, joueurs []Agents.Joueur) {
	message := Agents.Message{joueurs, plateauInitial, "C'est la fin du tour. ", "roundEnd"}
	addMessage(message)
	time.Sleep(durationSleep)
}

func Simulation(joueurs []Agents.Joueur, nbTours int, nbJoueurs int) {
	//INITIALISATION
	rand.Seed(time.Now().UnixNano())
	plateau := Agents.InitJeu(nbJoueurs, nbTours)
	//LANCEMENT DU JEU
	nbJoueurVivants := nbJoueurs

	for _, val := range joueurs {
		Agents.MakePrefs(val, &joueurs)
	}
	InitCompteur(plateau, joueurs)
	/*
		DEROULEMENT DE LA PARTIE
	*/
	indicePremier := 0
	premierjoueur := joueurs[0] //on détermine le premier joueur
	FirstPlayer(plateau, joueurs, indicePremier)
	// Le jeu continue tant que pas d'ouragan et au moins 1 joueurs sont en vie
	for plateau.TourActuel != plateau.NbTour && nbJoueurVivants >= 1 {
		var profile [][]int
		//Tirage de la carte Météo et incrémentation d'un tour
		plateau = Agents.NewDay(plateau)
		fmt.Println("_______Tour : ", plateau.TourActuel)
		fmt.Println("Météo : ", plateau.Meteo)
		InitMeteo(plateau, joueurs, plateau.Meteo)
		//Changement du premier joueur
		if plateau.TourActuel == 1 {
			premierjoueur = joueurs[0]
			indicePremier = 0
		} else {
			premierjoueur, indicePremier = Agents.AuTourDe(joueurs, premierjoueur)
			FirstPlayer(plateau, joueurs, indicePremier)
		}
		fmt.Println("Pour ce tour c'est ", premierjoueur.Nom, " qui commence")

		joueurJouant := premierjoueur
		indiceJoueur := indicePremier
		typeAction := 0
		nbRecolte := 0
		//Action des joueurs
		for i := 0; i < nbJoueurVivants; i++ {
			fmt.Println(joueurJouant.Nom, ", c'est à toi !")
			profile = append(profile, joueurJouant.Prefs)
			plateau, typeAction, nbRecolte = Agents.Joue(plateau, joueurJouant, nbJoueurVivants)
			ActionPlayer(plateau, joueurs, indiceJoueur, typeAction, nbRecolte)
			joueurJouant, indiceJoueur = Agents.AuTourDe(joueurs, joueurJouant)
			if i != nbJoueurVivants-1 {
				NextPlayer(plateau, joueurs, indiceJoueur)
			}
		}
		//Vote
		if (plateau.StockEau > 0) && (plateau.StockNourriture > 0) {
			if plateau.StockEau < nbJoueurVivants {
				nbDePersonneATue := nbJoueurVivants - plateau.StockEau
				IDMort := Agents.Vote(profile, nbDePersonneATue)
				for i := 0; i < nbJoueurVivants; i++ {
					for _, val := range IDMort {
						if joueurs[i].ID == val {
							joueurs[i].EstMort = true
							fmt.Println(joueurs[i].Nom + " a été tué")
							DeathPlayer(plateau, joueurs, i)
							nbJoueurVivants--
						}
					}
				}
			}
			plateau.StockEau = plateau.StockEau - nbJoueurVivants

			if plateau.StockNourriture < nbJoueurVivants {
				nbDePersonnesATue := nbJoueurVivants - plateau.StockNourriture
				IDMort := Agents.Vote(profile, nbDePersonnesATue)
				for i := 0; i < nbJoueurVivants; i++ {
					for _, val := range IDMort {
						if joueurs[i].ID == val {
							joueurs[i].EstMort = true
							fmt.Println(joueurs[i].Nom + " a été tué")
							DeathPlayer(plateau, joueurs, i)
							nbJoueurVivants--
						}
					}
				}
			}

			plateau.StockNourriture = plateau.StockNourriture - nbJoueurVivants

			if (plateau.PlaceRadeau <= 0) && (plateau.Meteo == 4) {
				for i, val := range joueurs {
					if val.EstMort == false {
						joueurs[i].EstMort = true
						fmt.Println(joueurs[i].Nom + " a été tué")
						DeathPlayer(plateau, joueurs, i)
						nbJoueurVivants--
					}
				}
			} else if (plateau.PlaceRadeau < nbJoueurVivants) && (plateau.Meteo == 4) {
				nbDePersonneATue := nbJoueurVivants - plateau.PlaceRadeau
				IDMort := Agents.Vote(profile, nbDePersonneATue)
				for i := 0; i < nbJoueurVivants; i++ {
					for _, val := range IDMort {
						if joueurs[i].ID == val {
							joueurs[i].EstMort = true
							fmt.Println(joueurs[i].Nom + " a été tué")
							DeathPlayer(plateau, joueurs, i)
							nbJoueurVivants--
						}
					}
				}
			}

			RoundEnd(plateau, joueurs)

			//Survie des naufragés
			if plateau.Meteo == 4 {
				fmt.Println("Un ouragan a ravagé l'île...")
				for id, val := range joueurs {
					if !val.EstMort {
						fmt.Println(val.Nom + " a survécu")
						SurvivePlayer(plateau, joueurs, id)
					}
				}
				GameEnd(plateau, joueurs, true)
			}
		} else {
			fmt.Println("Les joueurs restants sont mort de faim et/ou de soif")
			GameEnd(plateau, joueurs, false)
		}
	}
	GameEnd(plateau, joueurs, true)
}

func main() {
	fmt.Println("start server")
	s := NewServeur()
	http.Handle("/", websocket.Handler(s.React))

	if err := http.ListenAndServe(":5000", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
