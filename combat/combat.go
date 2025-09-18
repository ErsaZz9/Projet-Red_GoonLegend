package combat

import (
	"Projet-Red_GoonLegend/character"
	"Projet-Red_GoonLegend/utils"
	"fmt"

	"github.com/manifoldco/promptui"
)

// Structure pour les monstres
type Monster struct {
	Nom       string
	PVMax     int
	PV        int
	Attaque   int
	Initiative int
}

// Initialise un gobelin d'entraînement
func InitGoblin() Monster {
	return Monster{
		Nom:        "Gobelin d'entrainement",
		PVMax:      40,
		PV:         40,
		Attaque:    5,
		Initiative: 3,
	}
}

// Pattern d'attaque du gobelin
func (m *Monster) GoblinPattern(hero *character.Character, tour int) {
	degats := m.Attaque
	if tour%3 == 0 {
		degats = m.Attaque * 2
		fmt.Printf("%s utilise une attaque puissante !\n", m.Nom)
	}
	
	hero.PV -= degats
	if hero.PV < 0 {
		hero.PV = 0
	}
	
	fmt.Printf("%s inflige %d degats a %s\n", m.Nom, degats, hero.Nom)
	fmt.Printf("PV de %s : %d/%d\n", hero.Nom, hero.PV, hero.PVMax)
}

// Tour du joueur
func CharacterTurn(hero *character.Character, monster *Monster) {
	for {
		utils.ClearScreen()
		fmt.Printf("=== Tour de %s ===\n", hero.Nom)
		fmt.Printf("PV : %d/%d | Mana : %d/%d\n", hero.PV, hero.PVMax, hero.Mana, hero.ManaMax)
		fmt.Printf("Ennemi : %s (%d/%d PV)\n", monster.Nom, monster.PV, monster.PVMax)
		fmt.Println()

		options := []string{
			"Attaquer",
			"Sorts",
			"Inventaire",
		}

		prompt := promptui.Select{
			Label: "Que voulez-vous faire ?",
			Items: options,
		}

		index, _, err := prompt.Run()
		if err != nil {
			return
		}

		switch index {
		case 0:
			// Attaque basique
			degats := 5
			monster.PV -= degats
			if monster.PV < 0 {
				monster.PV = 0
			}
			fmt.Printf("%s inflige %d degats a %s\n", hero.Nom, degats, monster.Nom)
			fmt.Printf("PV de %s : %d/%d\n", monster.Nom, monster.PV, monster.PVMax)
			fmt.Println("\n(Appuyez sur Entree pour continuer)")
			fmt.Scanln()
			return
		case 1:
			// Utiliser les sorts
			useSpellsInCombat(hero, monster)
			return
		case 2:
			// Utiliser l'inventaire
			useInventoryInCombat(hero, monster)
			return
		}
	}
}

// Utiliser l'inventaire en combat
func useInventoryInCombat(hero *character.Character, monster *Monster) {
	for {
		utils.ClearScreen()
		fmt.Printf("=== Inventaire en combat ===\n")
		fmt.Printf("PV : %d/%d\n", hero.PV, hero.PVMax)
		fmt.Printf("Ennemi : %s (%d/%d PV)\n", monster.Nom, monster.PV, monster.PVMax)
		fmt.Println()

		options := make([]string, 0, len(hero.Inventaire)+1)
		for _, item := range hero.Inventaire {
			options = append(options, item.Nom)
		}
		options = append(options, "← Retour")

		prompt := promptui.Select{
			Label: "Choisissez un objet :",
			Items: options,
		}

		index, _, err := prompt.Run()
		if err != nil {
			return
		}

		if index == len(options)-1 {
			return
		}

		item := hero.Inventaire[index]
		
		switch item.Nom {
		case "Potion de soin":
			hero.UsePotion()
			fmt.Println("Vous utilisez une Potion de soin !")
		case "Potion de mana":
			hero.RestoreMana(item.Soin)
			hero.RemoveInventory(item.Nom)
			fmt.Printf("Vous utilisez une Potion de mana ! Mana : %d/%d\n", hero.Mana, hero.ManaMax)
		case "Potion de poison":
			hero.ApplyPoison(item.Degats, 3)
			hero.RemoveInventory(item.Nom)
			fmt.Println("Vous utilisez une Potion de poison !")
		default:
			fmt.Printf("Impossible d'utiliser %s en combat.\n", item.Nom)
		}

		fmt.Println("\n(Appuyez sur Entree pour continuer)")
		fmt.Scanln()
		return
	}
}

