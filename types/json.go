package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Movie struct {
	Name   string
	Year   int  `json: "released"`
	Color  bool `json: "color,omitempty"`
	Actors []string
}

var movies = []Movie{
	{Name: "M1", Year: 1998, Color: false, Actors: []string{"person1", "person2"}},
	{Name: "M2", Year: 2998, Color: true, Actors: []string{"person1-", "person2-"}},
	{Name: "M3", Year: 4998, Color: false, Actors: []string{"person1---", "person2---"}},
	{Name: "M4", Year: 2398, Color: true, Actors: []string{"person1000", "person2000"}},
}

func main() {
	data, err := json.MarshalIndent(movies, "", "    ")
	if err != nil {
		fmt.Fprintf(os.Stderr, "json marshal error: %v \n", err)
	}
	fmt.Printf("%s \n", data)
}
