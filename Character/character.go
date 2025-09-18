package character

import (
	"Projet-Red_GoonLegend/classes"
	"Projet-Red_GoonLegend/items"
	"fmt"
)

// struct pour les competences
type Skill struct {
	Nom       string
	Degat     int
	CountMana int
}

// struct pour les equipements
type Equipment struct {
	Head  *items.Item
	Chest *items.Item
	Feet  *items.Item
}

// struct pour le personnage
type Character struct {
	Nom                string
	Classe             classes.Classe
	Niveau             int
	PVMax              int
	PV                 int
	Purse              int
	Inventaire         []items.Item
	Skills             []Skill
	Equipement         Equipment
	MaxInventory       int
	InventoryUpgrades  int
	Initiative         int
	Experience         int
	ExperienceMax      int
	Mana               int
	ManaMax            int
}

// character init (perso de base)
func InitCharacter(nom string, classe classes.Classe, inventaire []items.Item) Character {
	return Character{
		Nom:                nom,
		Classe:             classe,
		Niveau:             1,
		PVMax:              classe.PVMax,
		PV:                 classe.PVMax / 2,
		Purse:              100,
		Inventaire:         inventaire,
		Skills:             []Skill{{Nom: "Coup de poing", Degat: 10, CountMana: 0}},
		Equipement:         Equipment{},
		MaxInventory:       10, // capacite initiale
		InventoryUpgrades:  0,  // compteur d'upgrades
		Initiative:         classe.Dexterite,
		Experience:         0,
		ExperienceMax:      100,
		Mana:               classe.ManaMax,
		ManaMax:            classe.ManaMax,
	}
}

