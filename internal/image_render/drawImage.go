package render

import (
	"fmt"
	"log"

	imgfetch "github.com/alan-ar1/imgfetch/pkg/imgfetch"
)

func Draw(filepath string, rows int, columns int) {
	size := imgfetch.ImageTermSize{Rows: rows, Columns: columns}
	imgSeq, err := imgfetch.GetImageSeq(filepath, size)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(imgSeq)
}
