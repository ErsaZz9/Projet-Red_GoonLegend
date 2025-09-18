package gameMap

import (
	"Projet-Red_GoonLegend/character"
	"Projet-Red_GoonLegend/combat"
	"Projet-Red_GoonLegend/items"
	"Projet-Red_GoonLegend/utils"
	"fmt"

	"github.com/manifoldco/promptui"
)

// Structure pour une zone de la map
type Zone struct {
	Nom        string
	Description string
	Enemies    []combat.Monster
	Treasures  []string
	IsCleared  bool
	NextZones  []string
	GoldReward int
}

// Structure pour la map complète
type GameMap struct {
	CurrentZone string
	Zones       map[string]Zone
	Visited     map[string]bool
}

// Initialise la map du jeu
func InitMap() GameMap {
	zones := make(map[string]Zone)
	
	// Zone 1: Village (zone de départ)
	zones["village"] = Zone{
		Nom:        "Village de Labubu",
		Description: "Un petit village paisible où les Labubus vivent en harmonie... ou pas ?",
		Enemies:    []combat.Monster{},
		Treasures:  []string{"Potion de soin", "Potion de mana"},
		IsCleared:  true,
		NextZones:  []string{"foret"},
		GoldReward: 10,
	}
	
	// Zone 2: Forêt Labubu
	zones["foret"] = Zone{
		Nom:        "Forêt Sombre des Labubus",
		Description: "Une forêt mystérieuse où les Labubus maléfiques se cachent dans l'ombre.",
		Enemies:    []combat.Monster{
			{Nom: "Labubu Sombre", PVMax: 60, PV: 60, Attaque: 8, Initiative: 5},
			{Nom: "Labubu Toxique", PVMax: 45, PV: 45, Attaque: 6, Initiative: 7},
		},
		Treasures:  []string{"Livre de Sort : Soin", "Fourrure de Loup"},
		IsCleared: false,
		NextZones:  []string{"chateau"},
		GoldReward: 25,
	}
	
	// Zone 3: Château des Labubus
	zones["chateau"] = Zone{
		Nom:        "Château de la Reine Labubu",
		Description: "Le château sombre où règne la terrible Reine Labubu et ses gardes.",
		Enemies:    []combat.Monster{
			{Nom: "Garde Labubu", PVMax: 80, PV: 80, Attaque: 10, Initiative: 6},
			{Nom: "Reine Labubu", PVMax: 150, PV: 150, Attaque: 15, Initiative: 8},
		},
		Treasures:  []string{"Épée de Labubu", "Couronne de la Reine"},
		IsCleared: false,
		NextZones:  []string{},
		GoldReward: 50,
	}
	
	return GameMap{
		CurrentZone: "village",
		Zones:      zones,
		Visited:    make(map[string]bool),
	}
}

// Affiche la map actuelle
func (gm *GameMap) DisplayMap() {
	utils.ClearScreen()
	fmt.Println("=== CARTE DU MONDE LABUBU ===")
	fmt.Println()
	
	// Afficher la zone actuelle
	current := gm.Zones[gm.CurrentZone]
	fmt.Printf("📍 Zone actuelle: %s\n", current.Nom)
	fmt.Printf("📝 %s\n", current.Description)
	fmt.Println()
	
	// Afficher les zones disponibles
	fmt.Println("🗺️  Zones disponibles:")
	for _, zoneName := range current.NextZones {
		zone := gm.Zones[zoneName]
		status := "🔒"
		if gm.Visited[zoneName] {
			status = "✅"
		}
		fmt.Printf("  %s %s\n", status, zone.Nom)
	}
	
	fmt.Println()
	fmt.Println("Options:")
	fmt.Println("1. Explorer la zone actuelle")
	fmt.Println("2. Se déplacer vers une autre zone")
	fmt.Println("3. Retour au menu principal")
}

// Menu principal de la map
func (gm *GameMap) RunMapMenu(hero *character.Character) {
	for {
		gm.DisplayMap()
		
		options := []string{
			"Explorer la zone actuelle",
			"Se déplacer vers une autre zone",
			"Retour au menu principal",
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
			gm.ExploreCurrentZone(hero)
		case 1:
			gm.MoveToZone(hero)
		case 2:
			return
		}
	}
}

