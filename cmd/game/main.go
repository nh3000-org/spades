package main

import (
	"context"
	"log"
	"os"
	"runtime"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"github.com/nh3000-org/spades/config"
)

type PlayerStat struct {
	Score  string
	Bags   string
	Bids   string
	Tricks string
	Hand   string
}

var PS = PlayerStat{}

type NonPersonStat struct {
	Score  string
	Bags   string
	Bids   string
	Tricks string
	Hand   string
}

var NP = NonPersonStat{}

type Deck struct {
	Cards []string
}

var C = Deck{}

var Rules = "Draw\n-Pick a Card\n-Keep or Discard The Draw\n\nBid\n-Number of Tricks\n-Nil\n\nScoring\n-Tricks * 10\n-Nil 100 If Made\nNil Minus 100 If Failure\n-Blind Nil 200 If Made\n-Vlind Nill Minus 200 If Failure\n\nBags\n-Over Tricks\n-10 Bags Minus 100"

var memoryStats runtime.MemStats
var ctxmain context.Context
var ctxmaincan context.CancelFunc

func discard() {

}
func keep() {

}
func pick() {

}
func deal() {

}
func main() {
	var a = app.NewWithID("org.nh3000.spades")
	config.FyneApp = a
	var w = a.NewWindow("Spades 4 Two")
	config.FyneMainWin = w
	config.PreferedLanguage = "eng"
	if strings.HasPrefix(os.Getenv("LANG"), "en") {
		config.PreferedLanguage = "eng"
	}
	if strings.HasPrefix(os.Getenv("LANG"), "sp") {
		config.PreferedLanguage = "spa"
	}
	if strings.HasPrefix(os.Getenv("LANG"), "hn") {
		config.PreferedLanguage = "hin"
	}
	config.Selected = config.Dark
	config.FyneApp.Settings().SetTheme(config.MyTheme{})
	MyLogo, iconerr := fyne.LoadResourceFromPath("icon.png")
	if iconerr != nil {
		log.Println("icon.png error ", iconerr.Error())
	}
	config.FyneApp.SetIcon(MyLogo)
	config.FyneMainWin.SetTitle(config.GetLangs("title"))
	config.FyneMainWin.Resize(fyne.NewSize(640, 480))
	config.FyneMainWin.ShowAndRun()
}
