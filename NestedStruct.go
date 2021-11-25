package main

import "fmt"

type Europe struct {
	Country string
	Info    Info
}

type Info struct {
	Name      string
	Language  string
	Ethnicity string
}

func main() {
	eng := &Europe{
		Country: "England",
		Info: Info{
			Name:      "Englang",
			Language:  "English",
			Ethnicity: "Angelcynn",
		},
	}

	ity := &Europe{
		Country: "Italy",
		Info: Info{
			Name:      "Italia",
			Language:  "Italian",
			Ethnicity: "Romance",
		},
	}

	fmt.Println(ity.Info.Ethnicity)
	fmt.Println(eng.Info.Language)
}
