# Galerapagos_IA04

## Membres du groupe
- Julie Szerbojm
- Romane Guari      
- Sihan Xie
- Hiba Hammi

## Sujet 
- Description du sujet : Implémentation d'une partie du jeu Galérapagos avec 3 à 12 joueurs virtuels, l'utilisateur étant considéré comme le maître du jeu.
- Enjeux de la simulation : Observation du déroulement de la partie jusqu'à obtenir des joueurs gagnants et/ou perdants.

## Architecture logicielle 
- Front-end en ReactJS
- Back-end en Golang

## Agents

### Plateau de jeu
Décrit par :
- Nombre de tour de la partie
- Tour actuel
- Météo
- Stock d'eau 
- Stock de nourriture
- Stock de bois
- Place disponible dans le radeau

### Joueur
Décrit par :
- un identifiant
- un nom
- des traits de caractères : intelligence et/ou égoïsme
- des compétences : pêcheur ou bucheron
- un booléen à true si le joueur est encore vivant
- une liste de préférence des joueurs 

### Déroulement d’un tour de jeu :
- Changement du premier joueur
- Tirage de la carte Météo
- Action des joueurs (pêcher, ramener de l'eau ou récupérer du bois)
- Survie des naufragés


