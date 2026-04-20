package handler

import (
	"log"
	"os"

	spriteDownload "pokeinfotui/internal/download"
	imgTrim "pokeinfotui/internal/trim_image"

	//drawImg "pokeinfotui/internal/image_render"
	pokego "github.com/JoshGuarino/PokeGo/pkg"
)

func Handler(input string) string {
	client := pokego.NewClient()
	s, err := client.Pokemon.GetPokemon("charmander")
	if err != nil {
		log.Fatal(err)
	}
	url := s.Sprites.FrontDefault
	filename := s.Name + ".png"
	spriteDownload.DownloadPrep(url, filename)
	filenameTrimmed := imgTrim.TrimImage(filename)
	return filenameTrimmed
}

// unused
func remove(filename string, filenameTrimmed string) {
	os.Remove(filenameTrimmed)
	os.Remove(filename)
}
