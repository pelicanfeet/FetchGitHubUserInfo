package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type GitHubUser struct {
	Login     string `json:"login"`
	Name      string `json:"name"`
	Followers int    `json:"followers"`
}

func fetchGitHubUser(username string) (*GitHubUser, error) {
	apiURL := fmt.Sprintf("https://api.github.com/users/%s", username)

	response, err := http.Get(apiURL)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API request failed with status code: %d", response.StatusCode)
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var user GitHubUser
	err = json.Unmarshal(body, &user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func main() {
	username := "pelicanfeet" // Replace with the GitHub username you want to fetch

	user, err := fetchGitHubUser(username)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("Username: %s\n", user.Login)
	fmt.Printf("Name: %s\n", user.Name)
	fmt.Printf("Followers: %d\n", user.Followers)
}
