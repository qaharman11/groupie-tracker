package api

type Artists struct {
	Id             int      `json:"id"`
	Image          string   `json:"image"`
	Name           string   `json:"name"`
	Members        []string `json:"members"`
	CreationDate   int      `json:"creationDate"`
	FirstAlbum     string   `json:"firstAlbum"`
	Locations      string   `json:"locations"`
	ConcertDates   string   `json:"concertDates"`
	Relations      string   `json:"relations"`
	DatesLocations map[string][]string
}

type Relations struct {
	Id      int                 `json:"id"`
	Details map[string][]string `json:"datesLocations"`
}

var (
	BaseURL     string = "https://groupietrackers.herokuapp.com/api"
	FullArtists []Artists

	Relinfo struct {
		Index []Relations
	}
)
