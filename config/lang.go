package config

var MyLangs = map[string]string{
	"eng-title": "Spades 4 Two",
	"spa-title": "Picas 4 Dos",
	"hin-title": "हुकुम 4 दो",

	"eng-preferences": "Preferences",
	"spa-dificulty":   "Preferencias",
	"hin-dificulty":   "प्राथमिकताएँ",

	"eng-rules": "Draw\n-Pick a Card\n-Keep or Discard The Draw\n\nBid\n-Number of Tricks\n-Nil\n-Blind Nil\n\nScoring\n-Tricks * 10\n-Nil 100 If Made\n-Nil Minus 100 If Failure\n-Blind Nil 200 If Made\n-Vlind Nill Minus 200 If Failure\n\nBags\n-Over Tricks\n-10 Bags Minus 100",
	"spa-rules": "Robar\n-Elige una carta\n-Quedarse o descartar el robo\n\nPujar\n-Número de bazas\n-Nula\n-Ciego 0\n\nPuntuación\n-Bazas * 10\n-Nula 100 si se hace\n-Nula menos 100 si falla\n-Ciego 0 200 si se hace\n-Nula menos 200 si falla\n\nBolsas\n-Más de bazas\n-10 bolsas menos 100",
	"hin-rules": "ड्रा\n-कार्ड चुनें\n-ड्रा को रखें या छोड़ें\n\nबोली\n-ट्रिक्स की संख्या\n-शून्य\n-ब्लाइंड शून्य\n\nस्कोरिंग\n-ट्रिक्स * 10\n-शून्य 100 अगर बने\n-शून्य माइनस 100 अगर फेलियर\n-ब्लाइंड शून्य 200 अगर बने\n-वीलिंड शून्य माइनस 200 अगर फेलियर\n\nबैग्स\n-ओवर ट्रिक्स\n-10 बैग्स माइनस 100",

	"eng-player": "Player Name",
	"spa-player": "Nombre del Jugador",
	"hin-player": "खिलाड़ी का नाम",

	"eng-difficulty": "Difficulty",
	"spa-difficulty": "Dificultad",
	"hin-difficulty": "कठिनाई",

	"eng-deckback": "Select a Deck Back",
	"spa-deckback": "Seleccione una Cubierta Trasera",
	"hin-deckback": "एक डेक बैक चुनें",
}
var PreferedLanguage string

func GetLangs(c string) string {
	value, err := MyLangs[PreferedLanguage+"-"+c]
	if !err {
		return "lang-error" + " " + PreferedLanguage + "-" + c
	}

	return value
}
