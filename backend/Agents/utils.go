package Agents

import (
	"errors"
	"fmt"
	"math/rand"
)

type Jeu struct {
	Meteo           int
	StockEau        int
	StockNourriture int
	StockBois       int
	PlaceRadeau     int
	NbTour          int
	TourActuel      int
}

func InitJeu(nbJ int, nbTour int) Jeu {
	nourriture := 0
	eau := 0
	switch nbJ {
	case 3:
		nourriture = 5
		eau = 6
	case 4:
		nourriture = 7
		eau = 8
	case 5:
		nourriture = 8
		eau = 10
	case 6:
		nourriture = 10
		eau = 12
	case 7:
		nourriture = 12
		eau = 14
	case 8:
		nourriture = 13
		eau = 16
	case 9:
		nourriture = 15
		eau = 18
	case 10:
		nourriture = 16
		eau = 20
	case 11:
		nourriture = 18
		eau = 22
	case 12:
		nourriture = 20
		eau = 24
	}
	return Jeu{rand.Intn(4), eau, nourriture, 0, 0, nbTour, 1}
}

func CheckTours(nbTours int) error {
	if nbTours > 12 || nbTours < 6 {
		err := fmt.Sprint("Le nombre de tours doit être compris entre 6 et 12.")
		return errors.New(err)
	}
	return nil
}
