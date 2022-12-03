package grille

import (
	"fmt"
	"math/rand"
	"time"
)

//class Grille : représente une grille de jeu de bataille navale
type Grille struct {
	//tableau de 10x10 cases
	cases [10][10]Case
}

//class Case : représente une case de la grille
type Case struct {
	//true si la case contient un bateau
	estBateau bool
	//true si la case a été touchée
	estTouchee bool
	//true si le bateau a été coulé
	estCoule bool
}

//constructeur de la classe Grille
func NewGrille() Grille {
	var g Grille
	return g
}

//initialise la grille avec des bateaux
func (g *Grille) InitGrille() {
	//initialisation du générateur de nombres aléatoires
	rand.Seed(time.Now().UnixNano())

	//on place les bateaux
	for i := 0; i < 5; i++ {
		g.placerBateau()
	}
}

//place un bateau sur la grille

func (g *Grille) placerBateau() {
	//on choisit une case au hasard
	x := rand.Intn(10)
	y := rand.Intn(10)

	//on vérifie que la case est vide
	if g.cases[x][y].estBateau {
		//si la case n'est pas vide, on recommence
		g.placerBateau()
	} else {
		//sinon on place le bateau
		g.cases[x][y].estBateau = true
	}
}

//affiche la grille
func (g *Grille) AfficherGrille() {
	//afficher les cases
	// _ : case vide
	// X : case touchée
	// T : bateau touché
	// C : bateau coulé
	// | : séparateur de colonne

	//afficher les numéros de colonne
	fmt.Print("   ")
	for i := 0; i < 10; i++ {
		fmt.Print(i, " ")
	}
	fmt.Println()
	
	for i := 0; i < 10; i++ {
		fmt.Print(i, " |")
		for j := 0; j < 10; j++ {
			if g.cases[i][j].estTouchee {
				if g.cases[i][j].estBateau {
					if g.cases[i][j].estCoule {
						fmt.Print("C")
					} else {
						fmt.Print("T")
					}
				} else {
					fmt.Print("X")
				}
			} else {
				fmt.Print("_")
			}
			fmt.Print("|")
		}
		fmt.Println("",i)
	}

	//afficher les numéros de colonne
	fmt.Print("   ")
	for i := 0; i < 10; i++ {
		fmt.Print(i, " ")
	}
	
}

//tire sur une case de la grille
func (g *Grille) Tirer(x int, y int) bool {
	//on vérifie que la case n'a pas déjà été touchée
	if g.cases[x][y].estTouchee {
		fmt.Println("Vous avez déjà tiré sur cette case")
		return false
	} else {
		//sinon on marque la case comme touchée
		g.cases[x][y].estTouchee = true
		return g.cases[x][y].estBateau
	}

}

func (g *Grille) EstCoule(x int, y int) bool {
	if g.cases[x][y].estBateau {
		g.cases[x][y].estCoule = true
		return true
	} else {
		return false
	}
}

//vérifie si la partie est terminée
func (g *Grille) EstTerminee() bool {
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if g.cases[i][j].estBateau && !g.cases[i][j].estTouchee {
				return false
			}
		}
	}
	return true
}

//compte le nombre de bateaux restants
func (g *Grille) CompterBateaux() int {
	nb := 0
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if g.cases[i][j].estBateau && !g.cases[i][j].estTouchee {
				nb++
			}
		}
	}
	return nb
}
