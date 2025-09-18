package menu

import (
	"Projet-Red_GoonLegend/character"
	"Projet-Red_GoonLegend/combat"
	"Projet-Red_GoonLegend/map"
	"Projet-Red_GoonLegend/utils"
	"Projet-Red_GoonLegend/shop"
	"fmt"

	"github.com/manifoldco/promptui"
)

func Run(hero *character.Character) {
	for {
		utils.ClearScreen()

		options := []string{
			"Afficher les informations du personnage",
			"Acceder a l'inventaire",
			"Marchand",
			"Forgeron",
			"Entrainement",
			"Explorer le monde",
			"Qui sont-ils?",
			"Quitter",
		}

		prompt := promptui.Select{
			Label: "=== Menu Principal ===",
			Items: options,
			Templates: &promptui.SelectTemplates{
				Label: "{{ . }}",
			},
			Size: 8,
		}

		index, _, err := prompt.Run()
		if err != nil {
			fmt.Println("Annule.")
			return
		}

		switch index {
		case 0:
			utils.ClearScreen()
			hero.DisplayInfo()
			fmt.Println("\n(Appuyez sur Entree pour continuer)")
			fmt.Scanln()
		case 1:
			accessInventoryMenu(hero)
		case 2:
			runShop(hero)
		case 3:
			runForge(hero) // defini dans forge_menu.go (package menu)
		case 4:
			combat.TrainingFight(hero)
		case 5:
			gameMap := gameMap.InitMap()
			gameMap.RunMapMenu(hero)
		case 6:
			showArtists()
		case 7:
			fmt.Println("Fermeture")
			return
		}
	}
}

// Sous-menu inventaire
func accessInventoryMenu(hero *character.Character) {
	for {
		utils.ClearScreen()

		options := make([]string, 0, len(hero.Inventaire)+1)
		for _, it := range hero.Inventaire {
			options = append(options, it.Nom)
		}
		options = append(options, "← Retour")

		prompt := promptui.Select{
			Label: "=== Inventaire ===",
			Items: options,
			Templates: &promptui.SelectTemplates{
				Label: "{{ . }}",
			},
			Size: 10,
		}

		index, _, err := prompt.Run()
		if err != nil {
			fmt.Println("Annule.")
			return
		}

		// Retour
		if index == len(options)-1 {
			return
		}

		item := hero.Inventaire[index]

		switch item.Nom {
		case "Potion de soin":
			hero.UsePotion()
		case "Potion de poison":
			hero.ApplyPoison(item.Degats, 3)
			hero.RemoveInventory(item.Nom)
		case "Livre de Sort : Boule de Feu":
			hero.SpellBook("Boule de Feu")
			hero.RemoveInventory(item.Nom)
			fmt.Println("Vous avez appris : Boule de Feu !")
		case "Livre de Sort : Soin":
			hero.SpellBook("Soin")
			hero.RemoveInventory(item.Nom)
			fmt.Println("Vous avez appris : Soin !")
		case "Livre de Sort : Eclair":
			hero.SpellBook("Eclair")
			hero.RemoveInventory(item.Nom)
			fmt.Println("Vous avez appris : Eclair !")
		case "Livre de Sort : Meteore":
			hero.SpellBook("Meteore")
			hero.RemoveInventory(item.Nom)
			fmt.Println("Vous avez appris : Meteore !")
		case "Livre de Sort : Glace":
			hero.SpellBook("Glace")
			hero.RemoveInventory(item.Nom)
			fmt.Println("Vous avez appris : Glace !")
		case "Livre de Sort : Vent":
			hero.SpellBook("Vent")
			hero.RemoveInventory(item.Nom)
			fmt.Println("Vous avez appris : Vent !")
		case "Livre de Sort : Terre":
			hero.SpellBook("Terre")
			hero.RemoveInventory(item.Nom)
			fmt.Println("Vous avez appris : Terre !")
		case "Chapeau de l'aventurier", "Tunique de l'aventurier", "Bottes de l'aventurier":
			hero.Equip(item)
		default:
			fmt.Printf("Impossible d'utiliser %s pour l'instant.\n", item.Nom)
		}

		fmt.Println("\n(Appuyez sur Entree pour continuer)")
		fmt.Scanln()
	}
}

