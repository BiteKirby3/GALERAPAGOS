package Agents

import (
	"errors"
	"fmt"
)

type Joueur struct {
	ID        int
	nom       string
	egoisme   int // de 0 à 10 si 10 : joueur très égoiste!
	pecheur   bool
	bricoleur bool
	estMort   bool
	prefs     []Joueur
}

func MakePrefs(j Joueur, autres []Joueur) (prefs []Joueur) {
	//TODO : selon le caractère de j et le caractère des autres joueurs faire une liste de pref
	return autres
}

func NewJoeurEgoiste(id int, nom string) Joueur {
	var prefs []Joueur
	return Joueur{id, nom, 10, false, false, false, prefs}
}

func NewJoueurAltruiste(id int, nom string) Joueur {
	var prefs []Joueur
	return Joueur{id, nom, 0, false, false, false, prefs}
}

func CheckJoueurs(joueurs []Joueur) error {
	if len(joueurs) < 3 || len(joueurs) > 12 {
		err := fmt.Sprint("Le nombre de joueurs doit être compris entre 3 et 12.")
		return errors.New(err)
	}
	for _, j := range joueurs {
		if j.pecheur && j.bricoleur {
			err := fmt.Sprint("Le joueur ", j.ID, " ne peut pas être pêcheur ET bricoleur.")
			return errors.New(err)
		}
		if j.egoisme > 10 || j.egoisme < 0 {
			err := fmt.Sprint("Le niveau d'égoïsme du joueur ", j.ID, " doit être compris entre 0 et 10.")
			return errors.New(err)
		}
	}
	return nil
}

func DevientPecheur(j Joueur) Joueur {
	j.pecheur = true
	return j
}

func DevientBricoleur(j Joueur) Joueur {
	j.bricoleur = true
	return j
}
