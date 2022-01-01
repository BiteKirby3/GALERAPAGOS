package Agents

import (
	"math/rand"
)

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
	if j.Bucheron {
		return 1 + rand.Intn(4)
	}
	return rand.Intn(3) + 1
}

func GetScorePeche(j Joueur, plateau Jeu, nbjoueurs int) int {
	deIntelligence := rand.Intn(11)
	if deIntelligence <= j.Intelligence {
		if nbjoueurs > plateau.StockNourriture {
			if j.Pecheur && !j.Bucheron {
				return 2
			}
			if !j.Pecheur && !j.Bucheron {
				if nbjoueurs < plateau.StockEau {
					return 1
				}
			}

		}
		return 0
	} else {
		return rand.Intn(3)
	}
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
	deIntelligence := rand.Intn(11)
	if deIntelligence <= j.Intelligence {
		if plateau.Meteo == 0 {
			return 0
		}
		if plateau.StockEau > 2*nbjoueurs {
			return 0
		}
		if (plateau.StockEau > nbjoueurs) && (plateau.Meteo >= 2) {
			return 1
		}
		if plateau.StockEau < nbjoueurs {
			return 2
		}
	} else {
		return rand.Intn(3)
	}
	return 0
}
func GetScoreBois(j Joueur, plateau Jeu, nbjoueurs int) int {
	deIntelligence := rand.Intn(11)
	if deIntelligence <= j.Intelligence {
		if nbjoueurs > plateau.StockBois {
			if !j.Pecheur && j.Bucheron {
				return 2
			}
			if !j.Pecheur && !j.Bucheron {
				if nbjoueurs < plateau.StockNourriture {
					return 1
				}
			}
		}
		return 0
	} else {
		return rand.Intn(3)
	}
	return 0
}
