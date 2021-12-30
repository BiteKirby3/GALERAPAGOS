package Agents

import "math/rand"

func Pecher(j Joueur) int {
	poisson := rand.Intn(4) * 2
	if j.Pecheur {
		return 2 * poisson
	}
	return poisson
}

func ChercherEau(meteo int) int {
	return meteo
}

func ConstructionRadeau(j Joueur) int {
	if j.Bricoleur {
		return 1 + rand.Intn(4)
	}
	return rand.Intn(3) + 1
}

func GetScorePeche(j Joueur, plateau Jeu) float64 {
	/*
			POUR UN JOUEUR INTELLIGENT :
			score pêche dépend :
			- du stock de nourriture actuel
			- du nb de gens à nourrir (nb de joueurs en vie)
			- de si le joueur est pêcheur ou non
			POUR UN JOUEUR "BETE"
			score pêche dépend :
			- de si le joueur est pêcheur ou non
		égoisme?
	*/

	return 0
}
func GetScoreEau(j Joueur, plateau Jeu) float64 {
	/*
			score eau dépend :
			- du stock d'eau
			- du nb de joueurs vivants
			- de la météo actuelle
			autre chose??
		égoisme?
	*/
	return 0
}
func GetScoreBois(j Joueur, plateau Jeu) float64 {
	/*
			score bois dépend :
			- dépend du nb de joueur en vie et du stock de bois
			- du tour actuel
			- si le joueur est un bricoleur
		égoisme?
	*/
	return 0
}
