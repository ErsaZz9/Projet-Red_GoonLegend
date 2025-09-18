package classes

// Définition d’une classe jouable
type Classe struct {
	Nom         string
	PVMax       int
	ManaMax     int
	Force       int
	Dexterite   int
	Description string
}

// Liste des classes disponibles
var Classes = []Classe{
	{"Humain", 100, 50, 10, 10, "Polyvalent"},
	{"Elfe", 80, 100, 5, 15, "Rapide et agile"},
	{"Nain", 120, 30, 15, 5, "Solide et robuste"},
}
