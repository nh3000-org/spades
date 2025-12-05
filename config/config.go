package config

import (
	"log"

	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/storage"
)

var FyneApp fyne.App
var FyneMainWin fyne.Window
var PlayerName string
var DealerPlayer bool
var Difficulty string
var DeckBack string // red yellow purple grey green
var PlayerBid int   // 1-13 14 for nil 15 for blind nil
var PlayerBags int
var NPCBid int      // 1-13 14 for nil 15 for blind nil
var NPCBags int
var PlayerScore int
var NPCScore int
var PlayerHistory int
var NPCHistory int
var PlayerTricks int
var NPCTricks int
func GetDateTime(offsethours string) string {
	ct := time.Now()

	hours, _ := time.ParseDuration(offsethours)
	future := ct.Add(hours)

	return future.String()
}

func DataStore(file string) fyne.URI {
	DataLocation, dlerr := storage.Child(FyneApp.Storage().RootURI(), file)
	if dlerr != nil {
		log.Println("DataStore error ", dlerr)
	}
	return DataLocation
}
