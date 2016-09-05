package phpipam

import (
  "fmt"
  "net/http"
  "io/ioutil"
  "encoding/json"
)

type Login struct {
  Code    int `json:"code"`
  Success bool `json:"success"`
  Data    LoginData `json:"data"`
  Message string `json:"message"`
}

type LoginData struct {
  Token   string `json:"token"`
  Expires string `json:"expires"`
  Test    string `json:"test"`
}

func NewLogin(hostname string, application string, username string, password string) (*Login) {
  var loginData = new(Login)
  client := &http.Client{}
  req, err := http.NewRequest("POST", "https://" + hostname + "/api/" + application + "/user/", nil)
  req.SetBasicAuth(username, password)
  resp, err := client.Do(req)
  body, err := ioutil.ReadAll(resp.Body)
  if (err!=nil) {
    fmt.Print(err)
  }
  json_err := json.Unmarshal([]byte(body), &loginData)
  if(json_err != nil){
      fmt.Println("Failed to Unmarshal:", json_err)
  }
  return loginData
}
