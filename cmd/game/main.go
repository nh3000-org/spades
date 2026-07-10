package main

import (
	"bytes"
	"image"
	"sort"
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

var PS Game
var NPC Game

type Game struct {
	Stats         config.PS
	Left          *canvas.Text
	Bid           *canvas.Text
	Bags          *canvas.Text
	Score         *canvas.Text
	Tricks        *canvas.Text
	ScoreBar      *fyne.Container
	Cards         *fyne.Container
	CardList      []string
	Bidlabel      *canvas.Text
	Bidbagslabel  *canvas.Text
	Bidscorelabel *canvas.Text

	Bidtrickslabel *canvas.Text

	ButtonKeep       *widget.Button
	ButtonDiscard    *widget.Button
	ButtonBlindnil   *widget.Button
	ButtonSTM        *widget.Button
	ButtonRegularnil *widget.Button
	Bidname          *canvas.Text
	ButtonBid        *widget.Button
	BidValue         *widget.Entry
	ButtonBidAll     *widget.Button
	Bidbar           *fyne.Container
}

var (
	PlayerGame       config.PS
	NPCGame          config.PS
	memoryStats      runtime.MemStats
	spadesheader     *canvas.Image
	CardsLeft        *canvas.Text
	PlayerTurn       = "PLAYER"
	TurnCount        int // 1 is first turn // 2 is second trun
	Buttonkeep       = 1
	Buttondiscard    = 2
	Buttonblindnil   = 3
	Buttonregularnil = 4
	Buttonstm        = 5
	Buttonbid        = 6
	done             = false
	Mycardimage      canvas.Image
	left             int

	GameBoard    fyne.Container
	ActionArea   *fyne.Container
	LastAction   = ""
	MyDeck       = NewDeck()
	Gameplayer   = Player{}
	NPCplayer    = Player{}
	DrawCard     string
	DrawCardSort string
	DrawRank     string
	DrawSuit     string
	SortOrder    string
	DeckCard     *fyne.Container
	Playarea     = container.NewCenter()
	ActionBid    *canvas.Text
	Sizew        float32
	Sizeh        float32
)

func cropImage(c string) *canvas.Image {
	//log.Println("cropImage ", c)
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
	}).SubImage(image.Rect(0, 0, 100, 300))
	//log.Println(my_sub_image.Bounds())
	var b []byte
	err1 := png.Encode(bytes.NewBuffer(b), my_sub_image)
	if err1 != nil {
		log.Println("inithand", err1)

	}
	cardimg := canvas.NewImageFromImage(my_sub_image)

	cardimg.SetMinSize(fyne.NewSize(100*Sizew, 300*Sizeh))
	return cardimg
}

func Turn(button int) {

	TurnCount++
	log.Println("Turn count", TurnCount, "playerturn", PlayerTurn)
	done = false
	if TurnCount > 2 {
		if PlayerTurn == "PLAYER" {
			PlayerTurn = "NPC"
			done = true
			log.Println("in player")

		}
		if PlayerTurn == "NPC" && !done {
			PlayerTurn = "PLAYER"
			log.Println("in npc")
		}
		TurnCount = 1
	}

	if TurnCount == 1 {
		// player
		if PlayerTurn == "PLAYER" {
			PS.ButtonKeep.Enable()
			PS.ButtonDiscard.Enable()
			if button == Buttonkeep {
				PS.ButtonKeep.Disable()
				PS.ButtonDiscard.Enable()
				PS.ButtonBlindnil.Disable()
				PS.ButtonRegularnil.Disable()
				PS.ButtonBid.Disable()
				PS.BidValue.Disable()
			}
			if button == Buttondiscard {
				PS.ButtonKeep.Enable()
				PS.ButtonDiscard.Disable()
				PS.ButtonBlindnil.Disable()
				PS.ButtonRegularnil.Disable()
				PS.ButtonBid.Disable()
				PS.BidValue.Disable()

			}
			HandleCard()
		}
		// npc

	}
	if TurnCount == 2 {
		if PlayerTurn == "PLAYER" {
			if button == Buttonkeep {
				/* 				Playerkeep.Disable()
				   				Playerdiscard.Enable()
				   				Playerblindnil.Disable()
				   				Playerregularnil.Disable()
				   				Playerbid.Disable() */
			}
			if button == Buttondiscard {
				/* 				Playerkeep.Enable()
				   				Playerdiscard.Disable()
				   				Playerblindnil.Disable()
				   				Playerregularnil.Disable()
				   				Playerbid.Disable() */
			}
			HandleCard()

			NPC.ButtonKeep.Enable()
			NPC.ButtonDiscard.Disable()
			NPC.ButtonBlindnil.Disable()
			NPC.ButtonRegularnil.Disable()
			//NPCBid.Disable()
			NPC.ButtonKeep.Tapped(nil)
			HandleCard()
			NPC.ButtonDiscard.Tapped(nil)

			HandleCard()
			TurnCount = 0
			PlayerTurn = "PLAYER"
			PS.ButtonKeep.Enable()
			PS.ButtonDiscard.Enable()
		}

	}

}

