package bateau

import (
	"math/rand"
	"time"
)

// class Bateau : représente un bateau
type Bateau struct {
	//position du bateau
	XDebut int
	YDebut int
	XFin   int
	YFin   int
	Taille int
}

// Initialise le bateau avec une position aléatoire et une Taille aléatoire (entre 2 et 5) et retourne le bateau initialisé qui prend en parametre un bateau
func (b *Bateau) InitBateau() {

	//initialisation du générateur de nombres aléatoires
	rand.Seed(time.Now().UnixNano())

	//on choisit une position aléatoire
	b.XDebut = rand.Intn(10)
	b.YDebut = rand.Intn(10)

	//on choisit une direction aléatoire
	direction := rand.Intn(4)

	//on choisit une Taille aléatoire
	Taille := rand.Intn(4) + 2
	b.Taille = Taille

	//on initialise la position de fin du bateau
	switch direction {
	case 0: //haut
		b.XFin = b.XDebut
		b.YFin = b.YDebut - Taille
	case 1: //droite
		b.XFin = b.XDebut + Taille
		b.YFin = b.YDebut
	case 2: //bas
		b.XFin = b.XDebut
		b.YFin = b.YDebut + Taille
	case 3: //gauche
		b.XFin = b.XDebut - Taille
		b.YFin = b.YDebut
	}
}