// Utiliser les sorts en combat
func useSpellsInCombat(hero *character.Character, monster *Monster) {
	for {
		utils.ClearScreen()
		fmt.Printf("=== Sorts disponibles ===\n")
		fmt.Printf("Mana : %d/%d\n", hero.Mana, hero.ManaMax)
		fmt.Printf("Ennemi : %s (%d/%d PV)\n", monster.Nom, monster.PV, monster.PVMax)
		fmt.Println()

		options := make([]string, 0, len(hero.Skills)+1)
		for _, skill := range hero.Skills {
			options = append(options, fmt.Sprintf("%s (Cout: %d mana)", skill.Nom, skill.CountMana))
		}
		options = append(options, "← Retour")

		prompt := promptui.Select{
			Label: "Choisissez un sort :",
			Items: options,
		}

		index, _, err := prompt.Run()
		if err != nil {
			return
		}

		if index == len(options)-1 {
			return
		}

		skill := hero.Skills[index]
		
		// Vérifier le mana
		if !hero.UseMana(skill.CountMana) {
			fmt.Printf("Mana insuffisant ! Il vous faut %d mana pour utiliser %s.\n", skill.CountMana, skill.Nom)
			fmt.Println("\n(Appuyez sur Entree pour continuer)")
			fmt.Scanln()
			continue
		}

		// Utiliser le sort
		degats := skill.Degat
		
		if skill.Nom == "Soin" {
			// Sort de soin
			hero.PV -= degats // degats négatifs = soin
			if hero.PV > hero.PVMax {
				hero.PV = hero.PVMax
			}
			fmt.Printf("%s utilise %s et se soigne de %d PV !\n", hero.Nom, skill.Nom, -degats)
			fmt.Printf("PV de %s : %d/%d\n", hero.Nom, hero.PV, hero.PVMax)
		} else if skill.Nom == "Bouclier" {
			// Sort de bouclier (protection temporaire)
			fmt.Printf("%s utilise %s et se protège !\n", hero.Nom, skill.Nom)
			fmt.Println("Vous êtes protégé pour ce tour !")
		} else if skill.Nom == "Meteore" {
			// Sort de météore (dégâts élevés)
			monster.PV -= degats
			if monster.PV < 0 {
				monster.PV = 0
			}
			fmt.Printf("%s invoque une météore et inflige %d dégâts à %s !\n", hero.Nom, degats, monster.Nom)
			fmt.Printf("PV de %s : %d/%d\n", monster.Nom, monster.PV, monster.PVMax)
		} else if skill.Nom == "Glace" {
			// Sort de glace (ralentissement)
			monster.PV -= degats
			if monster.PV < 0 {
				monster.PV = 0
			}
			fmt.Printf("%s lance un sort de glace et inflige %d dégâts à %s !\n", hero.Nom, degats, monster.Nom)
			fmt.Printf("PV de %s : %d/%d\n", monster.Nom, monster.PV, monster.PVMax)
		} else if skill.Nom == "Vent" {
			// Sort de vent (attaque rapide)
			monster.PV -= degats
			if monster.PV < 0 {
				monster.PV = 0
			}
			fmt.Printf("%s invoque un vent puissant et inflige %d dégâts à %s !\n", hero.Nom, degats, monster.Nom)
			fmt.Printf("PV de %s : %d/%d\n", monster.Nom, monster.PV, monster.PVMax)
		} else if skill.Nom == "Terre" {
			// Sort de terre (dégâts moyens)
			monster.PV -= degats
			if monster.PV < 0 {
				monster.PV = 0
			}
			fmt.Printf("%s fait trembler la terre et inflige %d dégâts à %s !\n", hero.Nom, degats, monster.Nom)
			fmt.Printf("PV de %s : %d/%d\n", monster.Nom, monster.PV, monster.PVMax)
		} else {
			// Sort d'attaque
			monster.PV -= degats
			if monster.PV < 0 {
				monster.PV = 0
			}
			fmt.Printf("%s utilise %s et inflige %d degats a %s !\n", hero.Nom, skill.Nom, degats, monster.Nom)
			fmt.Printf("PV de %s : %d/%d\n", monster.Nom, monster.PV, monster.PVMax)
		}
		
		fmt.Printf("Mana restant : %d/%d\n", hero.Mana, hero.ManaMax)

		fmt.Println("\n(Appuyez sur Entree pour continuer)")
		fmt.Scanln()
		return
	}
}

