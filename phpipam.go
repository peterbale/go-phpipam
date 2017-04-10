package phpipam

import (
	"io/ioutil"
	"net/http"
)

type Client struct {
	Token       string
	ServerUrl   string
	Application string
}

func New(hostname string, application string, username string, password string) (*Client, error) {
	apiKey, err := NewLogin(hostname, application, username, password)
	if err != nil {
		return nil, err
	}
	return &Client{
		Token:       apiKey.Data.Token,
		ServerUrl:   hostname,
		Application: application,
	}, nil
}

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
