# Galerapagos_IA04

## Membres du groupe
- Julie Szerbojm
- Romane Guari      
- Sihan Xie
- Hiba Hammi

## Sujet 
- Description du sujet : Implémentation d'une partie du jeu Galérapagos avec 3 à 12 joueurs virtuels, l'utilisateur étant considéré comme le maître du jeu.
- Enjeux de la simulation : Observation du déroulement de la partie (gagnants et perdants).

## Typologie des agents
### Caractère (taux de 0 à 10)
#### Egoïsme
- 10 : le joueur est très égoïste, la priorité du joueur est de s'en sortir seul;
- 0 : le joueur est altruiste, il veut sauver tout le monde quitte à se sacrifier pour les autres;
- 5 : le joueur mixte, sa survie est importante mais il préfèrerait sauver tout le monde.

Réflexion d'autres traits de caractère :
#### Intelligence 
- 10 : le joueur joue de façon intelligente par rapport à sa situation
- 0 : le joueur joue "mal"

### Capacités 
- Pêcheur : ramène plus de poissons que la moyenne;
- Bricoleur : construit plus rapidement le radeau.

## Architecture logicielle 
- Front-end en React
- Back-end en Golang
