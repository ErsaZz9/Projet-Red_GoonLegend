package character

import "fmt"

// struct pour les champs (g)

type Character struct {
	Nom        string
	Classe     string
	Niveau     int
	PVMax      int
	PV         int
	Inventaire []string
}

// character init on initialise
func InitCharacter(nom string, classe string, niveau int, pvMax int, pv int, inventaire []string) Character {
	return Character{
		Nom:        nom,
		Classe:     classe,
		Niveau:     niveau,
		PVMax:      pvMax,
		PV:         pv,
		Inventaire: inventaire,
	}
}

// affiche les infos du personnage (l)
func (c Character) DisplayInfo() {
	fmt.Println("=== Informations du personnage ===")
	fmt.Println("Nom       :", c.Nom)
	fmt.Println("Classe    :", c.Classe)
	fmt.Println("Niveau    :", c.Niveau)
	fmt.Printf("PV        : %d/%d\n", c.PV, c.PVMax)
	fmt.Println("Inventaire:", c.Inventaire)
	fmt.Println("=================================")
}

// affiche l’inventaire (a)
func (c Character) AccessInventory() {
	fmt.Println("=== Inventaire ===")
	if len(c.Inventaire) == 0 {
		fmt.Println("L'inventaire est vide.")
		return
	}
	for i, item := range c.Inventaire {
		fmt.Printf("%d. %s\n", i+1, item)
	}
	fmt.Println("==================")
}

// on crée une une fonction va permettre de simplifier le système d'ajout à l'inventaire, évitant de faire des append partout [astuce] (n)
func (c *Character) AddInventory(item string) {
	c.Inventaire = append(c.Inventaire, item)
}

// et vice verqa
func (c *Character) RemoveInventory(item string) {
	for i, invItem := range c.Inventaire {
		if invItem == item {
			// comme dhab, ici on prend tout avant i, (la potion la ou j'en suis), et tout après i. et return 🥶 [astuce] (d)
			c.Inventaire = append(c.Inventaire[:i], c.Inventaire[i+1:]...)
			return
		}
	}
}