func HandleCard() {
	left = len(MyDeck)
	if left == 0 {
		PS.ButtonKeep.Hide()
		PS.ButtonDiscard.Hide()
		PS.ButtonBlindnil.Hide()
		PS.ButtonSTM.Hide()
		PS.ButtonRegularnil.Enable()
		PS.ButtonBid.Enable()
		PS.BidValue.Enable()
		NPC.ButtonKeep.Hide()
		NPC.ButtonDiscard.Hide()
		NPC.ButtonBlindnil.Hide()
		NPC.ButtonSTM.Hide()
		NPC.ButtonRegularnil.Enable()
		NPC.ButtonBid.Enable()
		NPC.BidValue.Enable()
		Mycardimage.Hide()

		bid()
		return
	}
	dc, dcerr := MyDeck.Draw()
	if dcerr != nil {
		log.Println("draw card error", dcerr)
	}

	CardsLeft.Text = strconv.Itoa(len(MyDeck))

	//log.Println("draw card ", dc.Rank, dc.Suit)
	DrawRank = dc.Rank.String()
	DrawSuit = dc.Suit.String()
	DrawCard = DrawRank + DrawSuit + ".png"
	SortOrder = ""
	switch DrawSuit {
	case "S":
		SortOrder = "01"
	case "H":
		SortOrder = "02"
	case "C":
		SortOrder = "03"
	case "D":
		SortOrder = "04"
	}
	switch DrawRank {
	case "2":
		SortOrder = SortOrder + "01"
	case "3":
		SortOrder = SortOrder + "02"
	case "4":
		SortOrder = SortOrder + "03"
	case "5":
		SortOrder = SortOrder + "04"
	case "6":
		SortOrder = SortOrder + "05"
	case "7":
		SortOrder = SortOrder + "06"
	case "8":
		SortOrder = SortOrder + "07"
	case "9":
		SortOrder = SortOrder + "08"
	case "10":
		SortOrder = SortOrder + "09"
	case "J":
		SortOrder = SortOrder + "10"
	case "Q":
		SortOrder = SortOrder + "11"
	case "K":
		SortOrder = SortOrder + "12"
	case "A":
		SortOrder = SortOrder + "13"
	}
	//get ranking of ranks
	DrawCardSort = SortOrder + ":" + DrawRank + DrawSuit + ".png"

	log.Println("draw card sort", DrawCardSort, PlayerTurn)

	Mycardimage.Resource = getcards.NewEmbeddedResource(DrawCard)

	Mycardimage.SetMinSize(fyne.NewSize(100*Sizew, 200*Sizeh))
	Mycardimage.FillMode = canvas.ImageFillContain
	Mycardimage.Refresh()
	Mycardimage.Show()

}

// create and instantiate all gui elements
// update gui as needed

