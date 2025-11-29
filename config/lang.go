package config

import "log"

var MyLangs = map[string]string{
	"eng-title": "Spades 4 Two",
	"spa-title": "Picas 4 Dos",
	"hin-title": "हुकुम 4 दो",
}
var PreferedLanguage string

func GetLangs(c string) string {
	value, err := MyLangs[PreferedLanguage+"-"+c]
	if !err {
		return "lang-error" + " " + PreferedLanguage + "-" + c
	}

	return value
}