func (c *Character) SpellBook(sortName string) {
	for _, skill := range c.Skills {
		if skill.Nom == sortName {
			fmt.Printf("Vous connaissez deja %s.\n", sortName)
			return
		}
	}

	switch sortName {
	case "Boule de Feu":
		c.Skills = append(c.Skills, Skill{Nom: "Boule de Feu", Degat: 30, CountMana: 5})
		fmt.Println("Vous avez appris Boule de Feu !")
	case "Eclair":
		c.Skills = append(c.Skills, Skill{Nom: "Eclair", Degat: 40, CountMana: 8})
		fmt.Println("Vous avez appris Eclair !")
	case "Soin":
		c.Skills = append(c.Skills, Skill{Nom: "Soin", Degat: -30, CountMana: 3})
		fmt.Println("Vous avez appris Soin !")
	case "Bouclier":
		c.Skills = append(c.Skills, Skill{Nom: "Bouclier", Degat: 0, CountMana: 2})
		fmt.Println("Vous avez appris Bouclier !")
	case "Meteore":
		c.Skills = append(c.Skills, Skill{Nom: "Meteore", Degat: 60, CountMana: 12})
		fmt.Println("Vous avez appris Meteore !")
	case "Glace":
		c.Skills = append(c.Skills, Skill{Nom: "Glace", Degat: 25, CountMana: 4})
		fmt.Println("Vous avez appris Glace !")
	case "Vent":
		c.Skills = append(c.Skills, Skill{Nom: "Vent", Degat: 20, CountMana: 3})
		fmt.Println("Vous avez appris Vent !")
	case "Terre":
		c.Skills = append(c.Skills, Skill{Nom: "Terre", Degat: 35, CountMana: 6})
		fmt.Println("Vous avez appris Terre !")
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
	fmt.Printf("Mana      : %d/%d\n", c.Mana, c.ManaMax)
	fmt.Printf("Experience: %d/%d\n", c.Experience, c.ExperienceMax)
	fmt.Printf("Initiative: %d\n", c.Initiative)
	fmt.Printf("Purse     : %d\n", c.Purse)
	fmt.Printf("Capacite  : %d slots\n", c.MaxInventory)
	fmt.Println("Inventaire:")
	for _, item := range c.Inventaire {
		fmt.Printf("- %s\n", item.Nom)
	}
	fmt.Println("Skills:")
	for _, skill := range c.Skills {
		fmt.Printf("- %s (Cout: %d mana)\n", skill.Nom, skill.CountMana)
	}
	fmt.Println("Equipements equipes:")
	if c.Equipement.Head != nil {
		fmt.Printf("Tete  : %s\n", c.Equipement.Head.Nom)
	}
	if c.Equipement.Chest != nil {
		fmt.Printf("Torse : %s\n", c.Equipement.Chest.Nom)
	}
	if c.Equipement.Feet != nil {
		fmt.Printf("Pieds : %s\n", c.Equipement.Feet.Nom)
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

// ajoute un objet a l’inventaire
func (c *Character) AddInventory(item items.Item) {
	if c.IsInventoryFull() {
		fmt.Printf("Vous ne pouvez pas stocker plus de %d objets dans votre inventaire !\n", c.MaxInventory)
		return
	}
	c.Inventaire = append(c.Inventaire, item)
	fmt.Printf("%s ajoute a l'inventaire.\n", item.Nom)
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
		if item.Nom == "Potion de soin" && item.Soin > 0 {
			index = i
			c.PV += item.Soin
			if c.PV > c.PVMax {
				c.PV = c.PVMax
			}
			c.Inventaire = append(c.Inventaire[:i], c.Inventaire[i+1:]...)
			fmt.Printf("%s utilisee ! PV actuels : %d/%d\n", item.Nom, c.PV, c.PVMax)
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
	}
}

// verifier la mort
func (c *Character) IsDead() bool {
	if c.PV <= 0 {
		fmt.Printf("%s est mort...\n", c.Nom)
		resurrectHP := c.PVMax / 2
		if resurrectHP <= 0 {
			resurrectHP = 1
		}
		c.PV = resurrectHP
		fmt.Printf("%s a ete ressuscite avec %d PV.\n", c.Nom, c.PV)
		return true
	}
	return false
}

func (c *Character) IsInventoryFull() bool {
	return len(c.Inventaire) >= c.MaxInventory
}

// --- Gestion de l'equipement ---

func (c *Character) Equip(item items.Item) {
	switch item.Nom {
	case "Chapeau de l'aventurier":
		c.replaceEquip(&c.Equipement.Head, item, 10)
	case "Tunique de l'aventurier":
		c.replaceEquip(&c.Equipement.Chest, item, 25)
	case "Bottes de l'aventurier":
		c.replaceEquip(&c.Equipement.Feet, item, 15)
	default:
		fmt.Println("Cet objet ne peut pas etre equipe.")
	}
}

func (c *Character) replaceEquip(slot **items.Item, newItem items.Item, bonus int) {
	if *slot != nil {
		c.Inventaire = append(c.Inventaire, **slot)
		c.PVMax -= bonus
		if c.PV > c.PVMax {
			c.PV = c.PVMax
		}
	}
	*slot = &newItem
	c.PVMax += bonus
	fmt.Printf("%s equipe ! PV max : %d\n", newItem.Nom, c.PVMax)
	c.RemoveInventory(newItem.Nom)
}

// --- Upgrade d'inventaire ---

func (c *Character) UpgradeInventorySlot() {
	if c.InventoryUpgrades >= 3 {
		fmt.Println("Vous ne pouvez plus augmenter votre inventaire (maximum atteint).")
		return
	}
	c.MaxInventory += 10
	c.InventoryUpgrades++
	fmt.Printf("Inventaire augmente ! Capacite maximale : %d\n", c.MaxInventory)
}

// Ajouter de l'expérience
func (c *Character) AddExperience(exp int) {
	c.Experience += exp
	fmt.Printf("Vous gagnez %d points d'experience !\n", exp)
	
	for c.Experience >= c.ExperienceMax {
		c.Experience -= c.ExperienceMax
		c.Niveau++
		c.ExperienceMax = int(float64(c.ExperienceMax) * 1.5) // Augmentation progressive
		
		// Bonus de niveau
		c.PVMax += 10
		c.PV = c.PVMax // Restauration complète
		c.ManaMax += 5
		c.Mana = c.ManaMax // Restauration complète
		
		fmt.Printf("Niveau %d atteint ! PV max: %d, Mana max: %d\n", c.Niveau, c.PVMax, c.ManaMax)
	}
}

// Utiliser du mana
func (c *Character) UseMana(cost int) bool {
	if c.Mana >= cost {
		c.Mana -= cost
		return true
	}
	return false
}

// Restaurer du mana
func (c *Character) RestoreMana(amount int) {
	c.Mana += amount
	if c.Mana > c.ManaMax {
		c.Mana = c.ManaMax
	}
}
