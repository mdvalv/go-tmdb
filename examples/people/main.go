package main

import (
	"fmt"

	"github.com/mdvalv/go-tmdb"
	"github.com/mdvalv/go-tmdb/examples"
)

type example struct {
	client *tmdb.Client
}

func (e example) GetPerson() {
	person, _, err := e.client.People.GetPerson(1196961, nil)
	examples.PanicOnError(err)
	examples.PrettyPrint(person)
}

func (e example) GetPersonAppendToResponse() {
	opt := tmdb.PersonDetailsOptions{
		AppendToResponse: "changes,combined_credits,external_ids,images,movie_credits,tagged_images,translations,tv_credits",
	}
	person, _, err := e.client.People.GetPerson(1196961, &opt)
	examples.PanicOnError(err)
	examples.PrettyPrint(person)
}

func (e example) GetChanges() {
	opt := tmdb.ChangesOptions{
		StartDate: "2019-05-20",
	}
	changes, _, err := e.client.People.GetChanges(68813, &opt)
	examples.PanicOnError(err)
	examples.PrettyPrint(changes)
}

func (e example) GetMovieCredits() {
	credits, _, err := e.client.People.GetMovieCredits(20387, nil)
	examples.PanicOnError(err)
	fmt.Println("-> Cast:")
	for _, credit := range credits.Cast {
		fmt.Printf("%s as %s\n", credit.Title, credit.Character)
	}
	fmt.Println("-> Crew:")
	for _, credit := range credits.Crew {
		fmt.Printf("%s in %s as %s\n", credit.Title, credit.Department, credit.Job)
	}
}

func (e example) GetTVCredits() {
	credits, _, err := e.client.People.GetTVCredits(20387, nil)
	examples.PanicOnError(err)
	fmt.Println("-> Cast:")
	for _, credit := range credits.Cast {
		fmt.Printf("%s as %s\n", credit.Name, credit.Character)
	}
	fmt.Println("-> Crew:")
	for _, credit := range credits.Crew {
		fmt.Printf("%s in %s as %s\n", credit.Name, credit.Department, credit.Job)
	}
}

func (e example) GetCombinedCredits() {
	credits, _, err := e.client.People.GetCombinedCredits(20387, nil)
	examples.PanicOnError(err)
	fmt.Println("-> Cast:")
	for _, credit := range credits.Cast {
		switch credit.GetMediaType() {
		case "movie":
			movieCast, err := credit.ToMovieCast()
			if err != nil {
				panic(err)
			}
			fmt.Printf("movie: %s as %s\n", movieCast.Title, movieCast.Character)
		case "tv":
			tvCast, err := credit.ToTVShowCast()
			if err != nil {
				panic(err)
			}
			fmt.Printf("tv: %s as %s\n", tvCast.Name, tvCast.Character)
		}
	}
	fmt.Println("-> Crew:")
	for _, credit := range credits.Crew {
		switch credit.GetMediaType() {
		case "movie":
			movieCast, err := credit.ToMovieCrew()
			if err != nil {
				panic(err)
			}
			fmt.Printf("movie: %s in %s as %s\n", movieCast.Title, movieCast.Department, movieCast.Job)
		case "tv":
			tvCast, err := credit.ToTVShowCrew()
			if err != nil {
				panic(err)
			}
			fmt.Printf("tv: %s in %s as %s\n", tvCast.Name, tvCast.Department, tvCast.Job)
		}
	}
}

func (e example) GetExternalIDs() {
	extIds, _, err := e.client.People.GetExternalIDs(2340180, nil)
	examples.PanicOnError(err)
	examples.PrettyPrint(*extIds)
}

func (e example) GetImages() {
	images, _, err := e.client.People.GetImages(1034197)
	examples.PanicOnError(err)
	examples.PrettyPrint(*images)
}

func (e example) GetTranslations() {
	translation, _, err := e.client.People.GetTranslations(472630, nil)
	examples.PanicOnError(err)
	examples.PrettyPrint(*translation)
}

func (e example) GetLatest() {
	latest, _, err := e.client.People.GetLatest(nil)
	examples.PanicOnError(err)
	examples.PrettyPrint(*latest)
}

func (e example) GetPopular() {
	popular, _, err := e.client.People.GetPopular(nil)
	examples.PanicOnError(err)
	examples.PrettyPrint(*popular)
}

func (e example) GetTaggedImages() {
	images, _, err := e.client.People.GetTaggedImages(505710, nil)
	examples.PanicOnError(err)
	examples.PrettyPrint(*images)
}

func (e example) GetPeopleChanges() {
	changes, _, err := e.client.People.GetPeopleChanges(nil)
	examples.PanicOnError(err)
	examples.PrettyPrint(*changes)
}

func main() {
	example := example{
		client: examples.GetClient(),
	}

	examples.RunExamples(
		example.GetPerson,                 // 1
		example.GetPersonAppendToResponse, // 2
		example.GetChanges,                // 3
		example.GetMovieCredits,           // 4
		example.GetTVCredits,              // 5
		example.GetCombinedCredits,        // 6
		example.GetExternalIDs,            // 7
		example.GetImages,                 // 8
		example.GetTranslations,           // 9
		example.GetLatest,                 // 10
		example.GetPopular,                // 11
		example.GetTaggedImages,           // 12
		example.GetPeopleChanges,          // 13
	)
}
