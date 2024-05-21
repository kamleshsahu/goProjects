package fileSystem

import (
	"awesomeProject/fileSystem/service"
	"fmt"
)

func Run() {

	movieDir := service.NewDirectory("movie")

	movie1, err := service.NewFile("bahubali.mov")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	movie2, err := service.NewFile("bahubali2.mov")

	bmovies := service.NewDirectory("bollywoodMovies")
	bmovies.Add(movie2)
	bmovies.Add(movie1)

	movieDir.Add(bmovies)

	hmovies := service.NewDirectory("hollywoodMovies")
	tollyMovies := service.NewDirectory("tollywoodMovies")
	hmovies.Add(tollyMovies)
	m, err := service.NewFile("mayadede.mov")
	tollyMovies.Add(m)
	//hmovies.Add(service.NewFile("ironman.mov"))
	movieDir.Add(hmovies)
	movieDir.Ls()
}
