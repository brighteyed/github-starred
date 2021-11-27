package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/brighteyed/github-starred/repository"
	"github.com/google/go-github/v40/github"
)

func main() {
	user := flag.String("user", "", "github user name")
	flag.Parse()

	if *user == "" {
		flag.Usage()
		os.Exit(1)
	}

	client := github.NewClient(nil)

	// Check how many repositories starred the user
	_, resp, err := client.Activity.ListStarred(context.Background(), *user, &github.ActivityListStarredOptions{
		Sort:      "",
		Direction: "",
		ListOptions: github.ListOptions{
			Page:    0,
			PerPage: 1,
		},
	})

	if err != nil {
		log.Fatal(err)
	}

	// Peek random starred repository
	rand.Seed(time.Now().UnixMicro())
	repos, _, err := client.Activity.ListStarred(context.Background(), *user, &github.ActivityListStarredOptions{
		Sort:      "",
		Direction: "",
		ListOptions: github.ListOptions{
			Page:    rand.Intn(resp.LastPage-1) + 1,
			PerPage: 1,
		},
	})

	if err != nil {
		log.Fatal(err)
	}

	if len(repos) == 0 {
		log.Fatal("No repo found!")
	}

	// Output random starred repository
	starred := repos[0].GetRepository()
	fmt.Println(repository.NewStarred(*starred.HTMLURL, *starred.Description))
}
