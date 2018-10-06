package id3

import (
	"fmt"
	"os"

	id3 "github.com/mikkyang/id3-go"
)

//set tags elemets to tag in file named fname.
func SetTagElements(fname string, title string, artist string, album string) {
	fmt.Printf("title %s add to %v\n", title, fname)
	mp3file, err := id3.Open(fname)
	if err != nil {
		println(err)
		os.Exit(-1)
	}
	if title != "" {
		mp3file.SetTitle(title)
	}
	if artist != "" {
		mp3file.SetArtist(artist)
	}
	if album != "" {
		mp3file.SetAlbum(album)
	}
	defer mp3file.Close()
}
