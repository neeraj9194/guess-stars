package api

import (
	"fmt"
	"log"
	
	"github.com/andygrunwald/go-trending"
)

func TrendingProjects() {
	trend := trending.NewTrending()

	// Show projects of today
	projects, err := trend.GetProjects(trending.TimeToday, "")
	if err != nil {
		log.Fatal(err)
	}
	for index, project := range projects {
		no := index + 1
		if len(project.Language) > 0 {
			fmt.Printf("%d: %s (written in %s with %d ★ )\n", no, project.Name, project.Language, project.Stars)
		} else {
			fmt.Printf("%d: %s (with %d ★ )\n", no, project.Name, project.Stars)
		}
	}
}
