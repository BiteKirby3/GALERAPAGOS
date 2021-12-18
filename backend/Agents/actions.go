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
