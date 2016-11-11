package helpers

import (
	"encoding/json"
	"net/http"
)

type Person struct {
	Name    string `json:"name"`
	Surname string `json:"surname"`
	Gender  string `json:"gender"`
	Region  string `json:"region"`
}

func GetPerson() (Person, error) {
	person := Person{}

	response, err := http.Get("http://uinames.com/api/?gender=male&region=united%20states")
	if err != nil {
		return Person{}, err
	}

	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&person)
	if err != nil {
		return Person{}, err
	}

	return person, nil
}

func GetName() (string, error) {
	person, err := GetPerson()
	if err != nil {
		return "", err
	}

	return person.Name + " " + person.Surname, nil
}
