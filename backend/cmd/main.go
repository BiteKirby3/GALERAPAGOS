package main

import (
	"fmt"
	"gitlab.utc.fr/ia04_group/galerapagos_ia04/backend/Agents"
)

func main() {
	//INITIALISATION
	plateau := Agents.InitJeu(4, 10)
	var j1, j2, j3, j4 Agents.Joueur
	j1 = Agents.NewJoueurEgoiste(1, "Pierre")
	j2 = Agents.NewJoueurAltruiste(2, "Paul")
	j3 = Agents.NewJoueurAltruiste(3, "Jacques")
	j4 = Agents.NewJoueurEgoiste(4, "Test")
	var joueurs []Agents.Joueur
	joueurs = append(joueurs, j1)
	joueurs = append(joueurs, j2)
	joueurs = append(joueurs, j3)
	joueurs = append(joueurs, j4)

	//LANCEMENT DU JEU
	nbjoueursvivants := len(joueurs)
	err := Agents.CheckJoueurs(joueurs)
	err2 := Agents.CheckTours(plateau.NbTour)
	if err != nil || err2 != nil {
		print("Impossible de lancer le jeu...")
	} else {
		/*
			DEROULEMENT DE LA PARTIE
		*/
		premierjoueur := j1 //on détermine le premier joueur

		// Le jeu continue tant que pas d'ouragan et au moins 2 joueurs sont en vie
		for plateau.TourActuel != plateau.NbTour && nbjoueursvivants > 2 {
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

			joueurjouant := premierjoueur
			//Action des joueurs
			if plateau.Meteo == 4 {
				fmt.Println("Un ouragan ravage l'île, c'est votre dernière chance...")
				for i := 1; i <= nbjoueursvivants; i++ {
					fmt.Println(joueurjouant.Nom, ", c'est à toi !")
					plateau = Agents.Joue(plateau, joueurjouant)
					joueurjouant = Agents.AuTourDe(joueurs, joueurjouant)
				}
			} else {
				/*
					Dernier tour :
					Les joueurs encore en vie agissent en conséquence
				*/
				for i := 1; i <= nbjoueursvivants; i++ {
					fmt.Println(joueurjouant.Nom, ", c'est à toi !")
					//TODO : SES ACTIONS au dernier tour
					joueurjouant = Agents.AuTourDe(joueurs, joueurjouant)
				}
			}

			//Survie des naufragés
			//récap et est-ce que les naufragés décide de voter pour tuer qqn?
			//si un joueur meurt, modifier la liste
		}
	}

}
