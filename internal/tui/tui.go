// Copyright 2017 Zack Guo <zack.y.guo@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT license that can
// be found in the LICENSE file.

//go:build ignore
// +build ignore

package main

import (
	"fmt"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"net/http"

	pokego "github.com/JoshGuarino/PokeGo/pkg"
	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

func main() {
	client := pokego.NewClient()
	s, err := client.Pokemon.GetPokemon("379")
	if err != nil {
		log.Fatal(err)
	}

	url := s.Sprites.FrontDefault
	fmt.Println(url)
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("failed to fetch image: %v", err)
	}
	image, _, err := image.Decode(resp.Body)
	if err != nil {
		log.Fatalf("failed to decode fetched image: %v", err)
	}
	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer ui.Close()

	img := widgets.NewImage(nil)
	img.SetRect(0, 0, 72, 36)
	render := func() {
		img.Image = image
		ui.Render(img)
	}
	render()

	uiEvents := ui.PollEvents()
	for {
		e := <-uiEvents
		switch e.ID {
		case "q", "<C-c>":
			return
		case "<Up>", "k":
			img.MonochromeThreshold += 20
			render()
		case "<Down>", "j":
			img.MonochromeThreshold += 20
			render()
		}
		render()
	}
}
