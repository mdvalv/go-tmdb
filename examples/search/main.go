package main

import (
	"fmt"

	"github.com/mdvalv/go-tmdb"
	"github.com/mdvalv/go-tmdb/examples"
	"github.com/pkg/errors"
)

type example struct {
	client *tmdb.Client
}

func (e example) Companies() {
	companies, _, err := e.client.Search.Companies("walt disney", nil)
	if err != nil {
		panic(errors.Wrap(err, "failed to search for companies"))
	}
	examples.PrettyPrint(*companies)
}

func (e example) Collections() {
	collections, _, err := e.client.Search.Collections("hunger games", nil)
	if err != nil {
		panic(errors.Wrap(err, "failed to search for collections"))
	}
	examples.PrettyPrint(*collections)
}

func (e example) Keywords() {
	keywords, _, err := e.client.Search.Keywords("coming of age", nil)
	if err != nil {
		panic(errors.Wrap(err, "failed to search for keywords"))
	}
	examples.PrettyPrint(*keywords)
}

func (e example) Movies() {
	movies, _, err := e.client.Search.Movies("alice junior", nil)
	if err != nil {
		panic(errors.Wrap(err, "failed to search for movies"))
	}
	examples.PrettyPrint(*movies)
}

func (e example) TVShows() {
	year := 2020
	opt := tmdb.SearchTVShowsOptions{
		FirstAirDateYear: &year,
	}
	tvShows, _, err := e.client.Search.TVShows("feel good", &opt)
	if err != nil {
		panic(errors.Wrap(err, "failed to search for tv shows"))
	}
	examples.PrettyPrint(*tvShows)
}

func (e example) People() {
	people, _, err := e.client.Search.People("alice wu", nil)
	if err != nil {
		panic(errors.Wrap(err, "failed to search for people"))
	}
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
	if err != nil {
		panic(errors.Wrap(err, "failed to search multi media"))
	}
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
		example.Companies,
		example.Collections,
		example.Keywords,
		example.Movies,
		example.TVShows,
		example.People,
		example.Multi,
	)
}
