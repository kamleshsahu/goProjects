package _interface

import "awesomeProject/lld/bookMyShow/entity"

type ITheatreService interface {
	AddTheatre(cityId int, theatre entity.Theatre) (int, error)
	AddScreen(theatreId int, screen entity.Screen) (int, error)
	AddShow(theatreId int, show entity.Show) (int, error)
	GetShows(movieId int, cityId int) ([]entity.Show, error)
}

type IMovieService interface {
	AddMovie(movie entity.Movie) (int, error)
	GetMovies(filter map[string]bool) ([]entity.Movie, error)
}

// has theatre
// has movies
type IBookingService interface {
	GetTheatreByMovie(movieId string) ([]entity.Theatre, error)
	GetMovieByCity(city string) ([]entity.Movie, error)
	BookShow(showId int) (entity.Booking, error)
}
