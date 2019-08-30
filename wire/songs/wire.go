//+build wireinject

package songs

import "github.com/google/wire"

func InitializeMovie(line string) Movie {
	wire.Build(NewSong, NewAlbum, NewMovie)
	return Movie{}
}
