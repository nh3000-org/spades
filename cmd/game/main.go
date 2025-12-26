package main

import (
	"bytes"
	"image"
	"strconv"

	tc "image/color"
	"image/png"
	"log"
	"os"
	"runtime"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
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

func discard(player bool) {

}
func keep(player bool) {

}
func pick(player bool) {

}

func hand(player bool) {

}
func cropImage(c string) *canvas.Image {
	log.Println("cropImage ", c)
	back := getcards.NewEmbeddedResource(c)
	img, _, err := image.Decode(bytes.NewReader(back.Content()))
	type subImager interface {
		SubImage(r image.Rectangle) image.Image
	}
	if err != nil {
		log.Println("image does not support cropping")
	}
	// img is an Image interface. This checks if the underlying value has a
	// method called SubImage. If it does, then we can use SubImage to crop the
	// image.
	simg, ok := img.(subImager)
	if !ok {
		log.Println("image does not support cropping")
	}
	my_sub_image := simg.(interface {
		SubImage(r image.Rectangle) image.Image
	}).SubImage(image.Rect(0, 0, 200, 300))
	log.Println(my_sub_image.Bounds())
	var b []byte
	err1 := png.Encode(bytes.NewBuffer(b), my_sub_image)
	if err1 != nil {
		log.Println("inithand", err1)

	}
	cardimg := canvas.NewImageFromImage(my_sub_image)

	return cardimg
}

var CardsLeft *canvas.Text
var PlayerBid *canvas.Text
var PlayerBags *canvas.Text
var PlayerScore *canvas.Text
var PlayerTricks *canvas.Text
var PlayerScoreBar *fyne.Container
var NPCBid *canvas.Text
var NPCBags *canvas.Text
var NPCScore *canvas.Text
var NPCTricks *canvas.Text
var NPCScoreBar *fyne.Container

var PlayerCards *fyne.Container
var NPCCards *fyne.Container

var PlayerTurn = "PLAYER"
var TurnCount int // 1 is first turn // 2 is second trun
var Buttonkeep = 1
var Buttondiscard = 2
var Buttonblindnil = 3
var Buttonregularnil = 4
var done = false

func Turn(button int) {
	log.Println("Turn count", TurnCount, "playerturn", PlayerTurn)
	TurnCount++
	done = false
	if TurnCount > 2 {
		if PlayerTurn == "PLAYER" {
			PlayerTurn = "NPC"
			done = true
		}
		if PlayerTurn == "NPC" && !done {
			PlayerTurn = "PLAYER"
		}
		TurnCount = 1
	}
	log.Println("Turn count", TurnCount, "playerturn", PlayerTurn)
	if TurnCount == 1 {
		// player
		if PlayerTurn == "PLAYER" {
			if button == Buttonkeep {
				Playerkeep.Disable()
				Playerdiscard.Enable()
				Playerblindnil.Disable()
				Playerregularnil.Disable()
				Playerbid.Disable()
			}
			if button == Buttondiscard {
				Playerkeep.Enable()
				Playerdiscard.Disable()
				Playerblindnil.Disable()
				Playerregularnil.Disable()
				Playerbid.Disable()
			}
			HandleCard()
		}
		// npc
		if PlayerTurn == "NPC" {

			Playerkeep.Disable()
			Playerdiscard.Enable()
			Playerblindnil.Disable()
			Playerregularnil.Disable()
			Playerbid.Disable()
			Playerkeep.Tapped(nil)
			HandleCard()

		}
	}
	if TurnCount == 2 {
		if PlayerTurn == "PLAYER" {
			if button == Buttonkeep {
				Playerkeep.Disable()
				Playerdiscard.Enable()
				Playerblindnil.Disable()
				Playerregularnil.Disable()
				Playerbid.Disable()
			}
			if button == Buttondiscard {
				Playerkeep.Enable()
				Playerdiscard.Disable()
				Playerblindnil.Disable()
				Playerregularnil.Disable()
				Playerbid.Disable()
			}
			HandleCard()
		}
		if PlayerTurn == "NPC" {

			NPCkeep.Enable()
			NPCdiscard.Disable()
			NPCblindnil.Disable()
			NPCregularnil.Disable()
			NPCbid.Disable()
			NPCdiscard.Tapped(nil)

			HandleCard()
		}

	}

}

var Mycardimage canvas.Image

func HandleCard() {

	dc, dcerr := MyDeck.Draw()
	if dcerr != nil {
		log.Println("draw card ", dcerr)
	}
	//CardsLeft = canvas.NewText(strconv.Itoa(len(MyDeck)), tc.White)
	//CardsLeft.TextSize = 32
	CardsLeft.Text = strconv.Itoa(len(MyDeck))

	log.Println("draw card ", strconv.Itoa(len(MyDeck)), dc.Rank, dc.Suit, dc.String(), dc.Rank.String(), dc.Suit.String(), Gameplayer.hand)
	DrawRank = dc.Rank.String()
	DrawSuit = dc.Suit.String()
	DrawCard = DrawRank + DrawSuit + ".png"
	/// make 52 image + deckback as separate the dislay
	// this how to refresh image
	//mycardimage.Resource = getcards.NewEmbeddedResource(mycard)
	Mycardimage.Resource = getcards.NewEmbeddedResource(DrawCard)
	//mycardimage := canvas.NewImageFromResource(getcards.NewEmbeddedResource(DrawCard))
	Mycardimage.SetMinSize(fyne.NewSize(100, 100))
	Mycardimage.FillMode = canvas.ImageFillContain
	Mycardimage.Refresh()
	Mycardimage.Show()

}

var GameBoard fyne.Container
var LastAction = ""
var MyDeck = NewDeck()
var Gameplayer = Player{}
var NPCplayer = Player{}
var DrawCard string
var DrawRank string
var DrawSuit string
var Playerkeep *widget.Button
var Playerdiscard *widget.Button
var Playerblindnil *widget.Button
var Playerregularnil *widget.Button
var Playerbid *widget.Button
var NPCkeep *widget.Button
var NPCdiscard *widget.Button
var NPCblindnil *widget.Button
var NPCregularnil *widget.Button
var NPCbid *widget.Button
var NPCbidbar *fyne.Container
var Playerbidbar *fyne.Container
var DeckCard *fyne.Container

// create and instantiate all gui elements
// update gui as needed
func setupgui() {
	labelcolor := tc.RGBA{253, 118, 87, 255}
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
	PlayerBid = canvas.NewText("0", tc.White)
	PlayerBid.TextSize = 32
	PlayerBags = canvas.NewText("0", tc.White)
	PlayerBags.TextSize = 32
	PlayerScore = canvas.NewText("0", tc.White)
	PlayerScore.TextSize = 32
	PlayerTricks = canvas.NewText("0", tc.White)
	PlayerTricks.TextSize = 32
	PlayerScoreBar = container.New(layout.NewGridLayoutWithColumns(9),
		playerbidname,
		playerbidlabel,
		PlayerBid,
		playerbidbagslabel,
		PlayerBags,
		playerbidscorelabel,
		PlayerScore,
		playerbidtrickslabel,
		PlayerTricks,
	)
	// npc
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
	NPCBid = canvas.NewText("0", tc.White)
	NPCBid.TextSize = 32
	NPCBags = canvas.NewText("0", tc.White)
	NPCBags.TextSize = 32
	NPCScore = canvas.NewText("0", tc.White)
	NPCScore.TextSize = 32
	NPCTricks = canvas.NewText("0", tc.White)
	NPCTricks.TextSize = 32
	NPCScoreBar = container.New(layout.NewGridLayoutWithColumns(9),
		npcbidname,
		npcbidlabel,
		NPCBid,
		npcbidbagslabel,
		NPCBags,
		npcbidscorelabel,
		NPCScore,
		npcbidtrickslabel,
		NPCTricks,
	)

	NPCkeep = widget.NewButton("Keep", func() {
		c := cropImage(DrawCard)
		c.FillMode = canvas.ImageFill(canvas.ImageScalePixels)
		c.SetMinSize(fyne.NewSize(100, 100))
		NPCCards.Add(c)

	})
	NPCdiscard = widget.NewButton("Discard", func() {

	})
	NPCblindnil = widget.NewButton("Blind Nil", func() {

	})
	NPCregularnil = widget.NewButton("Nil", func() {

	})
	NPCbid := widget.NewButton("Bid", func() {

	})
	NPCbidbar = container.NewGridWithColumns(5)
	NPCbidbar.Add(NPCkeep)
	NPCbidbar.Add(NPCdiscard)
	NPCbidbar.Add(NPCblindnil)
	NPCbidbar.Add(NPCregularnil)
	NPCbidbar.Add(NPCbid)

	Playerkeep = widget.NewButton("Keep", func() {
		LastAction = "KEEP"
		c := cropImage(DrawCard)
		c.FillMode = canvas.ImageFill(canvas.ImageScalePixels)
		c.SetMinSize(fyne.NewSize(100, 100))
		PlayerCards.Add(c)
		Turn(Buttonkeep)

	})
	Playerdiscard = widget.NewButton("Discard", func() {
		LastAction = "DISCARD"

		Turn(Buttondiscard)
	})
	Playerblindnil = widget.NewButton("Blind Nil", func() {

	})
	Playerregularnil = widget.NewButton("Nil", func() {

	})
	Playerbid = widget.NewButton("Bid", func() {

	})
	Playerbid.Disable()
	Playerbidbar = container.NewGridWithColumns(5)
	Playerbidbar.Add(Playerkeep)
	Playerbidbar.Add(Playerdiscard)
	Playerbidbar.Add(Playerblindnil)
	Playerbidbar.Add(Playerregularnil)
	Playerbidbar.Add(Playerbid)

	Playerregularnil.Disable()
	Playerbid.Disable()

	PlayerCards = container.NewGridWithColumns(13)
	NPCCards = container.NewCenter()
	DeckCard = container.NewCenter()
	CardsLeft = canvas.NewText(strconv.Itoa(len(MyDeck)), tc.White)
	CardsLeft.TextSize = 32
	DrawnCard := container.NewCenter()
	DrawnCard.Add(&Mycardimage)
	GameBoard = *container.NewVBox()
	GameBoard.Add(NPCScoreBar)
	GameBoard.Add(NPCCards)
	GameBoard.Add(NPCbidbar)
	GameBoard.Add(NPCCards)
	GameBoard.Add(DeckCard)
	GameBoard.Add(CardsLeft)
	GameBoard.Add(DrawnCard)
	GameBoard.Add(PlayerCards)

	GameBoard.Add(Playerbidbar)
	GameBoard.Add(PlayerScoreBar)
}
func deal() {
	runtime.GC()
	runtime.ReadMemStats(&memoryStats)

	config.FyneMainWin.SetTitle(config.GetLangs("title") + " " + strconv.FormatUint(memoryStats.Alloc/1024/1024, 10) + " Mib")

	//card := getcards.NewEmbeddedResource(config.DeckBack)
	//cropImage(strings.ToLower(config.DeckBack) + "_back.png")
	//deckbackimage := canvas.NewImageFromResource(getcards.NewEmbeddedResource("AS.png"))
	//deckbackimage := cropImage(strings.ToLower(config.DeckBack) + "_back.png")
	deckbackimage := canvas.NewImageFromResource(getcards.NewEmbeddedResource(strings.ToLower(config.DeckBack) + "_back.png"))
	deckbackimage.SetMinSize(fyne.NewSize(100, 100))
	deckbackimage.FillMode = canvas.ImageFillContain

	//PlayerCards.Add(deckbackimage)
	//NPCCards.Add(deckbackimage)

	MyDeck = NewDeck()
	MyDeck.Shuffle()

	Gameplayer = Player{deck: &MyDeck}
	NPCplayer = Player{deck: &MyDeck}
	log.Println(Gameplayer.hand, NPCplayer.hand)
	/* 	dc, dcerr := d.Draw()
	   	if dcerr != nil {
	   		log.Println("draw card ", dcerr)
	   	}
	   	log.Println("draw card ", dc.Rank, dc.Suit, dc.String(), dc.Rank.String(), dc.Suit.String(), gplayer.hand)
	   	DrawRank = dc.Rank.String()
	   	DrawSuit = dc.Suit.String()
	   	DrawCard = DrawRank + DrawSuit + ".png"
	   	/// make 52 image + deckback as separate the dislay
	   	// this how to refresh image
	   	//mycardimage.Resource = getcards.NewEmbeddedResource(mycard)
	   	mycardimage := canvas.NewImageFromResource(getcards.NewEmbeddedResource(DrawCard))
	   	mycardimage.SetMinSize(fyne.NewSize(100, 100))
	   	mycardimage.FillMode = canvas.ImageFillContain
	   	mycardimage.Refresh()
	   	mycardimage.Show()

	   	PlayerCards.Add(mycardimage) */
	//NPCCards.Add(deckbackimage)
	HandleCard()
	// NEW LAYOUT

	DeckCard.Add(deckbackimage)
	DrawnCard := container.NewCenter()
	DrawnCard.Add(&Mycardimage)
	/* 	GameBoard = *container.NewVBox()
	   	GameBoard.Add(NPCScoreBar)
	   	GameBoard.Add(NPCbidbar)
	   	GameBoard.Add(NPCCards)
	   	GameBoard.Add(DeckCard)
	   	GameBoard.Add(CardsLeft)
	   	GameBoard.Add(DrawnCard)
	   	GameBoard.Add(PlayerCards)

	   	GameBoard.Add(Playerbidbar)
	   	GameBoard.Add(PlayerScoreBar) */
	config.FyneMainWin.SetContent(&GameBoard)

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
	setupgui()
	runtime.GC()
	runtime.ReadMemStats(&memoryStats)
	config.FyneApp.SetIcon(MyLogo)
	config.FyneMainWin.SetTitle(config.GetLangs("title") + " " + strconv.FormatUint(memoryStats.Alloc/1024/1024, 10) + " Mib")
	config.FyneMainWin.Resize(fyne.NewSize(640, 480))
	splash()
	config.FyneMainWin.ShowAndRun()
}
