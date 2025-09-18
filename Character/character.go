package character

import (
	"Projet-Red_GoonLegend/classes"
	"Projet-Red_GoonLegend/items"
	"fmt"
)

// struct pour les compétences
type Skill struct {
	Nom       string
	Degat     int
	CountMana int
}

// struct pour le personnage
type Character struct {
	Nom        string
	Classe     classes.Classe
	Niveau     int
	PVMax      int
	PV         int
	Inventaire []items.Item // ⚡ slice d’objets (définis dans items/)
	Skills     []Skill
}

// character init (perso de base)
func InitCharacter(nom string, classe classes.Classe, inventaire []items.Item) Character {
	return Character{
		Nom:        nom,
		Classe:     classe,
		Niveau:     1,
		PVMax:      classe.PVMax,
		PV:         classe.PVMax / 2, // 50% des PV
		Inventaire: inventaire,
		Skills:     []Skill{{Nom: "Coup de poing", Degat: 10, CountMana: 0}},
	}
}

func (c *Character) SpellBook(sortName string) {
	for _, skill := range c.Skills {
		if skill.Nom == sortName {
			fmt.Printf("Vous connaissez déjà %s.\n", sortName)
			return
		}
	}

	// ici tu définis chaque sort possible
	switch sortName {
	case "Boule de Feu":
		c.Skills = append(c.Skills, Skill{Nom: "Boule de Feu", Degat: 30, CountMana: 5})
		fmt.Println("Vous avez appris Boule de Feu !")
	case "Éclair":
		c.Skills = append(c.Skills, Skill{Nom: "Éclair", Degat: 40, CountMana: 8})
		fmt.Println("Vous avez appris Éclair !")
	default:
		fmt.Printf("Le sort %s est inconnu.\n", sortName)
	}
}

// affiche les infos du perso
func (c Character) DisplayInfo() {
	fmt.Println("=== Informations du personnage ===")
	fmt.Println("Nom       :", c.Nom)
	fmt.Println("Classe    :", c.Classe.Nom)
	fmt.Println("Niveau    :", c.Niveau)
	fmt.Printf("PV        : %d/%d\n", c.PV, c.PVMax)
	fmt.Println("Inventaire:")
	for _, item := range c.Inventaire {
		fmt.Printf("- %s\n", item.Nom)
	}
	fmt.Println("Skills:")
	for _, skill := range c.Skills {
		fmt.Printf("- %s\n", skill.Nom)
	}
	fmt.Println("=================================")
}

// affiche l’inventaire
func (c Character) AccessInventory() {
	fmt.Println("=== Inventaire ===")
	if len(c.Inventaire) == 0 {
		fmt.Println("L'inventaire est vide.")
		return
	}
	for i, item := range c.Inventaire {
		fmt.Printf("%d. %s\n", i+1, item.Nom)
	}
	fmt.Println("==================")
}

// ajoute un objet à l’inventaire
func (c *Character) AddInventory(item items.Item) {
	if c.IsInventoryFull() {
		fmt.Println("Vous ne pouvez pas stocker plus de 10 objets dans votre inventaire !")
		return
	}

	c.Inventaire = append(c.Inventaire, item)
	fmt.Printf("%s ajouté à l’inventaire.\n", item.Nom)
}

// retire un objet de l’inventaire
func (c *Character) RemoveInventory(nom string) {
	for i, invItem := range c.Inventaire {
		if invItem.Nom == nom {
			c.Inventaire = append(c.Inventaire[:i], c.Inventaire[i+1:]...)
			return
		}
	}
}

// consommer une potion de soin
func (c *Character) UsePotion() {
	index := -1
	for i, item := range c.Inventaire {
		if item.Nom == "Potion de soin" && item.Soin > 0 { // ⚡ CHANGÉ
			index = i
			// applique le soin
			c.PV += item.Soin
			if c.PV > c.PVMax {
				c.PV = c.PVMax
			}
			// retire la potion consommée
			c.Inventaire = append(c.Inventaire[:i], c.Inventaire[i+1:]...)
			fmt.Printf("%s utilisée ! PV actuels : %d/%d\n", item.Nom, c.PV, c.PVMax)
			break
		}
	}
	if index == -1 {
		fmt.Println("Aucune potion disponible !")
	}
}

// appliquer un poison
func (c *Character) ApplyPoison(dmg int, duration int) {
	for i := 0; i < duration; i++ {
		c.PV -= dmg
		if c.PV < 0 {
			c.PV = 0
		}
		fmt.Printf("Le poison agit ! PV actuels : %d/%d\n", c.PV, c.PVMax)
		// ⚡ ici tu mettras time.Sleep plus tard si tu veux le délai
	}
}

// vérifier la mort
func (c *Character) IsDead() bool {
	if c.PV <= 0 {
		fmt.Printf("%s est mort...\n", c.Nom)
		// respawn 50%
		resurrectHP := c.PVMax / 2
		if resurrectHP <= 0 {
			resurrectHP = 1
		}
		c.PV = resurrectHP
		fmt.Printf("%s a été ressuscité avec %d PV.\n", c.Nom, c.PV)
		return true
	}
	return false
}

func (c *Character) IsInventoryFull() bool {
	return len(c.Inventaire) >= 10
}
