package Agents

import (
	"errors"
	"fmt"
	"math"
)

type Joueur struct {
	ID        int
	Nom       string
	Egoisme   int // de 0 à 10 si 10 : joueur très égoiste!
	Pecheur   bool
	Bricoleur bool
	EstMort   bool
	Prefs     []Joueur
}

//Constructeurs
func NewJoueurEgoiste(id int, nom string) Joueur {
	var prefs []Joueur
	return Joueur{id, nom, 10, false, false, false, prefs}
}

func NewJoueurAltruiste(id int, nom string) Joueur {
	var prefs []Joueur
	return Joueur{id, nom, 0, false, false, false, prefs}
}

//Vérification des attributs
func CheckJoueurs(joueurs []Joueur) error {
	if len(joueurs) < 3 || len(joueurs) > 12 {
		err := fmt.Sprint("Le nombre de joueurs doit être compris entre 3 et 12.")
		return errors.New(err)
	}
	for _, j := range joueurs {
		if j.Pecheur && j.Bricoleur {
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
	j.Bricoleur = true
	return j
}

func Meurt(j Joueur) Joueur {
	j.EstMort = true
	return j
}

func AuTourDe(joueurs []Joueur, premier Joueur) (j Joueur) {
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
			trouve = true
		}
	}
	if trouve == false {
		for i := 0; i < index; i++ {
			if !joueurs[i].EstMort && !trouve {
				j = joueurs[i]
				trouve = true
			}
		}
	}
	return j
}

func MakePrefs(j Joueur, autres []Joueur) (prefs []Joueur) {
	//TODO : selon le caractère de j et le caractère des autres joueurs faire une liste de pref
	return autres
}

func CalculPlaceRadeau(bois int) int {
	//6 bois pour une place
	return (bois / 6)
}

//Pour chaque action on calcule un score de 0 à 1 de faisabilité prenant en compte le plateau et le caractère du joueur
//l'action ayant le meilleure score est celle qui est effectuée par le joueur
func Joue(j Joueur, plateau Jeu, nbjoueurs int) Jeu {
	//TODO en fonction des caractéristiques du joueur et du plateau, calculer les score du joueur
	scorepeche := GetScorePeche(j, plateau, nbjoueurs)
	fmt.Println("pêcher :", scorepeche)
	scoreeau := GetScoreEau(j, plateau, nbjoueurs)
	fmt.Println("eau :", scoreeau)
	scorebois := GetScoreBois(j, plateau, nbjoueurs)
	fmt.Println("bois:", scorebois)
	scoremax := math.Max(math.Max(scorebois, scorepeche), scoreeau)
	fmt.Println("score max", scoremax)
	switch scoremax {
	case scorebois:
		plateau.StockBois += ConstructionRadeau(j)
		plateau.PlaceRadeau = CalculPlaceRadeau(plateau.StockBois)
	case scoreeau:
		plateau.StockEau += ChercherEau(plateau.Meteo)
	case scorepeche:
		plateau.StockNourriture += Pecher(j)
	}
	return plateau
}
