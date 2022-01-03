package Agents

import (
	"errors"
	"fmt"
	"math"
	"math/rand"
)

type Joueur struct {
	ID           int
	Nom          string
	Egoisme      int // de 0 à 10 si 10 : joueur très égoiste!
	Intelligence int // de 0 à 10 si 10 : joueur très intelligent!
	Pecheur      bool
	Bucheron     bool
	EstMort      bool
	Prefs        []int
}

//Constructeurs
func NewJoueur(id int, intelligence int, egoisme int, p bool, b bool, nom string) Joueur {
	var prefs []int
	return Joueur{id, nom, egoisme, intelligence, p, b, false, prefs}
}

//Vérification des attributs
func CheckJoueurs(joueurs []Joueur) error {
	if len(joueurs) < 3 || len(joueurs) > 12 {
		err := fmt.Sprint("Le nombre de joueurs doit être compris entre 3 et 12.")
		return errors.New(err)
	}
	for _, j := range joueurs {
		if j.Pecheur && j.Bucheron {
			err := fmt.Sprint("Le joueur ", j.ID, " ne peut pas être pêcheur ET bricoleur.")
			return errors.New(err)
		}
		if j.Egoisme > 10 || j.Egoisme < 0 {
			err := fmt.Sprint("Le niveau d'égoïsme du joueur ", j.ID, " doit être compris entre 0 et 10.")
			return errors.New(err)
		}
	}
	return nil
}

//Changement de caractéristiques
func DevientPecheur(j Joueur) Joueur {
	j.Pecheur = true
	return j
}

func DevientBricoleur(j Joueur) Joueur {
	j.Bucheron = true
	return j
}

func Meurt(j Joueur) Joueur {
	j.EstMort = true
	return j
}

func AuTourDe(joueurs []Joueur, premier Joueur) (j Joueur, pos int) {
	index := 0
	for i, j := range joueurs {
		if premier.ID == j.ID {
			index = i
		}
	}
	trouve := false
	for i := index + 1; i < len(joueurs); i++ {
		if !joueurs[i].EstMort && !trouve {
			j = joueurs[i]
			pos = i
			trouve = true
		}
	}
	if trouve == false {
		for i := 0; i < index; i++ {
			if !joueurs[i].EstMort && !trouve {
				j = joueurs[i]
				pos = i
				trouve = true
			}
		}
	}
	return j, pos
}

func MakePrefs(j Joueur, autres *[]Joueur) {
	mapScore := make(map[int]int)
	joueurIndice := 0
	for i, val := range *autres {
		if val.ID != j.ID {
			mapScore[val.ID] = rand.Intn(3)
			if val.Pecheur == true {
				mapScore[val.ID] = mapScore[val.ID] + 1
			} else if val.Bucheron == true {
				mapScore[val.ID] = mapScore[val.ID] + 1
			}
		} else {
			joueurIndice = i
		}
	}

	s := MaxCount(mapScore)
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}

	(*autres)[joueurIndice].Prefs = s
}

//Pour chaque action on calcule un score de 0 à 1 de faisabilité prenant en compte le plateau et le caractère du joueur
//l'action ayant le meilleure score est celle qui est effectuée par le joueur
func Joue(plateau Jeu, j Joueur, nbjoueurs int) (Jeu, int, int) {
	//TODO en fonction des caractéristiques du joueur et du plateau, calculer les score du joueur
	typeAction := 0
	nb := 0
	scorepeche := float64(GetScorePeche(j, plateau, nbjoueurs))
	scoreeau := float64(GetScoreEau(j, plateau, nbjoueurs))
	scorebois := float64(GetScoreBois(j, plateau, nbjoueurs))
	scoremax := math.Max(math.Max(scorebois, scorepeche), scoreeau)
	switch scoremax {
	case scorebois:
		typeAction = 2
		nb = ConstructionRadeau(j)
		plateau.StockBois += nb
		if plateau.StockBois >= 6 {
			nbPlaces := plateau.StockBois / 6
			plateau.PlaceRadeau += nbPlaces
			plateau.StockBois = plateau.StockBois - (nbPlaces * 6)
		}
	case scoreeau:
		typeAction = 0
		nb = ChercherEau(plateau.Meteo)
		plateau.StockEau += nb
	case scorepeche:
		typeAction = 1
		nb = Pecher(j)
		plateau.StockNourriture += nb
	}
	return plateau, typeAction, nb
}
