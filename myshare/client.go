package myshare

type Client struct {
	baseUrl string
	clubId  string
}

func NewClient(baseUrl string, clubId string) *Client {
	return &Client{baseUrl, clubId}
}
