package main

import "github.com/NikhilShankar/go-samples/wire/songs"

func main() {

	////Without wire
	//song := songs.NewSong("Wish you were here")
	//album := songs.NewAlbum()
	//album.AddSong(song)
	//movie := songs.NewMovie(album)
	//movie.GetSongs()

	//With Wire
	movie := songs.InitializeMovie("Wish you were here")
	movie.GetSongs()

}
