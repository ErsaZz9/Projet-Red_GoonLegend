package character

import "fmt"

type Character struct {
	Nom        string
	Classe     string
	Niveau     int
	PVMax      int
	PV         int
	Inventaire []string
}

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

func (c Character) DisplayInfo() {
	rouge := "\033[31m"
	reset := "\033[0m"

	fmt.Println(rouge + "=== Informations du personnage ===" + reset)
	fmt.Println("Nom       :", c.Nom)
	fmt.Println("Classe    :", c.Classe)
	fmt.Println("Niveau    :", c.Niveau)
	fmt.Printf("PV        : %d/%d\n", c.PV, c.PVMax)
	fmt.Println("Inventaire:", c.Inventaire)
	fmt.Println(rouge + "=================================" + reset)
}

func (c Character) AccessInventory() {
	rouge := "\033[31m"
	reset := "\033[0m"

	fmt.Println(rouge + "=== Inventaire ===" + reset)
	if len(c.Inventaire) == 0 {
		fmt.Println("L'inventaire est vide.")
		return
	}
	for i, item := range c.Inventaire {
		fmt.Printf("%d. %s\n", i+1, item)
	}
	fmt.Println(rouge + "==================" + reset)
}
