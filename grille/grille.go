package grille

import (
	"BatailleNavale/bateau"
	"fmt"
	"math/rand"
	"time"
)

// Représente la grille de jeu
type Grille struct {
	grille    [10][10]Case //grille de jeu
	nbBateaux int          //nombre de bateaux
}

// Représente une case de la grille
type Case struct {
	estBateau bool          //true si la case contient un bateau
	bateau    bateau.Bateau //le bateau contenu dans la case
	estTouche bool          //true si la case a été touchée
}

// Initialise la grille
func (g *Grille) InitGrille() {

	//initialitsation de toutes les cases à false
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			g.grille[i][j].estBateau = false
			g.grille[i][j].estTouche = false
		}
	}

	//initialisation du générateur de nombres aléatoires
	rand.Seed(time.Now().UnixNano())

	//on initialise le nombre de bateaux
	g.nbBateaux = 5

	//on initialise les bateaux
	for i := 0; i < g.nbBateaux; i++ {
		var b bateau.Bateau
		b.InitBateau()

		//check si le bateau chevauche un autre bateau
		for g.Chevauche(b.XDebut, b.YDebut, b.XFin, b.YFin) {
			b.InitBateau()
		}

		g.placerBateau(b)
	}
}

// Place un bateau sur la grille
func (g *Grille) placerBateau(b bateau.Bateau) {

	//on place le bateau sur la grille
	if b.XDebut == b.XFin { //le bateau est vertical
		for i := b.YDebut; i <= b.YFin; i++ {
			g.grille[b.XDebut][i].estBateau = true
			g.grille[b.XDebut][i].bateau = b
		}
	} else { //le bateau est horizontal
		for i := b.XDebut; i <= b.XFin; i++ {
			g.grille[i][b.YDebut].estBateau = true
			g.grille[i][b.YDebut].bateau = b
		}
	}
}

// Vérifie si le bateau chevauche un autre bateau
func (g *Grille) Chevauche(XDebut int, YDebut int, XFin int, YFin int) bool {
	//on vérifie que les coordonnées sont valides
	if XDebut < 0 || XDebut > 9 || YDebut < 0 || YDebut > 9 || XFin < 0 || XFin > 9 || YFin < 0 || YFin > 9 {
		return true
	}

	if XDebut > XFin || YDebut > YFin {
		return true
	}

	//on vérifie que le bateau ne chevauche pas un autre bateau
	for i := XDebut; i <= XFin; i++ {
		for j := YDebut; j <= YFin; j++ {
			if g.grille[i][j].estBateau {
				return true
			}
		}
	}

	//on retourne false si le bateau ne chevauche pas un autre bateau

	return false
}

// Affiche la grille
func (g *Grille) AfficherGrille() {
	//afficher les cases
	// _ : case vide
	// X : case touchée
	// T : bateau touché
	// C : bateau coulé
	// | : séparateur de colonne

	//afficher les numéros de colonne
	fmt.Print("    ")
	for i := 0; i < 10; i++ {
		fmt.Print(i, "   ")
	}
	fmt.Println()

	//afficher les cases
	for i := 0; i < 10; i++ {
		//afficher le numéro de ligne
		fmt.Print(i, " | ")

		//afficher les cases
		for j := 0; j < 10; j++ {
			if !g.grille[i][j].estTouche {
				if g.grille[i][j].estBateau {
					if g.EstCoule(i, j) {
						fmt.Print("C | ")
					} else {
						fmt.Print("T | ")
					}
				} else {
					fmt.Print("X | ")
				}
			} else {
				fmt.Print("_ | ")
			}
		}

		//afficher le numéro de ligne
		fmt.Print(i)

		//passer à la ligne
		fmt.Println()
	}

	//afficher les numéros de colonne
	fmt.Print("    ")
	for i := 0; i < 10; i++ {
		fmt.Print(i, "   ")
	}
	fmt.Println()

	fmt.Print("Nombre de bateaux restants : ", g.NbBateauxRestants())
	fmt.Println()
	fmt.Println()
}

// Tire sur une case
func (g *Grille) Tirer(x int, y int) bool {
	//on vérifie que la case n'a pas déjà été touchée
	if g.grille[x][y].estTouche {
		return false
	}

	//on marque la case comme touchée
	g.grille[x][y].estTouche = true

	//on retourne true si la case contient un bateau
	return g.grille[x][y].estBateau
}

// On verifie si le bateau est coule
func (g *Grille) EstCoule(x int, y int) bool {
	//on vérifie que la case contient un bateau
	if !g.grille[x][y].estBateau {
		return false
	}

	//on récupère le bateau
	b := g.grille[x][y].bateau

	//on vérifie que toutes les cases du bateau sont touchées
	for i := b.XDebut; i <= b.XFin; i++ {
		for j := b.XDebut; j <= b.YFin; j++ {
			if !g.grille[i][j].estTouche {
				return false
			}
		}
	}

	//on retourne true si toutes les cases du bateau sont touchées
	return true
}

// Compte le nombre de bateaux restants
func (g *Grille) NbBateauxRestants() int {

	nbBateauxRestants := 0

	//on compte le nombre de bateaux restants
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if g.grille[i][j].estBateau && !g.EstCoule(i, j) {
				nbBateauxRestants++
			}
		}
	}

	//on retourne le nombre de bateaux restants
	return nbBateauxRestants

}

// Vérifie si la partie est terminée
func (g *Grille) PartieTerminee() bool {
	return g.NbBateauxRestants() == 0
}
