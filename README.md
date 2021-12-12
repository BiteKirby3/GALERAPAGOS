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

## Description des agents
Joueur :
- ID int
- note int

Carte :
- type (eau, nourriture, antidote)
- quantité par défaut à 1

Meteo :
- nom String
- quantité integer (entre 0 et 3) il s'agit de la quantité d'eau qui peut être récupérée pendant le tour


### Agent = joueur
#### Attributs
- ID int 
- nom String
- listeJoueurs []Joueur
- estMalade bool
- estMort bool
- listeCartes []Carte
- egoisme int (varie entre 0 et 10)
- intelligence int (varie entre 0 et 10)
- pecheur boolean
- bricoleur boolean
#### Méthodes
- MAJNotesJoueur
    - mettre à jour les notes des joueurs en fonction de leurs actions
- DemandeRessources
	- cette méthode va permettre de demander à d'autres joueurs s'il peuvent et veulent partager un carte ressource (eau,nourriture,anti-venin)
- Jouer
	- 4 actions possibles = pêcher, récupérer du bois, chercher de l'eau, chercher des ressources dans l'épave
	- cette méthode appelera l'une de des deux fonctions suivantes (l'appel d'une de ces fonctions dépendra du nombre attribué en intelligence) :
		- JoueurIntelligent
			- une personne égoiste (egoisme > 5) aura plus tendance à chercher une nouvelle carte/ressource dans l'épave alors qu'une personne altruiste (egoisme < 5) aura plus tendance à chercher de l'eau, de la nourriture ou du bois pour le groupe
		- JoueurAléatoirement
			- comme son nom l'indique le joueur réalise une action au hasard
- NoteMax 
	- le joueur vote pour l'adversaire avec la note la plus élevée


### Agent = Epave
- quand un joueur l'interroge il retourne une carte epave


### Agent = Bois
- quand un joueur l'interroge il retourne une quantité de bois (entre 1 et 5) ou si le joueur tombe malade(dans ce cas il ne reçoit pas de bois) 
	

### Agent = Méteo
- listeMeteo []Meteo
- permet de connaitre la quantité d'eau par tour
- permet de savoir si c'est le tour de l'ouragan

### Agent = Radeau
- permet de savoir combien de place son disponible sur le radeau

	
### Agent = StockEau
- permet de stocker l'eau du groupe

### Agent = StockNourriture
- permet de stocker la nourriture commune

### Agent = DebutJeu
- permet de lancer les agents
- demande le nombre de joueur à l'utilisateur et le premier joueur
- initialise les notes des adversaires pour chaque joueur 
(retirer 1 pour pecheur et bricoleur // le joueur précedent et suivant reçoivent respectivement un +2 et +1)

Remarque : Le joueur choisit d'éliminer l'adversaire avec la note la plus élevée


### Agent = JeuManager
- compte le nombre de tour
- indique qui est le premier joueur
- indique qui doit jouer
- demande à l'agent RessourcesManager s'il faut déclencher la phase de vote

### Agent = Vote
- récupère les votes
  - si égalité on demande au premier joueur de sélectionner un joueur
  - si le joueur choisi utilise ses cartes ressources il peut survivre
  - sinon le joueur décède 


### Agent = RessourcesManager
- Calcul la quantité d'eau et de nourriture à la fin du tour
- Demande au joueurs s'ils veulent partager leurs ressources
	
### Agent = Utilisateur
- Il s'agit d'un agent qui peut décide du nombre de joueur au départ et du joueur de départ. 
- Il peut également modifier certains paramètre au cours du jeu :
	- décider de l'arrivée de l'ouragan
	- choisir la quantité d'eau qui peut-être récupérée au prochain tour


### Récapitulatif d’un tour de jeu :
- Changement du premier joueur
- Tirage de la carte Météo
- Action des joueurs
- Survie des naufragés
- Fin du tour

## Architecture logicielle 
- Front-end en React
- Back-end en Golang
