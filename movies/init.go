package movies

import "net/http"

type Movie struct {
	Movies []MovieData `json:"movies"`
}

type MovieData struct {
	ID    string `json:"id"`
	Title string `json:"title"`
}

func InitMovie() *Movie {
	return &Movie{
		Movies: getMovies(),
	}
}

func (movies *Movie) PrintMovies(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("List of movies that I have been watched in 2019:\n"))
	for _, val := range movies.Movies {
		w.Write([]byte(val.Title + "\n"))
	}
}

// GetMovies : get list of movies
func getMovies() []MovieData {

	listMovies := []MovieData{
		MovieData{
			ID:    "1",
			Title: "Revenge",
		},
		MovieData{
			ID:    "2",
			Title: "Polar",
		},
	}

	return listMovies
}
