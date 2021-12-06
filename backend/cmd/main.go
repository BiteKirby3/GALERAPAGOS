package main

import (
	"fmt"
	"gitlab.utc.fr/ia04_group/galerapagos_ia04/backend/Agents"
)

func main() {
	var j1, j2, j3 Agents.Joueur
	j1 = Agents.NewJoeurEgoiste("1")
	j2 = Agents.NewJoueurAltruiste("2")
	j3 = Agents.NewJoueurAltruiste("3")
	var joueurs []Agents.Joueur
	joueurs = append(joueurs, j1)
	joueurs = append(joueurs, j2)
	joueurs = append(joueurs, j3)
	err := Agents.CheckJoueurs(joueurs)
	fmt.Println(err)
}
