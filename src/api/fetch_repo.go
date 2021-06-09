package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type RepoList struct {
	TotalCount        int    `json:"total_count"`
	IncompleteResults int    `json:"incomplete_results"`
	Repos             []Repo `json:"items"`
}

type Repo struct {
	Name        string `json:"name"`
	Author      Author `json:"owner"`
	Description string `json:"description"`
	Language    string `json:"language"`
	Stars       int    `json:"stargazers_count"`
}

type Author struct {
	Name string `json:"login"`
}

const githubBaseURL = "https://api.github.com/search/repositories"

// Generate URL to search git hub using language and that is created within a month.
func generateURL(language string) string {
	createdDate := time.Now().AddDate(0, -1, 0).Format("2006-01-02")
	param := "sort=stars&order=desc&per_page=100"

	if language != "" {
		param += fmt.Sprintf("&q=created:>%v+language:\"%v\"", createdDate, language)
	} else {
		param += fmt.Sprintf("&q=created:>%v", createdDate)
	}

	return githubBaseURL + "?" + param
}

func GetRepositoryList(language string) []Repo {

	client := &http.Client{}
	searchURL := generateURL(language)

	req, err := http.NewRequest("GET", searchURL, nil)
	if err != nil {
		fmt.Print(err.Error())
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Print(err.Error())
	}

	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Print(err.Error())
	}

	var responseObject RepoList
	json.Unmarshal(bodyBytes, &responseObject)

	return responseObject.Repos
}

func LanguageList() []string {
	languages := []string{
		"C++",
		"HTML",
		"Java",
		"JavaScript",
		"PHP",
		"Python",
		"Ruby",
		"C",
		"C#",
		"Dockerfile",
		"Go",
		"Swift",
		"TypeScript",
		"Visual Basic .NET",
		"YAML",
	}
	return languages
}
