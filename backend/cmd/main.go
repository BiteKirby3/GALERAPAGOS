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
				/*
					Dernier tour :
					Les joueurs encore en vie agissent en conséquence
				*/
				fmt.Println("Un ouragan ravage l'île, c'est votre dernière chance...")
				for i := 1; i <= nbjoueursvivants; i++ {
					fmt.Println(joueurjouant.Nom, ", c'est à toi !")
					//action?
					joueurjouant = Agents.AuTourDe(joueurs, joueurjouant)
				}
			} else {
				/*
					Tour de jeu classique
				*/
				for i := 1; i <= nbjoueursvivants; i++ {
					fmt.Println(joueurjouant.Nom, ", c'est à toi !")
					plateau = Agents.Joue(joueurjouant, plateau, nbjoueursvivants)
					Agents.AfficheJeu(plateau)
					joueurjouant = Agents.AuTourDe(joueurs, joueurjouant)
				}
				//Fin du tour, les joueurs ont tous joués, il faut maintenant retirer les stocks
				if nbjoueursvivants > plateau.StockEau {
					fmt.Println("Il n'y a pas assez d'eau, il faut éliminer", nbjoueursvivants-plateau.StockEau, "joueurs.")
					//vote et éliminiation
					nbjoueursvivants = nbjoueursvivants - (nbjoueursvivants - plateau.StockEau)

				}
				plateau.StockEau -= nbjoueursvivants
				if nbjoueursvivants > plateau.StockNourriture {
					fmt.Println("Il n'y a pas assez de nourriture, il faut éliminer", nbjoueursvivants-plateau.StockNourriture, "joueurs")
					//vote et élimination
					nbjoueursvivants = nbjoueursvivants - (nbjoueursvivants - plateau.StockEau)
				}
				plateau.StockNourriture -= nbjoueursvivants
			}

			//Survie des naufragés
			//récap et est-ce que les naufragés décide de voter pour tuer qqn?
			//si un joueur meurt, modifier la liste
		}
	}

}
