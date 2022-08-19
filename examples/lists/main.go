package main

import (
	"fmt"
	"os"

	"github.com/mdvalv/go-tmdb"
	"github.com/mdvalv/go-tmdb/examples"
	"github.com/pkg/errors"
)

type example struct {
	client *tmdb.Client
}

const (
	listId    = "list_id"
)

var sessionId = os.Getenv("SESSIONID")

func (e example) GetList() {
	list, _, err := e.client.Lists.GetList(listId, nil)
	if err != nil {
		panic(errors.Wrap(err, "failed to get list"))
	}

	for _, item := range list.Items {
		switch item.GetMediaType() {
		case "movie":
			movie, err := item.ToMovie()
			if err != nil {
				panic(err)
			}
			fmt.Printf("movie: %s\n", movie.Title)
		case "tv":
			tv, err := item.ToTVShow()
			if err != nil {
				panic(err)
			}
			fmt.Printf("tv: %s\n", tv.OriginalName)
		}
	}
}

func (e example) GetListWithOptions() {
	opt := tmdb.ListOptions{
		Language: "pt-BR",
	}
	list, _, err := e.client.Lists.GetList(listId, &opt)
	if err != nil {
		panic(errors.Wrap(err, "failed to get list"))
	}
	for _, item := range list.Items {
		switch item.GetMediaType() {
		case "movie":
			movie, err := item.ToMovie()
			if err != nil {
				panic(err)
			}
			fmt.Printf("movie: %s\n", movie.Title)
		case "tv":
			tv, err := item.ToTVShow()
			if err != nil {
				panic(err)
			}
			fmt.Printf("tv: %s\n", tv.OriginalName)
		}
	}
}

func (e example) GetItemStatus() {
	status, _, err := e.client.Lists.GetItemStatus(listId, 73939)
	if err != nil {
		panic(errors.Wrap(err, "failed to get item status"))
	}
	examples.PrettyPrint(*status)
}

func (e example) CreateList() {
	list := tmdb.CreateList{
		Name:        "new list",
		Description: "new list description",
		Language:    "asdf",
	}

	response, _, err := e.client.Lists.CreateList(sessionId, list)
	if err != nil {
		panic(errors.Wrap(err, "failed to create list"))
	}
	examples.PrettyPrint(*response)
}

func (e example) AddMovie() {
	response, _, err := e.client.Lists.AddMovie(sessionId, listId, 597219)
	if err != nil {
		panic(errors.Wrap(err, "failed to add movie"))
	}
	examples.PrettyPrint(*response)
}

func (e example) RemoveMovie() {
	response, _, err := e.client.Lists.RemoveMovie(sessionId, listId, 597219)
	if err != nil {
		panic(errors.Wrap(err, "failed to remove movie"))
	}
	examples.PrettyPrint(*response)
}

func (e example) Clear() {
	response, _, err := e.client.Lists.Clear(sessionId, listId)
	if err != nil {
		panic(errors.Wrap(err, "failed to clear list"))
	}
	examples.PrettyPrint(*response)
}

func (e example) Delete() {
	response, _, err := e.client.Lists.Delete(sessionId, listId)
	if err != nil {
		panic(errors.Wrap(err, "failed to delete list"))
	}
	examples.PrettyPrint(*response)
}

func main() {
	example := example{
		client: examples.GetClient(),
	}

	examples.RunExamples(
		example.GetList,
		example.GetListWithOptions,
		example.GetItemStatus,
		example.CreateList,
		example.AddMovie,
		example.RemoveMovie,
		example.Clear,
		example.Delete,
	)
}
