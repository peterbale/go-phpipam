package phpipam

type Client struct {
	Token				string
	ServerUrl		string
  Application string
}

func New(hostname string, application string, username string, password string) (*Client, error) {
  apiKey := NewLogin(hostname, application, username, password)
  return &Client{
    Token: 				apiKey.Data.Token,
    ServerUrl:		hostname,
    Application:	application,
  }, nil
}
