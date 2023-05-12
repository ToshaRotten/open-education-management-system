package github_client

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type GithubClient struct {
	url           string
	token         string
	client        *http.Client
	latestVersion string
}

func New() *GithubClient {
	var g GithubClient
	g.url = ""
	g.token = ""
	g.client = &http.Client{}
	return &g
}

func (g *GithubClient) SetGithubUrl(url string) {
	g.url = url
}

func (g *GithubClient) SetGithubToken(token string) {
	g.token = token
}

func (g *GithubClient) GetReleases() []Release {
	req, err := http.NewRequest("GET", g.url+"/releases", nil)
	if err != nil {
		fmt.Println(err)
	}

	req.Header.Set("Accept", "application/vnd.github+json")

	resp, err := g.client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	var releases []Release
	err = json.Unmarshal(data, &releases)
	if err != nil {
		fmt.Println(err)
	}
	g.latestVersion = releases[0].TagName
	return releases
}

func (g *GithubClient) GetProjectRepo() Repository {
	req, err := http.NewRequest("GET", g.url, nil)
	if err != nil {
		fmt.Println(err)
	}
	req.Header.Set("Accept", "application/vnd.github+json")

	resp, err := g.client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	var repo Repository
	err = json.Unmarshal(data, &repo)
	if err != nil {
		fmt.Println(err)
	}
	return repo
}

func (g *GithubClient) GetRepoContent() {
	req, err := http.NewRequest("GET", g.url+"/contents", nil)
	if err != nil {
		fmt.Println(err)
	}

	resp, err := g.client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
	}

	var files []File
	err = json.Unmarshal(data, &files)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
	}
}

func (g *GithubClient) DownloadSourceCode() {
	releases := g.GetReleases()
	g.latestVersion = releases[0].TagName

	url := fmt.Sprintf("https://github.com/ToshaRotten/open-education-management-system/archive/refs/tags/%s.zip", g.latestVersion)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
	}
	req.Header.Set("Accept", "application/vnd.github.v3+json")

	resp, err := g.client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
	}
	defer resp.Body.Close()

	filename := fmt.Sprintf("../source/zip/source_%s.zip", g.latestVersion)

	err = createFile(filename, resp.Body)
	if err != nil {
		fmt.Println("Error downloading file:", err)
	}
	fmt.Println("File downloaded successfully!")
}
