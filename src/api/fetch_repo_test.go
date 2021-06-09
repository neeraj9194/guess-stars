package api

import (
	"encoding/json"
	"fmt"
	"gopkg.in/h2non/gock.v1"
	"reflect"
	"testing"
	"time"
)

// Test generating a github URL
func TestGenerateURL(t *testing.T) {
	createdDate := time.Now().AddDate(0, -1, 0).Format("2006-01-02")
	lang := "Go"
	expected := fmt.Sprintf(
		"https://api.github.com/search/repositories?sort=stars&order=desc&per_page=100&q=created:>%v+language:\"Go\"",
		createdDate)
	url := generateURL(lang)
	if expected != url {
		t.Fatal(fmt.Sprintf("Failed. %v is not equal to %v", expected, url))
	}
}

func TestGetRepositoryList(t *testing.T) {

	responseBytes := []byte(`{
		"total_count":811046,
		"incomplete_results":false,
		"items":[
		   {
			  "id":23096959,
			  "name":"go",
			  "full_name":"golang/go",
			  "owner":{
				 "login":"golang"
			  },
			  "description":"The Go programming language",
			  "created_at":"2014-08-19T04:33:40Z",
			  "updated_at":"2021-06-09T09:27:37Z",
			  "stargazers_count":86640,
			  "watchers_count":86640,
			  "language":"Go"
		   }
		]
	}`)
	var response RepoList

	// Mock request
	gock.New("https://api.github.com/search/repositories").
		Reply(200).
		JSON(responseBytes)

	repoList := GetRepositoryList("Go")
	json.Unmarshal(responseBytes, &response)

	if !reflect.DeepEqual(response.Repos, repoList) {
		t.Fatal("Failed!")
	}
}

func TestGetRepositoryEmptyList(t *testing.T) {

	responseBytes := []byte(`{
		"total_count": 0,
		"incomplete_results": false,
		"items": []
	}`)

	// Mock request
	gock.New("https://api.github.com/search/repositories").
		Reply(200).
		JSON(responseBytes)

	repoList := GetRepositoryList("Go")

	if !reflect.DeepEqual([]Repo{}, repoList) {
		t.Fatal("Failed!")
	}
}

func TestGetRepositoryError(t *testing.T) {

	responseBytes := []byte(`{
		"message": "Validation Failed",
		"errors": [
			{
			"message": "None of the search qualifiers apply to this search type.",
			"resource": "Search",
			"field": "q",
			"code": "invalid"
			}
		],
		"documentation_url": "https://docs.github.com/v3/search/"
	}`)

	// Mock request
	gock.New("https://api.github.com/search/repositories").
		Reply(200).
		JSON(responseBytes)

	repoList := GetRepositoryList("Go")

	if repoList != nil {
		t.Fatal("Failed!")
	}
}