// Explorer la zone actuelle
func (gm *GameMap) ExploreCurrentZone(hero *character.Character) {
	zone := gm.Zones[gm.CurrentZone]
	
	utils.ClearScreen()
	fmt.Printf("=== Exploration de %s ===\n", zone.Nom)
	fmt.Println()
	
	if zone.Nom == "Village de Labubu" {
		fmt.Println("Vous vous promenez dans le village...")
		fmt.Println("Les habitants vous donnent quelques potions et de l'or !")
		
		// Ajouter des potions et de l'or
		hero.AddInventory(items.Item{Nom: "Potion de soin", Type: "Potion", Soin: 50})
		hero.AddInventory(items.Item{Nom: "Potion de mana", Type: "Potion", Soin: 30})
		hero.Purse += zone.GoldReward
		fmt.Printf("Vous gagnez %d or !\n", zone.GoldReward)
		
		fmt.Println("\n(Appuyez sur Entree pour continuer)")
		fmt.Scanln()
		return
	}
	
	// Zones avec ennemis
	if len(zone.Enemies) > 0 {
		fmt.Println("Vous rencontrez des ennemis !")
		fmt.Println()
		
		for i, enemy := range zone.Enemies {
			fmt.Printf("%d. %s (PV: %d, Attaque: %d)\n", i+1, enemy.Nom, enemy.PV, enemy.Attaque)
		}
		fmt.Println("0. Fuir")
		
		options := make([]string, 0, len(zone.Enemies)+1)
		for _, enemy := range zone.Enemies {
			options = append(options, fmt.Sprintf("Combattre %s", enemy.Nom))
		}
		options = append(options, "Fuir")
		
		prompt := promptui.Select{
			Label: "Que voulez-vous faire ?",
			Items: options,
		}
		
		index, _, err := prompt.Run()
		if err != nil {
			return
		}
		
		if index < len(zone.Enemies) {
			// Combat contre l'ennemi sélectionné
			enemy := zone.Enemies[index]
			fmt.Printf("Vous affrontez %s !\n", enemy.Nom)
			fmt.Println("\n(Appuyez sur Entree pour commencer le combat)")
			fmt.Scanln()
			
			// Lancer le combat
			combat.TrainingFight(hero)
			
			// Récompenses après victoire
			fmt.Printf("🎉 Victoire ! Vous gagnez %d or !\n", zone.GoldReward)
			hero.Purse += zone.GoldReward
			
			// Vérifier si la zone est maintenant débloquée
			if zone.Nom == "Forêt Sombre des Labubus" && len(zone.Enemies) == 1 {
				fmt.Println("🎉 Vous avez vaincu tous les ennemis de cette zone !")
				fmt.Println("La zone suivante est maintenant accessible !")
			}
		}
	} else {
		fmt.Println("Cette zone est vide...")
	}
	
	fmt.Println("\n(Appuyez sur Entree pour continuer)")
	fmt.Scanln()
}

// Se déplacer vers une autre zone
func (gm *GameMap) MoveToZone(hero *character.Character) {
	current := gm.Zones[gm.CurrentZone]
	
	if len(current.NextZones) == 0 {
		fmt.Println("Aucune zone accessible depuis ici.")
		fmt.Println("\n(Appuyez sur Entree pour continuer)")
		fmt.Scanln()
		return
	}
	
	options := make([]string, 0, len(current.NextZones))
	for _, zoneName := range current.NextZones {
		zone := gm.Zones[zoneName]
		options = append(options, zone.Nom)
	}
	
	prompt := promptui.Select{
		Label: "Vers quelle zone voulez-vous aller ?",
		Items: options,
	}
	
	index, _, err := prompt.Run()
	if err != nil {
		return
	}
	
	zoneName := current.NextZones[index]
	gm.CurrentZone = zoneName
	gm.Visited[zoneName] = true
	
	fmt.Printf("Vous vous dirigez vers %s...\n", gm.Zones[zoneName].Nom)
	fmt.Println("\n(Appuyez sur Entree pour continuer)")
	fmt.Scanln()
}
