package main

import (
	"context"

	"log"
	"os"
	"runtime"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/nh3000-org/spades/config"
	"github.com/nh3000-org/spades/config/cards"
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

var NPC = NonPersonStat{}

type Deck struct {
	Cards []string
}

var C = Deck{}

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

func splash() {

	header := canvas.NewImageFromResource(cards.NewEmbeddedResource("honors_spade-14.png"))
	header.FillMode = canvas.ImageFillContain
	header.SetMinSize(fyne.NewSize(100, 100))

	rules := widget.NewMultiLineEntry()
	rules.SetText(config.GetLangs("rules"))

	config.PlayerName = config.FyneApp.Preferences().StringWithFallback("Player", "Player1")
	playerlabel := widget.NewLabel(config.GetLangs("player"))
	player := widget.NewEntry()
	player.SetText(config.PlayerName)
	player.SetPlaceHolder(config.GetLangs("player"))

	config.Difficulty = config.FyneApp.Preferences().StringWithFallback("Difficulty", "Easy")
	difficultylabel := widget.NewLabel(config.GetLangs("difficulty"))
	difficulty := widget.NewRadioGroup([]string{"Easy", "Hard"}, func(string) {})
	difficulty.SetSelected(config.Difficulty)
	difficulty.Horizontal = true

	config.DeckBack = config.FyneApp.Preferences().StringWithFallback("Deckback", "Grey")
	deckbacklabel := widget.NewLabel(config.GetLangs("deckback"))
	deckback := widget.NewRadioGroup([]string{"Red", "Yellow", "Purple", "Grey", "Green"}, func(string) {})
	deckback.SetSelected(config.DeckBack)
	deckback.Horizontal = false

	rightbox := container.NewVBox(
		widget.NewLabelWithStyle(config.GetLangs("preferences"), fyne.TextAlignCenter, fyne.TextStyle{Bold: true}),
		playerlabel,
		player,

		difficultylabel,
		difficulty,

		deckbacklabel,
		deckback,
	)

	next := widget.NewButton("Next", func() {

		config.FyneApp.Preferences().SetString("Player", player.Text)
		config.PlayerName = player.Text
		config.FyneApp.Preferences().SetString("Difficulty", difficulty.Selected)
		config.Difficulty = difficulty.Selected
		config.FyneApp.Preferences().SetString("Deckback", deckback.Selected)
		config.DeckBack = deckback.Selected
		deal()
	})
	border := container.NewBorder(header, next, nil, rightbox, rules)
	config.FyneMainWin.SetContent(border)

	// pick a user name
	// pick a dificulty
	// pick a deck back

	// deal a deck
	/* 	next := widget.NewButton("Next", func() {
		deal()
	}) */
	//img := cards.GetImage("honor_spades-14.png")

	//header := fyne.NewStaticResource("HonorSpades",cards.GetImage("CardsFS/honor_spades-14.png").Resource.Content())
	//h := container.NewCenter(header)
	//border := container.NewBorder(header, next, nil, nil, header)

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
	config.Selected = config.Game
	config.FyneApp.Settings().SetTheme(config.MyTheme{})
	MyLogo, iconerr := fyne.LoadResourceFromPath("icon.png")
	if iconerr != nil {
		log.Println("icon.png error ", iconerr.Error())
	}
	config.FyneApp.SetIcon(MyLogo)
	config.FyneMainWin.SetTitle(config.GetLangs("title"))
	config.FyneMainWin.Resize(fyne.NewSize(640, 480))
	splash()
	config.FyneMainWin.ShowAndRun()
}
