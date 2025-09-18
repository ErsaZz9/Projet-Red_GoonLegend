package items

type Item struct {
	Nom    string
	Type   string
	Soin   int
	Degats int
	Bonus  int  // Bonus d'attaque ou de défense
	Slot   string // "Weapon", "Armor", "Accessory"
}