func setupgui() {
	labelcolor := tc.RGBA{253, 118, 87, 255}
	PS = Game{}
	PS.Bidlabel = canvas.NewText("Bid:", labelcolor)
	PS.Bidlabel.TextSize = 32
	PS.Bidbagslabel = canvas.NewText("Bags:", labelcolor)
	PS.Bidbagslabel.TextSize = 32
	PS.Bidscorelabel = canvas.NewText("Score:", labelcolor)
	PS.Bidscorelabel.TextSize = 32
	PS.Bidtrickslabel = canvas.NewText("Tricks:", labelcolor)
	PS.Bidtrickslabel.TextSize = 32
	PS.Bidname = canvas.NewText(config.PlayerName, tc.White)
	PS.Bidname.TextSize = 32
	PS.Bid = canvas.NewText("0", tc.White)
	PS.Bid.TextSize = 32
	PS.Bags = canvas.NewText("0", tc.White)
	PS.Bags.TextSize = 32
	PS.Score = canvas.NewText("0", tc.White)
	PS.Score.TextSize = 32
	PS.Tricks = canvas.NewText("0", tc.White)
	PS.Tricks.TextSize = 32

	PS.Cards = container.NewHBox()

	NPC = Game{}
	NPC.Bidlabel = canvas.NewText("Bid:", labelcolor)
	NPC.Bidlabel.TextSize = 32
	NPC.Bidbagslabel = canvas.NewText("Bags:", labelcolor)
	NPC.Bidbagslabel.TextSize = 32
	NPC.Bidscorelabel = canvas.NewText("Score:", labelcolor)
	NPC.Bidscorelabel.TextSize = 32
	NPC.Bidtrickslabel = canvas.NewText("Tricks:", labelcolor)
	NPC.Bidtrickslabel.TextSize = 32
	NPC.Bidname = canvas.NewText("Computer", tc.White)
	NPC.Bidname.TextSize = 32
	NPC.Bidlabel = canvas.NewText("Bid:", labelcolor)
	NPC.Bidlabel.TextSize = 32
	NPC.Bid = canvas.NewText("0", tc.White)
	NPC.Bid.TextSize = 32
	NPC.Bags = canvas.NewText("0", tc.White)
	NPC.Bags.TextSize = 32
	NPC.Score = canvas.NewText("0", tc.White)
	NPC.Score.TextSize = 32
	NPC.Tricks = canvas.NewText("0", tc.White)
	NPC.Tricks.TextSize = 32

	NPC.Cards = container.NewHBox()
	NPC.ButtonKeep = widget.NewButton("Keep", func() {

		NPC.CardList = append(NPC.CardList, DrawCardSort)
		sort.Strings(NPC.CardList)
		NPC.Cards.RemoveAll()
		for _, card := range NPC.CardList {

			s := strings.Split(card, ":")
			c := cropImage(s[1])
			c.FillMode = canvas.ImageFill(canvas.ImageScalePixels)

			c.SetMinSize(fyne.NewSize(50*Sizew, 100*Sizeh))
			NPC.Cards.Add(c)
		}
		NPC.ButtonBlindnil.Disable()
		NPC.ButtonSTM.Disable()
		NPC.Cards.Refresh()

	})
	NPC.ButtonDiscard = widget.NewButton("Discard", func() {
		NPC.ButtonBlindnil.Disable()
		NPC.ButtonSTM.Disable()
		Turn(Buttondiscard)
	})
	NPC.ButtonBlindnil = widget.NewButton("Blind Nil", func() {
		NPC.BidValue.Text = "BLINDNIL"
		NPC.BidValue.Disable()
		NPC.ButtonBlindnil.Disable()
		NPC.ButtonSTM.Disable()
		NPC.ButtonRegularnil.Disable()
		NPC.ButtonBid.Disable()
		Turn(Buttonblindnil)

	})
	NPC.ButtonSTM = widget.NewButton("Shoot The Moon", func() {
		NPC.BidValue.Text = "STM"

		NPC.ButtonBlindnil.Disable()
		NPC.ButtonSTM.Disable()
		NPC.BidValue.Disable()
		Turn(Buttonstm)
	})
	NPC.ButtonRegularnil = widget.NewButton("Nil", func() {
		NPC.ButtonBlindnil.Disable()
		NPC.ButtonSTM.Disable()
		NPC.ButtonRegularnil.Disable()
		NPC.BidValue.Text = "NIL"
		NPC.BidValue.Disable()
		Turn(Buttonregularnil)
	})
	NPC.ButtonBid = widget.NewButton("Bid", func() {
		NPC.ButtonBlindnil.Disable()
		NPC.ButtonSTM.Disable()
		NPC.ButtonRegularnil.Disable()
		NPC.BidValue.Disable()
		Turn(Buttonbid)
	})
	//NPC.BidValue = widget.NewSelect([]string{"00", "01", "02", "03", "04", "05", "06", "07", "08", "09", "10", "11", "12", "13"}, func(string) {})

	NPC.BidValue = widget.NewEntry()
	NPC.Bidbar = container.NewGridWithColumns(7)
	NPC.Bidbar.Add(NPC.ButtonKeep)
	NPC.Bidbar.Add(NPC.ButtonDiscard)
	NPC.Bidbar.Add(NPC.ButtonBlindnil)
	NPC.Bidbar.Add(NPC.ButtonSTM)
	NPC.Bidbar.Add(NPC.ButtonRegularnil)
	NPC.Bidbar.Add(NPC.ButtonBid)
	NPC.Bidbar.Add(NPC.BidValue)
	NPC.ButtonKeep.Disable()
	NPC.ButtonDiscard.Disable()
	NPC.ButtonBlindnil.Disable()
	NPC.ButtonSTM.Disable()
	NPC.ButtonRegularnil.Disable()
	NPC.ButtonBid.Disable()
	NPC.BidValue.Disable()
	NPC.ScoreBar = container.New(layout.NewGridLayoutWithColumns(9),
		NPC.Bidname,
		NPC.Bidlabel,
		NPC.Bid,
		NPC.Bidbagslabel,
		NPC.Bags,
		NPC.Bidscorelabel,
		NPC.Score,
		NPC.Bidtrickslabel,
		NPC.Tricks,
	)

	PS = Game{}
	PS.Bidlabel = canvas.NewText("Bid:", labelcolor)
	PS.Bidlabel.TextSize = 32
	PS.Bidbagslabel = canvas.NewText("Bags:", labelcolor)
	PS.Bidbagslabel.TextSize = 32
	PS.Bidscorelabel = canvas.NewText("Score:", labelcolor)
	PS.Bidscorelabel.TextSize = 32
	PS.Bidtrickslabel = canvas.NewText("Tricks:", labelcolor)
	PS.Bidtrickslabel.TextSize = 32
	PS.Bidname = canvas.NewText(config.PlayerName, tc.White)
	PS.Bidname.TextSize = 32
	PS.Bidlabel = canvas.NewText("Bid:", labelcolor)
	PS.Bidlabel.TextSize = 32
	PS.Bid = canvas.NewText("0", tc.White)
	PS.Bid.TextSize = 32
	PS.Bags = canvas.NewText("0", tc.White)
	PS.Bags.TextSize = 32
	PS.Score = canvas.NewText("0", tc.White)
	PS.Score.TextSize = 32
	PS.Tricks = canvas.NewText("0", tc.White)
	PS.Tricks.TextSize = 32
	PS.ScoreBar = container.New(layout.NewGridLayoutWithColumns(9),
		PS.Bidname,
		PS.Bidlabel,
		PS.Bid,
		PS.Bidbagslabel,
		PS.Bags,
		PS.Bidscorelabel,
		PS.Score,
		PS.Bidtrickslabel,
		PS.Tricks,
	)
	PS.Cards = container.NewHBox()
	PS.ButtonKeep = widget.NewButton("Keep", func() {
		LastAction = "KEEP"

		PS.CardList = append(PS.CardList, DrawCardSort)
		log.Println(DrawCard)
		sort.Strings(PS.CardList)
		PS.Cards.RemoveAll()
		for _, card := range PS.CardList {
			s := strings.Split(card, ":")
			c := cropImage(s[1])
			c.FillMode = canvas.ImageFill(canvas.ImageScalePixels)

			c.SetMinSize(fyne.NewSize(50*Sizew, 100*Sizeh))
			PS.Cards.Add(c)
		}
		PS.Cards.Refresh()
		PS.ButtonBlindnil.Disable()
		PS.ButtonSTM.Disable()
		//log.Println("playerkeep", PS.Cards)
		Turn(Buttonkeep)

	})
	PS.ButtonDiscard = widget.NewButton("Discard", func() {
		LastAction = "DISCARD"
		PS.ButtonBlindnil.Disable()
		PS.ButtonSTM.Disable()
		Turn(Buttondiscard)
	})
	PS.ButtonBlindnil = widget.NewButton("Blind Nil", func() {
		PS.ButtonBlindnil.Disable()
		PS.ButtonSTM.Disable()
		PS.BidValue.Text = "BLINDNIL"
		PS.BidValue.Disable()
		PS.ButtonBid.Disable()
		PS.ButtonRegularnil.Disable()
		Turn(Buttonblindnil)

	})
	PS.ButtonSTM = widget.NewButton("Shoot The Moon", func() {
		PS.ButtonBlindnil.Disable()
		PS.ButtonSTM.Disable()
		PS.BidValue.Text = "STM"
		PS.BidValue.Disable()
		PS.BidValue.Disable()
		PS.ButtonBid.Disable()
		PS.ButtonRegularnil.Disable()
		Turn(Buttonstm)
	})
	PS.ButtonRegularnil = widget.NewButton("Nil", func() {
		PS.ButtonBlindnil.Disable()
		PS.ButtonSTM.Disable()
		PS.ButtonRegularnil.Disable()
		PS.BidValue.Text = "NIL"
		PS.BidValue.Disable()
		Turn(Buttonregularnil)
	})
	PS.ButtonBid = widget.NewButton("Bid", func() {
		PS.ButtonBlindnil.Disable()
		PS.ButtonSTM.Disable()
		PS.ButtonRegularnil.Disable()
		PS.BidValue.Disable()
		Turn(Buttonbid)
	})
	PS.ButtonBid.Disable()
	//PS.BidValue = widget.NewSelect([]string{"00", "01", "02", "03", "04", "05", "06", "07", "08", "09", "10", "11", "12", "13"}, func(string) {})
	PS.BidValue = widget.NewEntry()
	PS.Bidbar = container.NewGridWithColumns(7)
	PS.Bidbar.Add(PS.ButtonKeep)
	PS.Bidbar.Add(PS.ButtonDiscard)
	PS.Bidbar.Add(PS.ButtonBlindnil)
	PS.Bidbar.Add(PS.ButtonSTM)
	PS.Bidbar.Add(PS.ButtonRegularnil)
	PS.Bidbar.Add(PS.ButtonBid)
	PS.Bidbar.Add(PS.BidValue)

	PS.ButtonRegularnil.Disable()
	PS.ButtonBid.Disable()

	//CardsLayout := layout.NewCustomPaddedLayout(1, 1, 1, 1)
	//PlayerCards = container.New(CardsLayout)

	//PlayerCards = container.NewGridWithColumns(13)
	PS.Cards = container.NewHBox()
	log.Println("setup player", PS.Cards.MinSize())

	//NPCCards = container.NewGridWithColumns(13)
	//NPCCards = container.NewHBox()

	//NPCCardsLayout := layout.NewCustomPaddedLayout(1, 1, 1, 1)
	//NPCCards = container.New(NPCCardsLayout)
	DeckBackImage = canvas.NewImageFromResource(getcards.NewEmbeddedResource("green_back.png"))

	DeckBackImage.SetMinSize(fyne.NewSize(50*Sizew, 50*Sizeh))
	DeckBackImage.FillMode = canvas.ImageFillContain
	ActionArea = container.NewGridWithColumns(2)
	ActionArea.Add(DeckBackImage)
	CardsLeft = canvas.NewText(strconv.Itoa(len(MyDeck)), tc.White)
	CardsLeft.TextSize = 32
	ActionArea.Add(CardsLeft)
	//DeckCard = container.NewCenter()

	Playarea.Add(&Mycardimage)

	GameBoard = *container.NewVBox()
	GameBoard.Add(NPC.ScoreBar)
	GameBoard.Add(NPC.Bidbar)
	GameBoard.Add(NPC.Cards)

	GameBoard.Add(Playarea)

	GameBoard.Add(PS.Cards)

	GameBoard.Add(PS.Bidbar)
	GameBoard.Add(PS.ScoreBar)
}
func play() {
	GameBoard.RemoveAll()
	GameBoard.Add(NPC.ScoreBar)
	GameBoard.Add(NPC.Bidbar)
	GameBoard.Add(NPC.Cards)

	ActionArea = container.NewGridWithColumns(2)
	GameBoard.Add(ActionArea)

	GameBoard.Add(ActionArea)

	GameBoard.Add(Playarea)
	GameBoard.Add(PS.Cards)

	GameBoard.Add(PS.Bidbar)
	GameBoard.Add(PS.ScoreBar)
}
func bid() {
	GameBoard.RemoveAll()
	GameBoard.Add(NPC.ScoreBar)
	if NPC.BidValue != nil {
		GameBoard.Add(NPC.Bidbar)
	}
	GameBoard.Add(NPC.Cards)
	ActionArea.RemoveAll()
	ActionArea = container.NewGridWithColumns(2)
	ActionArea.Add(DeckBackImage)
	turn := "Please Bid " + config.PlayerName
	if PlayerTurn == "NPC" {
		turn = "Computer Bidding"

	}
	ActionBid = canvas.NewText(turn, tc.White)
	ActionBid.TextSize = 64
	ActionArea.Add(ActionBid)
	ActionArea.Refresh()

	GameBoard.Add(ActionArea)

	GameBoard.Add(Playarea)
	GameBoard.Add(PS.Cards)
	if PS.BidValue != nil {
		GameBoard.Add(PS.Bidbar)
	}
	GameBoard.Add(PS.ScoreBar)
}

