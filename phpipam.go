package phpipam

import (
	"io/ioutil"
	"net/http"
)

// Client struct to define the phpipam master client
type Client struct {
	Token       string
	ServerURL   string
	Application string
}

// New method for all users to call to create a client using hostname string,
// application string, username string and password string, returns Client
// pointer and error
func New(hostname string, application string, username string, password string) (*Client, error) {
	apiKey, err := NewLogin(hostname, application, username, password)
	if err != nil {
		return nil, err
	}
	return &Client{
		Token:       apiKey.Data.Token,
		ServerURL:   hostname,
		Application: application,
	}, nil
}

// Do Client pointer method for all downstream methods to call to run requested
// action using req http.Request pointer, returns byte slice and error
func (c *Client) Do(req *http.Request) ([]byte, error) {
	var body []byte
	client := &http.Client{}
	req.Header.Add("token", c.Token)
	resp, err := client.Do(req)
	if err != nil {
		return body, err
	}
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return body, err
	}
	return body, nil
}
