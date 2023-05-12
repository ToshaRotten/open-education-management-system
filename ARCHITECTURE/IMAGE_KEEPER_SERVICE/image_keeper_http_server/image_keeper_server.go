package image_keeper_http_server

import (
	"main/github_client"
	"main/image_keeper"
)

type Server struct {
	GithubClient *github_client.GithubClient
	ImageKeeper  *image_keeper.ImageKeeper
}
