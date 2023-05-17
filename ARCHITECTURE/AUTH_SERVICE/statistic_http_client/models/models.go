package models

type Stats struct {
	TotalRequests   int64
	RegisteredUsers int
}

type Log struct {
	Message string `json:"message"`
}
