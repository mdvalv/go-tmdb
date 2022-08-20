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

const (
	imdbId = "imdb_id"
	tvdbId = "tvdb_id"
)

func (e example) FindMovie() {
	findings, _, err := e.client.Find.Find("tt0421994", imdbId, nil)
	if err != nil {
		panic(errors.Wrap(err, "failed to find movie"))
	}
	examples.PrettyPrint(*findings)
}

func (e example) FindPerson() {
	findings, _, err := e.client.Find.Find("nm6442992", imdbId, nil)
	if err != nil {
		panic(errors.Wrap(err, "failed to find person"))
	}
	for _, person := range findings.People {
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

func (e example) FindTVShow() {
	findings, _, err := e.client.Find.Find("tt10638036", imdbId, nil)
	if err != nil {
		panic(errors.Wrap(err, "failed to find tv show"))
	}
	examples.PrettyPrint(*findings)
}

func (e example) FindEpisode() {
	findings, _, err := e.client.Find.Find("tt8160066", imdbId, nil)
	if err != nil {
		panic(errors.Wrap(err, "failed to find episode"))
	}
	examples.PrettyPrint(*findings)
}

func (e example) FindSeason() {
	findings, _, err := e.client.Find.Find("668343", tvdbId, nil)
	if err != nil {
		panic(errors.Wrap(err, "failed to find season"))
	}
	examples.PrettyPrint(*findings)
}

func main() {
	example := example{
		client: examples.GetClient(),
	}

	examples.RunExamples(
		example.FindMovie,
		example.FindPerson,
		example.FindTVShow,
		example.FindEpisode,
		example.FindSeason,
	)
}
