package myshare

type response[T any] struct {
	StatusCode int `json:"statusCode"`
	Data       T   `json:"data"`
}
