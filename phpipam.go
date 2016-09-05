package phpipam

type Client struct {
	ApiKey      string
	BaseUrl     string
  Application string
}

func New(hostname string, application string, username string, password string) (*Client, error) {
  apiKey := NewLogin(hostname, application, username, password)
  return &Client{
    apiKey.Data.Token,
    hostname,
    application,
  }, nil
}
