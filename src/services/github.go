package services

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"os"
)

const (
	GITHUB_API_URL = "https://api.github.com"
)

type ListBranchesResponse []struct {
	Name      string `json:"name"`
	Protected bool   `json:"protected"`
}

type Commit struct {
	Sha    string        `json:"sha"`
	Commit CommitDetails `json:"commit"`
}

type CommitDetails struct {
	Author       AuthorDetails `json:"author"`
	Committer    AuthorDetails `json:"committer"`
	Message      string        `json:"message"`
	Url          string        `json:"url"`
	CommentCount int           `json:"comment_count"`
}

type AuthorDetails struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Date  string `json:"date"`
}

type DetailBranchResponse struct {
	Name      string `json:"name"`
	Commit    Commit `json:"commit"`
	Protected bool   `json:"protected"`
}

var headers = map[string]string{
	"Accept":               "application/vnd.github.v3+json",
	"X-GitHub-Api-Version": "2022-11-28",
}

func GetListBranches(owner string, repo string) (*ListBranchesResponse, error) {
	var branches ListBranchesResponse
	url := GITHUB_API_URL + "/repos/" + owner + "/" + repo + "/branches"

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, errors.New("Error creating request")
	}

	for key, value := range headers {
		req.Header.Add(key, value)
	}
	req.Header.Add("Authorization", os.Getenv("GITHUB_TOKEN"))
	resp, err := client.Do(req)
	if err != nil {
		return nil, errors.New("error making request")
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		return nil, errors.New("error in response")
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.New("error reading response")
	}
	err = json.Unmarshal(body, &branches)
	if err != nil {
		return nil, errors.New("error unmarshalling response")
	}
	return &branches, nil
}

func GetBranchDetails(owner string, repo string, branch string) (*DetailBranchResponse, error) {
	var branchDetails DetailBranchResponse
	url := GITHUB_API_URL + "/repos/" + owner + "/" + repo + "/branches/" + branch

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, errors.New("Error creating request")
	}

	for key, value := range headers {
		req.Header.Add(key, value)
	}
	req.Header.Add("Authorization", os.Getenv("GITHUB_TOKEN"))
	resp, err := client.Do(req)
	if err != nil {
		return nil, errors.New("Error making request")
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		return nil, errors.New("Error in response")
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.New("Error reading response")
	}
	err = json.Unmarshal(body, &branchDetails)
	if err != nil {
		return nil, errors.New("error unmarshalling response")
	}
	return &branchDetails, nil
}
