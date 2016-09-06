package phpipam

type Client struct {
	Token				string
	ServerUrl		string
  Application string
}

func New(hostname string, application string, username string, password string) (*Client, error) {
  apiKey, err := NewLogin(hostname, application, username, password)
	if (err!=nil) {
    return nil, err
  }
  return &Client{
    Token: 				apiKey.Data.Token,
    ServerUrl:		hostname,
    Application:	application,
  }, nil
}
