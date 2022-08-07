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

func (e example) GetPerson() {
	person, _, err := e.client.People.GetPerson(1196961)
	if err != nil {
		panic(errors.Wrap(err, "failed to get person"))
	}
	examples.PrettyPrint(person)
}

func (e example) GetPersonAppendToResponse() {
	person, _, err := e.client.People.GetPerson(1196961, "changes", "combined_credits", "external_ids", "images", "movie_credits", "tagged_images", "translations", "tv_credits")
	if err != nil {
		panic(errors.Wrap(err, "failed to get person"))
	}
	examples.PrettyPrint(person)
}

func (e example) GetChanges() {
	opt := tmdb.PeopleChangesOptions{
		StartDate: "2019-05-20",
	}
	changes, _, err := e.client.People.GetChanges(68813, &opt)
	if err != nil {
		panic(errors.Wrap(err, "failed to get changes"))
	}
	examples.PrettyPrint(changes)
}

func (e example) GetMovieCredits() {
	credits, _, err := e.client.People.GetMovieCredits(20387, nil)
	if err != nil {
		panic(errors.Wrap(err, "failed to get movie credits"))
	}
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
	if err != nil {
		panic(errors.Wrap(err, "failed to get tv credits"))
	}
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
	if err != nil {
		panic(errors.Wrap(err, "failed to get combined credits"))
	}
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
	if err != nil {
		panic(errors.Wrap(err, "failed to get external ids"))
	}
	examples.PrettyPrint(*extIds)
}

func (e example) GetImages() {
	images, _, err := e.client.People.GetImages(1034197)
	if err != nil {
		panic(errors.Wrap(err, "failed to get images"))
	}
	examples.PrettyPrint(*images)
}

func (e example) GetTranslations() {
	translation, _, err := e.client.People.GetTranslations(472630, nil)
	if err != nil {
		panic(errors.Wrap(err, "failed to get translation"))
	}
	examples.PrettyPrint(*translation)
}

func (e example) GetLatest() {
	latest, _, err := e.client.People.GetLatest(nil)
	if err != nil {
		panic(errors.Wrap(err, "failed to get latest person"))
	}
	examples.PrettyPrint(*latest)
}

func (e example) GetPopular() {
	popular, _, err := e.client.People.GetPopular(nil)
	if err != nil {
		panic(errors.Wrap(err, "failed to get popular person"))
	}
	examples.PrettyPrint(*popular)
}

func (e example) GetTaggedImages() {
	images, _, err := e.client.People.GetTaggedImages(505710, nil)
	if err != nil {
		panic(errors.Wrap(err, "failed to get images"))
	}
	examples.PrettyPrint(*images)
}

func (e example) GetPeopleChanges() {
	changes, _, err := e.client.People.GetPeopleChanges(nil)
	if err != nil {
		panic(errors.Wrap(err, "failed to get people changes"))
	}
	examples.PrettyPrint(*changes)
}

func main() {
	example := example{
		client: examples.GetClient(),
	}

	examples.RunExamples(
		example.GetPerson,
		example.GetPersonAppendToResponse,
		example.GetChanges,
		example.GetMovieCredits,
		example.GetTVCredits,
		example.GetCombinedCredits,
		example.GetExternalIDs,
		example.GetImages,
		example.GetTranslations,
		example.GetLatest,
		example.GetPopular,
		example.GetTaggedImages,
		example.GetPeopleChanges,
	)
}
