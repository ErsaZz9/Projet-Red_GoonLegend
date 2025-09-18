package items

type Item struct {
	Nom    string // ex: "Potion de soin", "Épée en bois"
	Type   string // ex: "Potion", "Arme", "Livre de sort"
	Prix   int    // prix chez le marchand
	Soin   int    // valeur de soin (ex: +50 PV pour une potion)
	Degats int    // dégâts infligés (ex: arme ou poison)
}