// Sous-menu marchand (UNE SEULE DEFINITION)
func runShop(hero *character.Character) {
	for {
		utils.ClearScreen()

		// Construire le menu a partir de shop.ShopItems
		options := make([]string, 0, len(shop.ShopItems)+1)
		for _, s := range shop.ShopItems {
			options = append(options, fmt.Sprintf("Acheter %s (%d or)", s.Nom, s.Prix))
		}
		options = append(options, "← Retour")

		prompt := promptui.Select{
			Label: "=== Marchand ===",
			Items: options,
			Templates: &promptui.SelectTemplates{
				Label: "{{ . }}",
			},
			Size: 10,
		}

		index, _, err := prompt.Run()
		if err != nil {
			fmt.Println("Annule.")
			return
		}
		if index == len(options)-1 {
			return
		}

		choix := shop.ShopItems[index]

		// Rappel du cout
		fmt.Printf("\nObjet: %s\nCout : %d or\n\n", choix.Nom, choix.Prix)

		// 1) Or insuffisant
		if hero.Purse < choix.Prix {
			fmt.Printf("Vous n'avez pas assez d'or pour acheter %s !\n", choix.Nom)
			fmt.Println("\n(Appuyez sur Entree pour continuer)")
			fmt.Scanln()
			continue
		}

		// 2) Items de type Upgrade -> effet direct (pas de place requise)
		if choix.Item.Type == "Upgrade" {
			hero.Purse -= choix.Prix
			hero.UpgradeInventorySlot()
			fmt.Printf("Vous avez achete : %s pour %d or\n", choix.Nom, choix.Prix)
			fmt.Printf("Or restant : %d\n", hero.Purse)
			fmt.Println("\n(Appuyez sur Entree pour continuer)")
			fmt.Scanln()
			continue
		}

		// 3) Items normaux -> verifier la place
		if hero.IsInventoryFull() {
			fmt.Printf("Vous ne pouvez pas stocker plus de %d objets dans votre inventaire !\n", hero.MaxInventory)
			fmt.Println("\n(Appuyez sur Entree pour continuer)")
			fmt.Scanln()
			continue
		}

		// 4) Achat standard
		hero.Purse -= choix.Prix
		hero.AddInventory(choix.Item)
		fmt.Printf("Vous avez achete : %s pour %d or\n", choix.Nom, choix.Prix)
		fmt.Printf("Or restant : %d\n", hero.Purse)
		fmt.Println("\n(Appuyez sur Entree pour continuer)")
		fmt.Scanln()
	}
}

// Affiche les artistes cachés
func showArtists() {
	utils.ClearScreen()
	fmt.Println("=== Qui sont-ils? ===")
	fmt.Println()
	fmt.Println("Les artistes cachés dans ce projet sont :")
	fmt.Println()
	fmt.Println("🎵 **TASK 2** - Two for the Price of One")
	fmt.Println("   Artiste : **Madonna**")
	fmt.Println("   Référence : 'Two for the Price of One' (1982)")
	fmt.Println()
	fmt.Println("🎵 **TASK 3** - Gimme! Gimme! Gimme!")
	fmt.Println("   Artiste : **ABBA**")
	fmt.Println("   Référence : 'Gimme! Gimme! Gimme! (A Man After Midnight)' (1979)")
	fmt.Println()
	fmt.Println("Ces références musicales sont cachées dans les noms des tâches !")
	fmt.Println()
	fmt.Println("(Appuyez sur Entree pour continuer)")
	fmt.Scanln()
}