var DeckBackImage *canvas.Image

func deal() {
	runtime.GC()
	runtime.ReadMemStats(&memoryStats)

	config.FyneMainWin.SetTitle(config.GetLangs("title") + " " + strconv.FormatUint(memoryStats.Alloc/1024/1024, 10) + " Mib")

	MyDeck = NewDeck()
	MyDeck.Shuffle()

	Gameplayer = Player{deck: &MyDeck}
	NPCplayer = Player{deck: &MyDeck}
	log.Println(Gameplayer.hand, NPCplayer.hand)

	HandleCard()

	Playarea.Add(&Mycardimage)

	config.FyneMainWin.SetContent(&GameBoard)

}

func splash() {

	runtime.GC()
	runtime.ReadMemStats(&memoryStats)
	config.FyneMainWin.SetTitle(config.GetLangs("title") + " " + strconv.FormatUint(memoryStats.Alloc/1024/1024, 10) + " Mib")
	spadesheader = canvas.NewImageFromResource(getcards.NewEmbeddedResource("honors_spade-14.png"))
	spadesheader.FillMode = canvas.ImageFillContain
	spadesheader.SetMinSize(fyne.NewSize(100*Sizew, 200*Sizeh))

	rules := widget.NewMultiLineEntry()
	rules.SetText(config.GetLangs("rules"))

	config.Fyne_scale = config.FyneApp.Preferences().StringWithFallback("Fyne_scale", "1.5")
	fslabel := widget.NewLabel(config.GetLangs("scale"))
	fs := widget.NewEntry()
	fs.SetText(config.Fyne_scale)
	fs.SetPlaceHolder(config.GetLangs("scale"))

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
	DeckBackImage = canvas.NewImageFromResource(getcards.NewEmbeddedResource(strings.ToLower(config.DeckBack) + "_back.png"))
	DeckBackImage.SetMinSize(fyne.NewSize(50*Sizew, 50*Sizeh))
	DeckBackImage.FillMode = canvas.ImageFillContain
	deckback.Horizontal = false

	rightbox := container.NewVBox(
		widget.NewLabelWithStyle(config.GetLangs("preferences"), fyne.TextAlignCenter, fyne.TextStyle{Bold: true}),

		fslabel,
		fs,

		playerlabel,
		player,

		difficultylabel,
		difficulty,

		deckbacklabel,
		deckback,
	)

	next := widget.NewButton("Next", func() {
		fscheck, err := strconv.ParseFloat(fs.Text, 64)
		if err != nil {
			log.Println("Invalid scale value")
			fs.SetText("1.0")
		}
		if fscheck < 1.0 || fscheck > 2.0 {
			log.Println("Scale value must be between 0.5 and 2.0")
			fs.SetText("1.0")
		}

		config.FyneApp.Preferences().SetString("Fyne_scale", fs.Text)
		config.FyneApp.Preferences().SetString("Player", player.Text)
		config.PlayerName = player.Text
		config.FyneApp.Preferences().SetString("Difficulty", difficulty.Selected)
		config.Difficulty = difficulty.Selected
		config.FyneApp.Preferences().SetString("Deckback", deckback.Selected)

		config.DeckBack = deckback.Selected
		config.DealerPlayer = true
		PlayerGame = config.NewPlayer(player.Text)
		NPCGame = config.NewPlayer("NPC")
		w, _ := strconv.ParseFloat(config.Fyne_scale, 32)
		h, _ := strconv.ParseFloat(config.Fyne_scale, 32)

		Sizew = float32(w)
		Sizeh = float32(h)
		deal()

	})

	border := container.NewBorder(spadesheader, next, nil, rightbox, rules)
	config.FyneMainWin.SetContent(border)
	//config.FyneMainWin.Canvas().Refresh(border)
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

	config.FyneApp.Settings().SetTheme(&config.MyTheme{})
	MyLogo, iconerr := fyne.LoadResourceFromPath("icon.png")
	if iconerr != nil {
		log.Println("icon.png error ", iconerr.Error())
	}
	config.PlayerName = config.FyneApp.Preferences().StringWithFallback("Player", "Player1")
	setupgui()
	runtime.GC()
	runtime.ReadMemStats(&memoryStats)
	config.FyneApp.SetIcon(MyLogo)
	config.FyneMainWin.SetTitle(config.GetLangs("title") + " " + strconv.FormatUint(memoryStats.Alloc/1024/1024, 10) + " Mib")
	config.FyneMainWin.Resize(fyne.NewSize(1024, 800))
	splash()
	config.FyneMainWin.ShowAndRun()
}
