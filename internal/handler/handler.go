package handler

import (
	"fmt"
	"log"
	"os"

	spriteDownload "pokeinfotui/internal/download"
	"pokeinfotui/internal/trim_image"

	//drawImg "pokeinfotui/internal/image_render"
	pokego "github.com/JoshGuarino/PokeGo/pkg"
)

func Handler(input string) string {
	client := pokego.NewClient()
	s, err := client.Pokemon.GetPokemon(input)
	if err != nil {
		log.Fatal(err)
	}
	url := s.Sprites.FrontDefault
	filename := s.Name + ".png"
	spriteDownload.DownloadPrep(url, filename)
	filenameTrimmed := trim_image.TrimImage(filename)
	fmt.Println(filenameTrimmed)
	return filenameTrimmed
}

func Remove(filename string, filenameTrimmed string) {
	os.Remove(filenameTrimmed)
	os.Remove(filename)
}

func AbilityList() [100]string {
	var stringAbilityList [100]string
	return stringAbilityList
	//needs inout and is still being handled in tui.go
}

func MoveList() [100]string {
	var stringAbilityList [100]string
	return stringAbilityList
	//needs inout and is still being handled in tui.go
}

func stringTypes() string {
	var stringTypes string
	return stringTypes
	//needs inout and is still being handled in tui.go
}
