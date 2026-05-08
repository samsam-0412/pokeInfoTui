package spritedownload

import (
	"io"
	"net/http"
	"os"
)

func DownloadPrep(url string, pokemonName string) {
	err := DownloadFile(url, pokemonName)
	if err != nil {
		panic(err)
	}

}

// spriteDownload.go from https://gist.github.com/cnu/026744b1e86c6d9e22313d06cba4c2e9
// DownloadFile will download a url and store it in local filepath.
// It writes to the destination file as it downloads it, without
// loading the entire file into memory.
func DownloadFile(url string, filepath string) error {
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	return nil
}
