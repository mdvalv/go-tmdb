// Search examples.
package main

import (
	"fmt"

	"github.com/mdvalv/go-tmdb"
	"github.com/mdvalv/go-tmdb/examples"
)

type example struct {
	client *tmdb.Client
}

func (e example) Companies() {
	companies, _, err := e.client.Search.Companies("walt disney", nil)
	examples.PanicOnError(err)
	examples.PrettyPrint(*companies)
}

func (e example) Collections() {
	collections, _, err := e.client.Search.Collections("hunger games", nil)
	examples.PanicOnError(err)
	examples.PrettyPrint(*collections)
}

func (e example) Keywords() {
	keywords, _, err := e.client.Search.Keywords("coming of age", nil)
	examples.PanicOnError(err)
	examples.PrettyPrint(*keywords)
}

func (e example) Movies() {
	movies, _, err := e.client.Search.Movies("alice junior", nil)
	examples.PanicOnError(err)
	examples.PrettyPrint(*movies)
}

func (e example) TVShows() {
	year := 2020
	opt := tmdb.SearchTVShowsOptions{
		FirstAirDateYear: &year,
	}
	tvShows, _, err := e.client.Search.TVShows("feel good", &opt)
	examples.PanicOnError(err)
	examples.PrettyPrint(*tvShows)
}

func (e example) People() {
	people, _, err := e.client.Search.People("alice wu", nil)
	examples.PanicOnError(err)
	for _, person := range people.People {
		fmt.Printf("-> %s (%s), known for:\n", person.Name, person.KnownForDepartment)
		for _, work := range person.KnownFor {
			switch work.GetMediaType() {
			case "movie":
				movie, err := work.ToMovie()
				if err != nil {
					panic(err)
				}
				fmt.Printf("movie: %s\n", movie.Title)
			case "tv":
				tv, err := work.ToTVShow()
				if err != nil {
					panic(err)
				}
				fmt.Printf("tv: %s\n", tv.OriginalName)
			}
		}
	}
}

func (e example) Multi() {
	multi, _, err := e.client.Search.Multi("carol", nil)
	examples.PanicOnError(err)
	for _, result := range multi.Results {
		switch result.GetMediaType() {
		case "movie":
			movie, err := result.ToMovie()
			if err != nil {
				panic(err)
			}
			fmt.Printf("movie: %s\n", movie.Title)
		case "tv":
			tv, err := result.ToTVShow()
			if err != nil {
				panic(err)
			}
			fmt.Printf("tv: %s\n", tv.OriginalName)
		case "person":
			person, err := result.ToPerson()
			if err != nil {
				panic(err)
			}
			fmt.Printf("person: %s\n", person.Name)
		}
	}
}

func main() {
	example := example{
		client: examples.GetClient(),
	}

	examples.RunExamples(
		example.Companies,   // 1
		example.Collections, // 2
		example.Keywords,    // 3
		example.Movies,      // 4
		example.TVShows,     // 5
		example.People,      // 6
		example.Multi,       // 7
	)
}
