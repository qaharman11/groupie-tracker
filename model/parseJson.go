package model

import (
	"encoding/json"
	"fmt"
	"groupie-tracker/api"
	"io/ioutil"
	"net/http"
)

func GetArtists() error {
	resp, err := http.Get(api.BaseURL + "/artists")
	if err != nil {
		fmt.Println("No response from request")
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)                         // response body is []byte
	if err := json.Unmarshal(body, &api.FullArtists); err != nil { // Parse []byte to the go struct pointer
		fmt.Println("Can not unmarshal JSON")
	}

	return nil
}

func GetRelation() error {
	resp, err := http.Get(api.BaseURL + "/relation")
	if err != nil {
		fmt.Println("No response from request")
		return err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err := json.Unmarshal(body, &api.Relinfo); err != nil {
		fmt.Println("Can not unmarshal JSON")
	}

	return nil
}
