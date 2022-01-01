package main

import (
	"fmt"
	"gitlab.utc.fr/ia04_group/galerapagos_ia04/backend/Agents"
	"math/rand"
	"time"
)

func main() {
	//INITIALISATION
	plateau := Agents.InitJeu(4, 10)
	rand.Seed(time.Now().UnixNano())
	var j1, j2, j3, j4 Agents.Joueur
	j1 = Agents.NewJoueur(1, 2, "Pierre")
	j2 = Agents.NewJoueur(2, 6, "Paul")
	j3 = Agents.NewJoueur(3, 8, "Jacques")
	j4 = Agents.NewJoueur(4, 4, "Test")
	var joueurs []Agents.Joueur
	joueurs = append(joueurs, j1)
	joueurs = append(joueurs, j2)
	joueurs = append(joueurs, j3)
	joueurs = append(joueurs, j4)

	//LANCEMENT DU JEU
	nbJoueurVivants := len(joueurs)
	err := Agents.CheckJoueurs(joueurs)
	err2 := Agents.CheckTours(plateau.NbTour)

	joueurs[0].Prefs = Agents.MakePrefs(j1, joueurs)
	joueurs[1].Prefs = Agents.MakePrefs(j2, joueurs)
	joueurs[2].Prefs = Agents.MakePrefs(j3, joueurs)
	joueurs[3].Prefs = Agents.MakePrefs(j4, joueurs)

	if err != nil || err2 != nil {
		print("Impossible de lancer le jeu...")
	} else {
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
}
