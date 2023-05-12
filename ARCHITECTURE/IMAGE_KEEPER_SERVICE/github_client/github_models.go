package github_client

type Repository struct {
	Name      string `json:"name"`
	URL       string `json:"html_url"`
	UpdatedAt string `json:"updated_at"`
}

type File struct {
	Name string `json:"name"`
	Type string `json:"type"`
	URL  string `json:"url"`
}

type Release struct {
	TagName string `json:"tag_name"`
}
