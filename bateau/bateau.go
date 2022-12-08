package bateau

import (
	"math/rand"
	"time"
)

// class Bateau : représente un bateau
type Bateau struct {
	//position du bateau
	Id     int
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

	//On choisit un id aléatoire unique pour chaque bateau
	b.Id = rand.Intn(1000)

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

	//On vérifie que le bateau est bien dans la grille et ne passe pas par les bords de la grille
	if b.XDebut < 0 || b.XDebut > 9 || b.YDebut < 0 || b.YDebut > 9 || b.XFin < 0 || b.XFin > 9 || b.YFin < 0 || b.YFin > 9 {
		b.InitBateau()
	}

}
