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

func (e example) GetTrendingParsingMediaTypes() {
	trending, _, err := e.client.Trending.GetTrending("week")
	if err != nil {
		panic(errors.Wrap(err, "failed to get trending"))
	}
	for _, result := range trending.Results {
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

func (e example) GetTrendingPeopleParsingKnowFor() {
	trending, _, err := e.client.Trending.GetTrendingPeople("week")
	if err != nil {
		panic(errors.Wrap(err, "failed to get trending people"))
	}
	for _, person := range trending.People {
		fmt.Println("->", person.Name, "known for:")
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

func (e example) GetTrending() {
	trending, _, err := e.client.Trending.GetTrending("day")
	if err != nil {
		panic(errors.Wrap(err, "failed to get trending"))
	}
	examples.PrettyPrint(trending)
}

func (e example) GetTrendingMovies() {
	trending, _, err := e.client.Trending.GetTrendingMovies("day")
	if err != nil {
		panic(errors.Wrap(err, "failed to get trending movies"))
	}
	examples.PrettyPrint(trending)
}

func (e example) GetTrendingTVShows() {
	trending, _, err := e.client.Trending.GetTrendingTVShows("day")
	if err != nil {
		panic(errors.Wrap(err, "failed to get trending tv shows"))
	}
	examples.PrettyPrint(trending)
}

func (e example) GetTrendingPeople() {
	trending, _, err := e.client.Trending.GetTrendingPeople("day")
	if err != nil {
		panic(errors.Wrap(err, "failed to get trending people"))
	}
	examples.PrettyPrint(trending)
}

func main() {
	example := example{
		client: examples.GetClient(),
	}

	examples.RunExamples(
		example.GetTrendingParsingMediaTypes,
		example.GetTrendingPeopleParsingKnowFor,
		example.GetTrending,
		example.GetTrendingMovies,
		example.GetTrendingTVShows,
		example.GetTrendingPeople,
	)
}
