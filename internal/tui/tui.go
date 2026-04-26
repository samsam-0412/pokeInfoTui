package main

import (
	"encoding/json"
	"fmt"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"math/rand"
	"os"
	spriteDownload "pokeinfotui/internal/download"
	"pokeinfotui/internal/handler"
	"pokeinfotui/internal/trim_image"
	"regexp"
	"strconv"

	pokego "github.com/JoshGuarino/PokeGo/pkg"
	"github.com/JoshGuarino/PokeGo/pkg/models"
	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
	//terminal "golang.org/x/crypto/ssh/terminal"
)

func main() {
	var input string
	var err error
	var s *models.Pokemon
	client := pokego.NewClient()
	fmt.Print("Enter a Pokemon name or pokedexId or random: ")
	fmt.Scanln(&input)
	if input == "random" {
		input = strconv.Itoa(rand.Intn(1025-1) + 1)
	}
	s, err = client.Pokemon.GetPokemon(input)
	if err != nil {
		log.Fatal(err)
	}

	//ttyWidth, ttyHheight, err := terminal.GetSize(0)
	url := s.Sprites.FrontDefault
	filename := s.Name + ".png"
	spriteDownload.DownloadPrep(url, filename)
	filenameTrimmed := trim_image.TrimImage(filename)
	fmt.Println(filenameTrimmed)

	var stringTypes string
	TypesList := s.Types
	for index := range TypesList {
		re := regexp.MustCompile(`[^\d\p{Latin}]`)
		currentElement, err := json.Marshal(TypesList[index].Type.Name)
		if err != nil {
		}
		jsonStr := re.ReplaceAllString(string(currentElement), "")
		stringTypes += jsonStr + "\n"
	}
	var stringAbilityList [100]string
	abilityList := s.Abilities
	for index := range abilityList {

		re := regexp.MustCompile(`[^\d\p{Latin}]`)

		currentElement, err := json.Marshal(abilityList[index].Ability.Name)
		if err != nil {
		}
		jsonStr := re.ReplaceAllString(string(currentElement), "")
		stringAbilityList[index] = jsonStr

	}
	var stringMoveList [500]string
	moveList := s.Moves
	for index := range moveList {

		re := regexp.MustCompile(`[^\d\p{Latin}]`)

		currentElementMove, err := json.Marshal(moveList[index].Move.Name)
		if err != nil {
		}
		jsonStrMove := re.ReplaceAllString(string(currentElementMove), "")
		stringMoveList[index] = jsonStrMove
	}

	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer ui.Close()

	imgFile, err := os.Open(filenameTrimmed)
	if err != nil {
		log.Fatal(err)
	}
	defer imgFile.Close()

	image, _, err := image.Decode(imgFile)
	if err != nil {
		log.Println("Cannot decode file:", err)
	}
	img := widgets.NewImage(image)
	img.SetRect(0, 0, 40, 40)

	name := widgets.NewParagraph()
	name.Title = "Name"
	name.Text = s.Name
	name.SetRect(40, 0, 75, 3)

	height := widgets.NewParagraph()
	height.Title = "Height"
	height.Text = strconv.Itoa(s.Height) + " m"
	height.SetRect(40, 3, 75, 6)

	weight := widgets.NewParagraph()
	weight.Title = "Weight"
	weight.Text = strconv.Itoa(s.Weight) + " kg"
	weight.SetRect(40, 6, 75, 9)

	abilityListWidget := widgets.NewList()
	abilityListWidget.Title = "Abilities"
	abilityListWidget.Rows = stringAbilityList[:]
	abilityListWidget.TextStyle = ui.NewStyle(ui.ColorYellow)
	abilityListWidget.WrapText = false
	abilityListWidget.SetRect(40, 9, 75, 15)

	moveListWidget := widgets.NewList()
	moveListWidget.Title = "Moves"
	moveListWidget.Rows = stringMoveList[:]
	moveListWidget.TextStyle = ui.NewStyle(ui.ColorYellow)
	moveListWidget.WrapText = false
	moveListWidget.SetRect(40, 15, 75, 21)

	typesWidget := widgets.NewParagraph()
	typesWidget.Title = "Types"
	typesWidget.Text = stringTypes
	typesWidget.SetRect(40, 21, 75, 25)

	keybind := widgets.NewParagraph()
	keybind.Title = "Keybinds"
	keybind.SetRect(40, 29, 76, 40)

	scrollAbilities := widgets.NewParagraph()
	scrollAbilities.Title = "Abilities"
	scrollAbilities.Text = "Scroll Abilities: K and J or Down and Up (Arrows)"
	scrollAbilities.SetRect(41, 30, 75, 33)

	scrollMove := widgets.NewParagraph()
	scrollMove.Title = "Abilities"
	scrollMove.Text = "Scroll Moves: H and L or Left and Right (Arrows)"
	scrollMove.SetRect(41, 33, 75, 36)

	quit := widgets.NewParagraph()
	quit.Title = "Exit"
	quit.Text = "Press q to quit"
	quit.SetRect(41, 36, 75, 39)
	render := func() {
		img.Image = image
		ui.Render(img, name, height, weight, abilityListWidget, moveListWidget, typesWidget, keybind, quit, scrollAbilities, scrollMove)
	}
	render()
	previousKey := ""
	uiEvents := ui.PollEvents()
	for {
		e := <-uiEvents
		switch e.ID {
		case "q":
			handler.Remove(filename, filenameTrimmed)
			return
		case "j", "<Down>":
			abilityListWidget.ScrollDown()
		case "k", "<Up>":
			abilityListWidget.ScrollUp()
		case "h", "<Left>":
			moveListWidget.ScrollUp()
		case "l", "<Right>":
			moveListWidget.ScrollDown()
		}
		if previousKey == "g" {
			previousKey = ""
		} else {
			previousKey = e.ID
		}
		ui.Render(abilityListWidget, moveListWidget)
	}
}
