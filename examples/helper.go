package examples

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"strconv"

	"github.com/mdvalv/go-tmdb"
	"github.com/pkg/errors"
)

func GetClient() *tmdb.Client {
	token := os.Getenv("API_TOKEN")
	client, err := tmdb.NewClient(token)
	if err != nil {
		panic(errors.Wrap(err, "failed to instantiate a new client"))
	}
	return client
}

func PrettyPrint(object interface{}) {
	buf := new(bytes.Buffer)
	enc := json.NewEncoder(buf)
	enc.SetEscapeHTML(false)
	enc.SetIndent("", "    ")
	if err := enc.Encode(object); err != nil {
		panic(errors.Wrap(err, "failed to pretty print"))
	}
	fmt.Println(string(buf.String()))
}

func RunExamples(examples ...func()) {
	args := os.Args[1:]
	if len(args) > 0 {
		if arg, err := strconv.Atoi(args[0]); err != nil || arg <= 0 || arg > len(examples) {
			fmt.Println("invalid option")
			return
		} else {
			examples[arg-1]()
		}
	} else {
		for _, example := range examples {
			example()
		}
	}
}

func PanicOnError(err error) {
	if err != nil {
		panic(errors.Wrap(err, "failed to run example"))
	}
}
