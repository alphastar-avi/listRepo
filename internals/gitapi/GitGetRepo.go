package gitapi

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func GetRepoNames(username, token string) []string {
	apiURL := fmt.Sprintf("https://api.github.com/users/%s/repos", username)

	req, _ := http.NewRequest("GET", apiURL, nil)
	req.Header.Set("Authorization", "token "+token)
	req.Header.Set("Accept", "application/vnd.github.v3+json")
	req.Header.Set("User-Agent", "my-github-app")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return []string{}
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return []string{}
	}

	var data []struct {
		Name string `json:"name"`
	}

	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return []string{}
	}

	var repoNames []string
	for _, repo := range data {
		repoNames = append(repoNames, repo.Name)
	}

	return repoNames
}
