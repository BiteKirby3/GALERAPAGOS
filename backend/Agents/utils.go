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

type Message struct {
	ListJoueurs []Joueur
	Plateau     Jeu
	Description string
	TypeEvent   string
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

	return Jeu{rand.Intn(4), eau, nourriture, 0, 0, nbTour, 0}
}

func NewDay(j Jeu) Jeu {
	j.TourActuel = j.TourActuel + 1
	if j.NbTour == j.TourActuel {
		j.Meteo = 4
	} else {
		j.Meteo = rand.Intn(4)
	}
	return j
}

func CheckTours(nbTours int) error {
	if nbTours > 12 || nbTours < 6 {
		err := fmt.Sprint("Le nombre de tours doit être compris entre 6 et 12.")
		return errors.New(err)
	}
	return nil
}

func Vote(profile [][]int, nbDePersonneATue int) []int {
	bestAlts, _ := MajoritySCF(profile)
	return bestAlts[:nbDePersonneATue]
}

func MajoritySWF(p [][]int) (count map[int]int, err error) {
	count = make(map[int]int)
	for i := 0; i < len(p); i++ {
		alterPref := p[i][0]
		_, ok := count[alterPref]
		if ok {
			count[alterPref] = count[alterPref] + 1
		} else {
			count[alterPref] = 1
		}
	}

	//add the tiebreak part
	var alts []int
	for alt := range count {
		alts = append(alts, alt)
	}
	countTb := TieBreak(alts)
	const N = 100
	for alt := range count {
		count[alt] = N*count[alt] + countTb[alt]
	}
	return count, err
}

func MajoritySCF(p [][]int) (bestAlts []int, err error) {
	count, _ := MajoritySWF(p)
	bestAlts = MaxCount(count)
	return bestAlts, err
}

func TieBreak(alts []int) map[int]int {
	n := len(alts)
	var slice = make([]int, n)
	for i := 0; i < n; i++ {
		slice[i] = i
	}
	k := 100 //repeated k times change
	for {
		r1 := rand.Intn(n)
		r2 := rand.Intn(n)
		temp := slice[r1]
		slice[r1] = slice[r2]
		slice[r2] = temp
		k--
		if k <= 0 {
			break
		}
	}
	count := make(map[int]int)
	for i := 0; i < n; i++ {
		count[alts[i]] = slice[i]
	}

	return count
}

func MaxCount(count map[int]int) (bestAlts []int) {
	bestAlts = make([]int, 0)
	for len(bestAlts) < len(count) {
		idMax := -1
		valMax := -1
		for k, v := range count {
			if (valMax < v) && (!Contains(bestAlts, k)) {
				idMax = k
				valMax = v
			}
		}
		bestAlts = append(bestAlts, idMax)
	}
	return bestAlts
}

func Contains(list []int, a int) bool {
	for _, val := range list {
		if val == a {
			return true
		}
	}
	return false
}
