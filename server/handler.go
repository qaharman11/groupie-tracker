package server

import (
	"groupie-tracker/api"
	"groupie-tracker/model"
	"net/http"
	"strconv"
	"strings"
	"text/template"
)

type dataErr struct {
	Id   int
	Text string
}

type allData struct {
	Art api.Artists
	Rel api.Relations
}

func GetArtistsById(id int) allData {
	oneArtistData := api.FullArtists[id-1]
	oneArtistDataRelation := api.Relinfo.Index[id-1]

	res := allData{
		Art: oneArtistData,
		Rel: oneArtistDataRelation,
	}
	return res
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		// http.Error(w, "404 not found", http.StatusNotFound)
		errorPage(w, "not found", http.StatusNotFound)
		return
	}
	html, err := template.ParseFiles("templates/home.html")
	if err != nil {
		errorPage(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	if r.Method != http.MethodGet {
		errorPage(w, "not allowed", http.StatusMethodNotAllowed)
		return
	}

	err = model.GetArtists()
	if err != nil {
		errorPage(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	err = html.Execute(w, api.FullArtists)
	if err != nil {
		errorPage(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}

func EachArtistHandler(w http.ResponseWriter, r *http.Request) {
	id := strings.Split(r.URL.Path, "/")
	idArtist, err := strconv.Atoi(id[2])
	if idArtist < 1 || idArtist > 52 {
		errorPage(w, "not found", http.StatusNotFound)
		return
	}
	if r.URL.Path != "/detail/"+id[2] {
		errorPage(w, "not found", http.StatusNotFound)
		return
	}
	html, err := template.ParseFiles("templates/artist.html")
	if err != nil {
		errorPage(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	if r.Method != http.MethodGet {
		errorPage(w, "not allowed", http.StatusMethodNotAllowed)
		return
	}

	if err != nil {
		errorPage(w, http.StatusText(http.StatusInternalServerError), http.StatusNotFound)
		return
	}

	if err != nil {
		errorPage(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	err = model.GetRelation()
	if err != nil {
		errorPage(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	res := GetArtistsById(idArtist)

	res.Art.DatesLocations = res.Rel.Details
	err = html.Execute(w, res.Art)
	if err != nil {
		errorPage(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}

func errorPage(w http.ResponseWriter, statusText string, statusID int) {
	temp, err := template.ParseFiles("templates/error.html")
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	res := dataErr{
		Id:   statusID,
		Text: statusText,
	}
	w.WriteHeader(statusID)
	temp.Execute(w, res)
}
