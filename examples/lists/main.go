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
	listID = "list_id"
)

var sessionID = os.Getenv("SESSIONID")

func (e example) GetList() {
	list, _, err := e.client.Lists.GetList(listID, nil)
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
	list, _, err := e.client.Lists.GetList(listID, &opt)
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
	status, _, err := e.client.Lists.GetItemStatus(listID, 73939)
	examples.PanicOnError(err)
	examples.PrettyPrint(*status)
}

func (e example) CreateList() {
	list := tmdb.CreateList{
		Name:        "new list",
		Description: "new list description",
	}

	response, _, err := e.client.Lists.CreateList(sessionID, list)
	examples.PanicOnError(err)
	examples.PrettyPrint(*response)
}

func (e example) AddMovie() {
	response, _, err := e.client.Lists.AddMovie(sessionID, listID, 597219)
	examples.PanicOnError(err)
	examples.PrettyPrint(*response)
}

func (e example) RemoveMovie() {
	response, _, err := e.client.Lists.RemoveMovie(sessionID, listID, 597219)
	examples.PanicOnError(err)
	examples.PrettyPrint(*response)
}

func (e example) Clear() {
	response, _, err := e.client.Lists.Clear(sessionID, listID)
	examples.PanicOnError(err)
	examples.PrettyPrint(*response)
}

func (e example) Delete() {
	response, _, err := e.client.Lists.Delete(sessionID, listID)
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
