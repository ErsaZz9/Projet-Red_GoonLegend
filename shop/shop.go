package shop

import "Projet-Red_GoonLegend/items"

type MarchandItem struct {
	Nom   string
	Prix  int
	Item  items.Item
}

var ShopItems = []MarchandItem{
	{Nom: "Potion de soin", Prix: 3, Item: items.Item{Nom: "Potion de soin", Type: "Potion", Soin: 50}},
	{Nom: "Potion de mana", Prix: 5, Item: items.Item{Nom: "Potion de mana", Type: "Potion", Soin: 30}},
	{Nom: "Potion de poison", Prix: 6, Item: items.Item{Nom: "Potion de poison", Type: "Potion", Degats: 10}},
	{Nom: "Livre de Sort : Boule de Feu", Prix: 25, Item: items.Item{Nom: "Livre de Sort : Boule de Feu", Type: "Livre"}},
	{Nom: "Livre de Sort : Soin", Prix: 20, Item: items.Item{Nom: "Livre de Sort : Soin", Type: "Livre"}},
	{Nom: "Livre de Sort : Eclair", Prix: 35, Item: items.Item{Nom: "Livre de Sort : Eclair", Type: "Livre"}},
	{Nom: "Livre de Sort : Meteore", Prix: 50, Item: items.Item{Nom: "Livre de Sort : Meteore", Type: "Livre"}},
	{Nom: "Livre de Sort : Glace", Prix: 30, Item: items.Item{Nom: "Livre de Sort : Glace", Type: "Livre"}},
	{Nom: "Livre de Sort : Vent", Prix: 25, Item: items.Item{Nom: "Livre de Sort : Vent", Type: "Livre"}},
	{Nom: "Livre de Sort : Terre", Prix: 40, Item: items.Item{Nom: "Livre de Sort : Terre", Type: "Livre"}},
	{Nom: "Épée en fer", Prix: 15, Item: items.Item{Nom: "Épée en fer", Type: "Arme", Bonus: 5, Slot: "Weapon"}},
	{Nom: "Arc en bois", Prix: 12, Item: items.Item{Nom: "Arc en bois", Type: "Arme", Bonus: 3, Slot: "Weapon"}},
	{Nom: "Bâton magique", Prix: 20, Item: items.Item{Nom: "Bâton magique", Type: "Arme", Bonus: 2, Slot: "Weapon"}},
	{Nom: "Armure de cuir", Prix: 25, Item: items.Item{Nom: "Armure de cuir", Type: "Armure", Bonus: 3, Slot: "Armor"}},
	{Nom: "Cape magique", Prix: 18, Item: items.Item{Nom: "Cape magique", Type: "Armure", Bonus: 1, Slot: "Armor"}},
	{Nom: "Anneau de force", Prix: 30, Item: items.Item{Nom: "Anneau de force", Type: "Accessoire", Bonus: 2, Slot: "Accessory"}},
	{Nom: "Fourrure de Loup", Prix: 4, Item: items.Item{Nom: "Fourrure de Loup", Type: "Ressource"}},
	{Nom: "Peau de Troll", Prix: 7, Item: items.Item{Nom: "Peau de Troll", Type: "Ressource"}},
	{Nom: "Cuir de Sanglier", Prix: 3, Item: items.Item{Nom: "Cuir de Sanglier", Type: "Ressource"}},
	{Nom: "Plume de Corbeau", Prix: 1, Item: items.Item{Nom: "Plume de Corbeau", Type: "Ressource"}},
	{Nom: "Augmentation d'inventaire", Prix: 30, Item: items.Item{Nom: "Augmentation d'inventaire", Type: "Upgrade"}},

}