// Combat d'entraînement principal
func TrainingFight(hero *character.Character) {
	monster := InitGoblin()
	tour := 1

	utils.ClearScreen()
	fmt.Println("=== Combat d'entrainement ===")
	fmt.Printf("Vous affrontez un %s !\n", monster.Nom)
	fmt.Printf("Initiative de %s : %d\n", hero.Nom, hero.Initiative)
	fmt.Printf("Initiative de %s : %d\n", monster.Nom, monster.Initiative)
	fmt.Println("\n(Appuyez sur Entree pour commencer)")
	fmt.Scanln()

	for {
		utils.ClearScreen()
		fmt.Printf("=== Tour %d ===\n", tour)
		fmt.Printf("PV de %s : %d/%d | Mana : %d/%d\n", hero.Nom, hero.PV, hero.PVMax, hero.Mana, hero.ManaMax)
		fmt.Printf("PV de %s : %d/%d\n", monster.Nom, monster.PV, monster.PVMax)
		fmt.Println()

		// Vérifier les conditions de fin
		if hero.PV <= 0 {
			fmt.Printf("%s est vaincu !\n", hero.Nom)
			hero.IsDead()
			fmt.Println("\n(Appuyez sur Entree pour continuer)")
			fmt.Scanln()
			return
		}
		if monster.PV <= 0 {
			fmt.Printf("%s est vaincu !\n", monster.Nom)
			// Gain d'expérience
			expGain := 50
			hero.AddExperience(expGain)
			fmt.Println("\n(Appuyez sur Entree pour continuer)")
			fmt.Scanln()
			return
		}

		// Déterminer qui joue en premier (initiative)
		heroFirst := hero.Initiative >= monster.Initiative
		
		if heroFirst {
			// Le héros joue en premier
			if hero.PV > 0 {
				CharacterTurn(hero, &monster)
			}
			
			// Vérifier si le monstre est encore vivant après l'attaque du joueur
			if monster.PV <= 0 {
				fmt.Printf("%s est vaincu !\n", monster.Nom)
				expGain := 50
				hero.AddExperience(expGain)
				fmt.Println("\n(Appuyez sur Entree pour continuer)")
				fmt.Scanln()
				return
			}
			
			// Tour du monstre
			if monster.PV > 0 {
				monster.GoblinPattern(hero, tour)
				fmt.Println("\n(Appuyez sur Entree pour continuer)")
				fmt.Scanln()
			}
		} else {
			// Le monstre joue en premier
			if monster.PV > 0 {
				monster.GoblinPattern(hero, tour)
				fmt.Println("\n(Appuyez sur Entree pour continuer)")
				fmt.Scanln()
			}
			
			// Vérifier si le héros est encore vivant après l'attaque du monstre
			if hero.PV <= 0 {
				fmt.Printf("%s est vaincu !\n", hero.Nom)
				hero.IsDead()
				fmt.Println("\n(Appuyez sur Entree pour continuer)")
				fmt.Scanln()
				return
			}
			
			// Tour du héros
			if hero.PV > 0 {
				CharacterTurn(hero, &monster)
			}
		}

		tour++
	}
}
