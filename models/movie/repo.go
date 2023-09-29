package movie

type Movie struct {
	Id int `json:"id"`
	Title string `json:"title"`
	Year int `json:"year"`
	Poster string `json:"poster"`
	Director string `json:"director"`
	Casts string `json:"casts"`
	Genre string `json:"genre"`
}