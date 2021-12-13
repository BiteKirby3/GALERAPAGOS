package Agents

import "math/rand"

func pecher(j Joueur) int {
	poisson := rand.Intn(4) * 2
	if j.pecheur {
		return 2 * poisson
	}
	return poisson
}

func chercherEau(meteo int) int {
	return meteo
}

func constructionRadeau(j Joueur) int {
	if j.bricoleur {
		return 1 + rand.Intn(4)
	}
	return rand.Intn(3) + 1
}
