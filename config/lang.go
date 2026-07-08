package config

var MyLangs = map[string]string{
	"eng-title": "Spades 4 Two",
	"spa-title": "Picas 4 Dos",
	"hin-title": "हुकुम 4 दो",

	"eng-preferences": "Preferences",
	"spa-dificulty":   "Preferencias",
	"hin-dificulty":   "प्राथमिकताएँ",

	"eng-scale": "Scale",
	"spa-scale": "Escala",
	"hin-scale": "पैमाना",

	"eng-rules": "Draw\n-Pick a Card\n-Keep or Discard The Draw\n\nBid\n-Number of Tricks\n-Nil\n-Blind Nil\n-Blind 13\n\nScoring\n-Tricks * 10\n-Nil 100 If Made Minus 100 If Failure\n-Blind Nil 200 If Made Minus 200 If Failure\n-Blind 13 200 If Made Minus 200 if Failure\n\nBags\n-Over Tricks\n-10 Bags Minus 100",
	"spa-rules": "Robar\n-Elige una carta\n-Conserva o descarta la carta robada\n\nApuesta\n-Número de bazas\n-Cero\n-Ciego Cero\n-Ciego 13\n\nPuntuación\n-Babas * 10\n-Cero 100 si se logra -100 si falla\n-Ciego Cero 200 si se logra -200 si falla\n-Ciego 13 200 si se logra -200 si falla\n\nBolsas\n-Bazas adicionales\n-10 bolsas -100",
	"hin-rules": "ड्रा\n-कार्ड चुनें\n-ड्रा को रखें या छोड़ें\n\nबिड\n-ट्रिक्स की संख्या\n-शून्य\n-ब्लाइंड शून्य\n-ब्लाइंड 13\n\nस्कोरिंग\n-ट्रिक्स * 10\n-शून्य 100 अगर बनाया तो माइनस 100 अगर फेलियर\n-ब्लाइंड शून्य 200 अगर बनाया तो माइनस 200 अगर फेलियर\n-ब्लाइंड 13 200 अगर बनाया तो माइनस 200 अगर फेलियर\n\nबैग्स\n-ओवर ट्रिक्स\n-10 बैग्स माइनस 100",

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
