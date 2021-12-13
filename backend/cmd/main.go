package main

import (
	"fmt"

	"gitlab.utc.fr/ia04_group/galerapagos_ia04/backend/Agents"
)

func main() {
	plateau := Agents.InitJeu(3, 10)
	var j1, j2, j3 Agents.Joueur
	j1 = Agents.NewJoeurEgoiste(1, "Pierre")
	j2 = Agents.NewJoueurAltruiste(2, "Paul")
	j3 = Agents.NewJoueurAltruiste(3, "Jacques")
	var joueurs []Agents.Joueur
	joueurs = append(joueurs, j1)
	joueurs = append(joueurs, j2)
	joueurs = append(joueurs, j3)
	err := Agents.CheckJoueurs(joueurs)
	fmt.Println(err)
	err2 := Agents.CheckTours(plateau.NbTour)
	fmt.Println(err2)
}
