package main

import (
	"strconv"

	tc "image/color"
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
	getcards "github.com/nh3000-org/spades/config/cards"
)

type Player struct {
	deck *Deck
	hand Cards
}

var PlayerGame config.PS
var NPCGame config.PS

var memoryStats runtime.MemStats
var spadesheader *canvas.Image

// var deckbackimage *canvas.Image
var deckbackname string

func discard(player bool) {

}
func keep(player bool) {

}
func pick(player bool) {

}

func hand(player bool) {

}

func deal() {
	runtime.GC()
	runtime.ReadMemStats(&memoryStats)

	config.FyneMainWin.SetTitle(config.GetLangs("title") + " " + strconv.FormatUint(memoryStats.Alloc/1024/1024, 10) + " Mib")
	//deckbackimage.FillMode = canvas.ImageFillContain
	//deckbackimage.SetMinSize(fyne.NewSize(50, 50))
	labelcolor := tc.RGBA{253, 118, 87, 255}
	// player status
	playerbidlabel := canvas.NewText("Bid:", labelcolor)
	playerbidlabel.TextSize = 32
	playerbidbagslabel := canvas.NewText("Bags:", labelcolor)
	playerbidbagslabel.TextSize = 32
	playerbidscorelabel := canvas.NewText("Score:", labelcolor)
	playerbidscorelabel.TextSize = 32
	playerbidtrickslabel := canvas.NewText("Tricks:", labelcolor)
	playerbidtrickslabel.TextSize = 32
	playerbidname := canvas.NewText(config.PlayerName, tc.White)
	playerbidname.TextSize = 32
	playerbid := canvas.NewText("0", tc.White)
	playerbid.TextSize = 32
	playerbags := canvas.NewText("0", tc.White)
	playerbags.TextSize = 32
	playerscore := canvas.NewText("0", tc.White)
	playerscore.TextSize = 32
	playertricks := canvas.NewText("0", tc.White)
	playertricks.TextSize = 32
	bidplayer := container.NewHBox(
		playerbidname,
		playerbidlabel,
		playerbid,
		playerbidbagslabel,
		playerbags,
		playerbidscorelabel,
		playerscore,
		playerbidtrickslabel,
		playertricks,
	)

	// npc status
	npcbidlabel := canvas.NewText("Bid:", labelcolor)
	npcbidlabel.TextSize = 32
	npcbidbagslabel := canvas.NewText("Bags:", labelcolor)
	npcbidbagslabel.TextSize = 32
	npcbidscorelabel := canvas.NewText("Score:", labelcolor)
	npcbidscorelabel.TextSize = 32
	npcbidtrickslabel := canvas.NewText("Tricks:", labelcolor)
	npcbidtrickslabel.TextSize = 32
	npcbidname := canvas.NewText("NPC", tc.White)
	npcbidname.TextSize = 32
	npcbid := canvas.NewText("0", tc.White)
	npcbid.TextSize = 32
	npcbags := canvas.NewText("0", tc.White)
	npcbags.TextSize = 32
	npcscore := canvas.NewText("0", tc.White)
	npcscore.TextSize = 32
	npctricks := canvas.NewText("0", tc.White)
	npctricks.TextSize = 32
	bidnpc := container.NewHBox(
		npcbidname,
		npcbidlabel,
		npcbid,
		npcbidbagslabel,
		npcbags,
		npcbidscorelabel,
		npcscore,
		npcbidtrickslabel,
		npctricks,
	)

	playerdiscard := canvas.NewText(" Player Discards 0", tc.White)
	playerhand := container.NewHBox()

	npcdiscard := canvas.NewText("NPC Discards 0", tc.White)

	keep := widget.NewButton("Keep", func() {

	})
	discard := widget.NewButton("Discard", func() {

	})
	blindnil := widget.NewButton("Blind Nil", func() {

	})
	regularnil := widget.NewButton("Nil", func() {

	})
	bid := widget.NewButton("Bid", func() {

	})
	bidbar := container.NewGridWithColumns(5)
	bidbar.Add(keep)
	bidbar.Add(discard)
	bidbar.Add(blindnil)
	bidbar.Add(regularnil)
	bidbar.Add(bid)

	card := canvas.NewImageFromResource(getcards.NewEmbeddedResource("2C.png"))
	deckbackimage := canvas.NewImageFromResource(getcards.NewEmbeddedResource("AS.png"))
	deckbackimage.SetMinSize(fyne.NewSize(100, 100))
	card.FillMode = canvas.ImageFillContain
	card.SetMinSize(fyne.NewSize(100, 100))
	playerhand.Add(card)
	playerboard := container.NewBorder(bidbar, nil, nil, nil, card)
	bidboard := container.NewBorder(nil, playerboard, nil, nil, deckbackimage)
	gameboard := container.NewBorder(bidnpc, bidplayer, npcdiscard, playerdiscard, bidboard)
	d := NewDeck()
	d.Shuffle()

	gplayer := Player{deck: &d}
	dc, dcerr := d.Draw()
	if dcerr != nil {
		log.Println("draw card ", dcerr)
	}
	log.Println("draw card ", dc.Rank, dc.Suit, dc.String(), dc.Rank.String(), dc.Suit.String(), gplayer.hand)
	mycard := dc.Rank.String() + dc.Suit.String() + ".png"
	/// make 52 image + deckback as separate the dislay
	deckbackimage.Resource = getcards.NewEmbeddedResource(mycard)
	deckbackimage.SetMinSize(fyne.NewSize(100, 100))
	deckbackimage.Refresh()
	deckbackimage.Show()

	//gplayer := Player{deck: &d}

	bidboard.Refresh()
	gameboard.Refresh()
	config.FyneMainWin.Canvas().Refresh(deckbackimage)
	config.FyneMainWin.SetContent(gameboard)

	//config.FyneMainWin.Canvas().Refresh(deckbackimage)

}

func splash() {
	runtime.GC()
	runtime.ReadMemStats(&memoryStats)
	config.FyneMainWin.SetTitle(config.GetLangs("title") + " " + strconv.FormatUint(memoryStats.Alloc/1024/1024, 10) + " Mib")
	spadesheader = canvas.NewImageFromResource(getcards.NewEmbeddedResource("honors_spade-14.png"))
	spadesheader.FillMode = canvas.ImageFillContain
	spadesheader.SetMinSize(fyne.NewSize(30, 30))

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

		//deckbackname = strings.ToLower(deckback.Selected) + "_back.png"
		//deckbackimage := canvas.NewImageFromResource(getcards.NewEmbeddedResource(deckbackname))
		config.DeckBack = deckback.Selected
		config.DealerPlayer = true
		PlayerGame = config.NewPlayer(player.Text)
		NPCGame = config.NewPlayer("NPC")
		deal()

	})
	border := container.NewBorder(spadesheader, next, nil, rightbox, rules)
	config.FyneMainWin.SetContent(border)

}
func main() {
	a := app.NewWithID("org.nh3000.spades")
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
	runtime.GC()
	runtime.ReadMemStats(&memoryStats)
	config.FyneApp.SetIcon(MyLogo)
	config.FyneMainWin.SetTitle(config.GetLangs("title") + " " + strconv.FormatUint(memoryStats.Alloc/1024/1024, 10) + " Mib")
	config.FyneMainWin.Resize(fyne.NewSize(640, 480))
	splash()
	config.FyneMainWin.ShowAndRun()
}
