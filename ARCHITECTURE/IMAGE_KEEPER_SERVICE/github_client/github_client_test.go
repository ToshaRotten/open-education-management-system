package github_client

import (
	"fmt"
	"testing"
)

func TestGithubClient_GetProjectRepo(t *testing.T) {
	client := New()
	client.SetGithubUrl("https://api.github.com/repos/ToshaRotten/open-education-management-system")
	//client.SetGithubToken()
	repo := client.GetProjectRepo()
	fmt.Println(repo)
}

func TestGithubClient_GetReleases(t *testing.T) {
	client := New()
	client.SetGithubUrl("https://api.github.com/repos/ToshaRotten/open-education-management-system")
	//client.SetGithubToken()
	releases := client.GetReleases()
	fmt.Println(releases)
}

func TestGithubClient_DownloadSourceCode(t *testing.T) {
	client := New()
	client.SetGithubUrl("https://api.github.com/repos/ToshaRotten/open-education-management-system")
	//client.SetGithubToken()
	client.DownloadSourceCode()
}
