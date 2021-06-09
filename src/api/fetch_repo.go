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
	IncompleteResults bool   `json:"incomplete_results"`
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
	fmt.Printf("%v\n", searchURL)

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

func LanguageList() map[string]string {
	return map[string]string{
		"C++":               "cpp",
		"HTML":              "HTML",
		"Java":              "Java",
		"JavaScript":        "JavaScript",
		"PHP":               "PHP",
		"Python":            "Python",
		"Ruby":              "Ruby",
		"C":                 "C",
		"C#":                "Csharp",
		"Dockerfile":        "Dockerfile",
		"Go":                "Go",
		"Swift":             "Swift",
		"TypeScript":        "TypeScript",
		"Visual Basic .NET": "Visual+Basic+.NET",
		"YAML":              "YAML",
	}
}
