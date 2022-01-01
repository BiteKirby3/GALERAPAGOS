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

func GetScorePeche(j Joueur, plateau Jeu, nbjoueurs int) float64 {
	if nbjoueurs > plateau.StockNourriture {
		if j.Pecheur && !j.Bricoleur {
			return 1
		}
		if !j.Pecheur && !j.Bricoleur {
			if nbjoueurs < plateau.StockEau {
				return 1
			}
		}

	}
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

func IsRadeauDone(nbjoueurs int, places int) bool {
	if places > nbjoueurs {
		return true
	}
	return false
}

func AssezNourriture(nbjoueurs int, stocknourriture int) bool {
	if stocknourriture > nbjoueurs {
		return true
	}
	return false
}

func GetScoreEau(j Joueur, plateau Jeu, nbjoueurs int) int {
	//Joueur intelligent et altruiste
	score := 0
	if plateau.Meteo == 0 {
		score = 0
	} else {
		if plateau.Meteo == 1 {
			score += 3
		}
		if plateau.Meteo == 2 {
			score += 7
		}
		if plateau.Meteo == 3 {
			score += 10
		}
		//si le nb de joueur est supérieur au nb d'eau
		if nbjoueurs > plateau.StockEau {
			score += 10
			//si le joueur n'est pas bricoleur ou pêcheur, il va chercher de l'eau
			if !j.Bricoleur && j.Pecheur {
				score += 10
			} else
			//si le joueur est bricoleur MAIS il y a assez de place dans le radeau pour tous
			if IsRadeauDone(nbjoueurs, plateau.PlaceRadeau) && j.Bricoleur && !j.Pecheur {
				score += 10
			} else
			//si le joueur est pêcheur MAIS il y a assez de poisson pour tous
			if AssezNourriture(nbjoueurs, plateau.StockEau) && !j.Bricoleur && j.Pecheur {
				score += 10
			} else {
				score += 5
			}

		} else {
			//il y a assez d'eau pour faire boire les joueurs ce tour ci
			//ET je suis avancé dans le jeu (+ de la moitié des jours)
			if plateau.TourActuel/plateau.NbTour >= 2 {
				if !j.Bricoleur && j.Pecheur {
					score += 8
				}
			} else { //ET je suis au début du jeu

			}

		}

	}
	if score != 0 {
		return score / 3
	}
	/*
		à chacun des paramètres on a une note entre 0 et 10
				score eau dépend :
				- du stock d'eau / nb de tour restant
				- du nb de joueurs vivants
				- de la météo actuelle
				- des compétences du joueur
				autre chose??
			égoisme?
	*/
	return 0
}
func GetScoreBois(j Joueur, plateau Jeu, nbjoueurs int) float64 {

	/*
			score bois dépend :
			- dépend du nb de joueur en vie et du stock de bois
			- du tour actuel
			- si le joueur est un bricoleur
		égoisme?
	*/
	return 0
}
