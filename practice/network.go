package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	response, err := http.Get("http://services.explorecalifornia.org/json/tours.php")
	if err != nil {
		panic(any(err))
	}
	fmt.Printf("Response data type %T\n", response)
	jsonResponse, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(any(err))
	}
	content := string(jsonResponse)
	//fmt.Println(content)
	tours := parseToursContent(content)
	for _, tour := range tours {
		fmt.Println(tour)
	}
	err = response.Body.Close()
	if err != nil {
		panic(any(err))
	}

}

func parseToursContent(content string) []Tour {
	tours := make([]Tour, 0, 20)
	decoder := json.NewDecoder(strings.NewReader(content))
	_, err := decoder.Token()
	if err != nil {
		panic(any(err))
	}

	var tour Tour
	for decoder.More() {
		err := decoder.Decode(&tour)
		if err != nil {
			panic(any(err))
		}
		tours = append(tours, tour)
	}

	return tours
}

type Tour struct {
	Name  string
	Price string
}
