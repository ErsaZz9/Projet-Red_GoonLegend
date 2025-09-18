package forge

import "Projet-Red_GoonLegend/items"

// Recette de forge
type Recipe struct {
	Nom       string
	Cout      int
	Materiaux map[string]int
	Resultat  items.Item
}

// Toutes les recettes disponibles
var ForgeRecipes = []Recipe{
	{
		Nom:  "Chapeau de l'aventurier",
		Cout: 5,
		Materiaux: map[string]int{
			"Plume de Corbeau": 1,
			"Cuir de Sanglier": 1,
		},
		Resultat: items.Item{Nom: "Chapeau de l'aventurier", Type: "Equipement"},
	},
	{
		Nom:  "Tunique de l'aventurier",
		Cout: 5,
		Materiaux: map[string]int{
			"Fourrure de Loup": 2,
			"Peau de Troll":    1,
		},
		Resultat: items.Item{Nom: "Tunique de l'aventurier", Type: "Equipement"},
	},
	{
		Nom:  "Bottes de l'aventurier",
		Cout: 5,
		Materiaux: map[string]int{
			"Fourrure de Loup": 1,
			"Cuir de Sanglier": 1,
		},
		Resultat: items.Item{Nom: "Bottes de l'aventurier", Type: "Equipement"},
	},
}
