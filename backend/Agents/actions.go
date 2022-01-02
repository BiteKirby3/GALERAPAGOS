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

func AssezDeau(nbjoueurs int, stockeau int) bool {
	if stockeau > nbjoueurs {
		return true
	}
	return false
}

func GetScorePeche(j Joueur, plateau Jeu, nbjoueurs int) float64 {
	//Joueur intelligent et altruiste
	score := 0
	//il n'y a pas assez de nourriture pour tout le monde à ce tour
	if nbjoueurs > plateau.StockNourriture {
		score += 10
		//si le joueur est pêcheur, il est plus enclin à se dévouer à aller pêcher
		if j.Pecheur && !j.Bricoleur {
			score += 10
		}
		//si le joueur est bricoleur
		if j.Bricoleur {
			//ET le radeau est fait
			if IsRadeauDone(nbjoueurs, plateau.PlaceRadeau) {
				score += 10
			} else { //ET le radeau n'est pas fait
				score += 5
			}
		}
		//le joueur n'est ni pêcheur ni bricoleur
		if !j.Pecheur && !j.Bricoleur {
			if AssezDeau(nbjoueurs, plateau.StockEau) {
				if IsRadeauDone(nbjoueurs, plateau.PlaceRadeau) {
					score += 10
				} else {
					score += 8
				}
			} else {
				if IsRadeauDone(nbjoueurs, plateau.PlaceRadeau) {
					score += 5
				} else {
					score += 8
				}
			}
		}

	} else {
		//si j'ai assez de nourriture pour ce tour mais pas assez d'eau
		if nbjoueurs > plateau.StockEau {
			score += 0
			//capacité des joueurs
			if j.Pecheur {
				score += 5
			} else {
				score += 0
			}
		} else {
			//j'ai assez de nourriture et assez d'eau pour ce tour
			//prendre en compte l'avancement ds le jeu?
			score += 5
			if j.Pecheur {
				score += 5
			} else {
				score += 0
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
	if score == 0 {
		return 0
	}
	return float64(score / 2)
}

func GetScoreEau(j Joueur, plateau Jeu, nbjoueurs int) float64 {
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
				score += 7
			}

		} else {
			//il y a assez d'eau pour faire boire les joueurs ce tour ci
			//ET je suis avancé dans le jeu (+ de la moitié des jours)
			if plateau.TourActuel/plateau.NbTour >= 2 {
				score += 2
				if !j.Bricoleur && j.Pecheur {
					score += 8
				} else {
					score += 0
				}
			} else { //ET je suis au début du jeu
				score += 8
				if !j.Bricoleur && !j.Pecheur {
					score += 8
				} else {
					score += 4
				}
			}

		}

	}
	if score == 0 {
		return 0
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
	return float64(score / 3)
}
func GetScoreBois(j Joueur, plateau Jeu, nbjoueurs int) float64 {
	score := 0
	if !IsRadeauDone(nbjoueurs, plateau.PlaceRadeau) {
		if AssezDeau(nbjoueurs, plateau.StockEau) {
			if AssezNourriture(nbjoueurs, plateau.StockNourriture) {
				score += 8
				if j.Bricoleur {
					score += 10
				} else {
					score += 8
				}
			} else {
				score += 3
				if j.Pecheur {
					score += 0
				} else if j.Bricoleur {
					score += 10
				} else {
					score += 3
				}
			}
		} else {
			if AssezNourriture(nbjoueurs, plateau.StockNourriture) {
				score += 3
			} else {
				if j.Pecheur {
					score += 0
				} else {
					score += 2
				}
			}
		}

	} else {
		score += 0
	}
	/*
			score bois dépend :
			- dépend du nb de joueur en vie et du stock de bois
			- du tour actuel
			- si le joueur est un bricoleur
			- je suis avancé ds le jeu
		égoisme?
	*/
	if score == 0 {
		return 0
	}
	return float64(score / 2)
}
