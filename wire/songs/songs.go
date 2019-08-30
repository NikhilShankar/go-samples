package songs

import "fmt"

type Lyrics string


type Song struct {
	Lyrics
}

type Album struct {
	Songs []Song
}

type Movie struct {
	Album
}

func (a *Album) GetSongs() {
	for _,s := range a.Songs {
		fmt.Println(s.Lyrics)
	}
}

func (a *Album) AddSong(s Song) {
	a.Songs = append(a.Songs, s)
}

func (m *Movie) GetSongs() {
	m.Album.GetSongs()
}

func NewSong(line string) Song{
	return Song{
		Lyrics(line),
	}
}

func NewAlbum(song Song) Album {
	s := make([]Song,0)
	s = append(s, song)
	return Album{
		Songs: s,
	}
}

func NewMovie(album Album) Movie {
	return Movie{
		Album: album,
	}
}
