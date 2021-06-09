package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
	"log"

	fetchRepo "github.com/neeraj9194/guess-stars/src/api"
)

func main() {
	lang := languageSelection()

	winResult := 0

	listProjects := fetchRepo.GetRepositoryList(lang)
	if listProjects == nil {
		log.Fatal("Could not fetch repositories from Github.")
	}
	
	listProjects = RandomTrendingList(5, listProjects)

	fmt.Println("---------Lets Start------------")

	for _, pr := range listProjects {
		var guessStars int

		fmt.Printf(
			"\n Name: %v\n Author: %v\n Language: %v\n Description: %v\n",
			pr.Name, pr.Author.Name, pr.Language, pr.Description)
		fmt.Println("Enter the number of stars for the given project:")
		_, err := fmt.Scan(&guessStars)
		if err != nil {
			fmt.Println("Please enter a valid integer.")
			return
		}
		var tolerancePer int
		if pr.Stars == 0 {
			// Div by 0
			if guessStars == 0 {
				fmt.Println("You are correct!")
				winResult++
			} else {
				fmt.Println("Wrong!")
			}
		} else {
			tolerancePer = ((guessStars - pr.Stars) * 100) / pr.Stars
			if tolerancePer >= -10 && tolerancePer <= 10 {
				fmt.Println("You are correct!")
				winResult++
			} else {
				fmt.Println("Wrong!")
			}
		}
	}

	if winResult > 4 {
		fmt.Printf("\n\nCongrat you won. You guessed %v repo correctly.\n", winResult)
	} else {
		fmt.Printf("\n\nYou lost. You guessed %v repo correctly.\n", winResult)
	}
	printProjects(listProjects)
}

func printProjects(projects []fetchRepo.Repo) {
	for index, project := range projects {
		no := index + 1
		if len(project.Language) > 0 {
			fmt.Printf("%d: %s (written in %s with %d ★ )\n", no, project.Name, project.Language, project.Stars)
		} else {
			fmt.Printf("%d: %s (with %d ★ )\n", no, project.Name, project.Stars)
		}
	}
}

func RandomTrendingList(limit int, projects []fetchRepo.Repo) []fetchRepo.Repo {
	var v []fetchRepo.Repo
	var randList []int
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < limit; i++ {
		temp := rand.Intn(len(projects))
		if !ElementExist(randList[:], temp) {
			randList = append(randList, temp)
			v = append(v, projects[temp])
		} else {
			i--
		}
	}
	return v
}

// Find takes a slice and looks for an element in it.
func ElementExist(slice []int, val int) bool {
	for _, item := range slice {
		if item == val {
			return true
		}
	}
	return false
}

func languageSelection() string {
	langList := fetchRepo.LanguageList()

	keys := make([]string, 0, len(langList))
	for k := range langList {
		keys = append(keys, k)
	}

	sort.Strings(keys)
	for i, item := range keys {
		fmt.Printf("%d: %v\n", i+1, item)
	}

	var language int
	fmt.Println("Select the language from the above list (0 for all):")
	_, err := fmt.Scan(&language)
	if err != nil {
		fmt.Println("Please enter a valid integer.")
		return ""
	}
	if language == 0 || language > len(keys) {
		return ""
	}
	return langList[keys[language-1]]
}
