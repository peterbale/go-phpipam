package phpipam

import (
	"crypto/tls"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

// Config struct for users to define when interacting with the library
type Config struct {
	Hostname      string
	Application   string
	Username      string
	Password      string
	SSLSkipVerify bool
}

// Client struct to define the phpipam master internal client
type Client struct {
	Token         string
	ServerURL     string
	Application   string
	SSLSkipVerify bool
}

// Login struct to define phpipam login return data
type Login struct {
	Code    int  `json:"code"`
	Success bool `json:"success"`
	Data    struct {
		Token   string `json:"token"`
		Expires string `json:"expires"`
		Test    string `json:"test"`
	} `json:"data"`
	Message string `json:"message"`
}

// NewClient New pointer method for all users to call to create a client using
// hostname string, application string, username string and password string,
// returns PhpIPAM pointer and error
func (c *Config) NewClient() (*Client, error) {
	apiKey, err := c.NewLogin()
	if err != nil {
		return nil, err
	}
	return &Client{
		Token:         apiKey.Data.Token,
		ServerURL:     c.Hostname,
		Application:   c.Application,
		SSLSkipVerify: c.SSLSkipVerify,
	}, nil
}

// NewLogin method to login to phpipam using server_url string, application
// string, username string and password string, returns Login struct pointer
// and error
func (c *Config) NewLogin() (*Login, error) {
	var loginData = new(Login)
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: c.SSLSkipVerify},
	}
	client := &http.Client{
		Transport: tr,
	}
	req, _ := http.NewRequest("POST", c.Hostname+"/api/"+c.Application+"/user/", nil)
	req.SetBasicAuth(c.Username, c.Password)
	resp, err := client.Do(req)
	if err != nil {
		return loginData, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return loginData, err
	}
	err = json.Unmarshal([]byte(body), &loginData)
	if err != nil {
		return loginData, err
	}
	if loginData.Code != 200 {
		return loginData, errors.New(loginData.Message)
	}
	return loginData, nil
}

// Do Client pointer method for all downstream methods to call to run requested
// action using req http.Request pointer, returns byte slice and error
func (c *Client) Do(req *http.Request) ([]byte, error) {
	var body []byte
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: c.SSLSkipVerify},
	}
	client := &http.Client{
		Transport: tr,
	}
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
