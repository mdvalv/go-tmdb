package main

import (
	"fmt"
	"os"

	"github.com/mdvalv/go-tmdb"
	"github.com/mdvalv/go-tmdb/examples"
)

type example struct {
	client *tmdb.Client
}

const (
	listId = "list_id"
)

var sessionId = os.Getenv("SESSIONID")

func (e example) GetList() {
	list, _, err := e.client.Lists.GetList(listId, nil)
	examples.PanicOnError(err)

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
	examples.PanicOnError(err)
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
	examples.PanicOnError(err)
	examples.PrettyPrint(*status)
}

func (e example) CreateList() {
	list := tmdb.CreateList{
		Name:        "new list",
		Description: "new list description",
	}

	response, _, err := e.client.Lists.CreateList(sessionId, list)
	examples.PanicOnError(err)
	examples.PrettyPrint(*response)
}

func (e example) AddMovie() {
	response, _, err := e.client.Lists.AddMovie(sessionId, listId, 597219)
	examples.PanicOnError(err)
	examples.PrettyPrint(*response)
}

func (e example) RemoveMovie() {
	response, _, err := e.client.Lists.RemoveMovie(sessionId, listId, 597219)
	examples.PanicOnError(err)
	examples.PrettyPrint(*response)
}

func (e example) Clear() {
	response, _, err := e.client.Lists.Clear(sessionId, listId)
	examples.PanicOnError(err)
	examples.PrettyPrint(*response)
}

func (e example) Delete() {
	response, _, err := e.client.Lists.Delete(sessionId, listId)
	examples.PanicOnError(err)
	examples.PrettyPrint(*response)
}

func main() {
	example := example{
		client: examples.GetClient(),
	}

	examples.RunExamples(
		example.GetList,            // 1
		example.GetListWithOptions, // 2
		example.GetItemStatus,      // 3
		example.CreateList,         // 4
		example.AddMovie,           // 5
		example.RemoveMovie,        // 6
		example.Clear,              // 7
		example.Delete,             // 8
	)
}
