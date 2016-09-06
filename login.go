package phpipam

import (
  "net/http"
  "io/ioutil"
  "encoding/json"
  "errors"
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

func NewLogin(server_url string, application string, username string, password string) (*Login, error) {
  var loginData = new(Login)
  client := &http.Client{}
  req, _ := http.NewRequest("POST", "https://" + server_url + "/api/" + application + "/user/", nil)
  req.SetBasicAuth(username, password)
  resp, err := client.Do(req)
  if (err!=nil) {
    return loginData, err
  }
  body, err := ioutil.ReadAll(resp.Body)
  if (err!=nil) {
    return loginData, err
  }
  err = json.Unmarshal([]byte(body), &loginData)
  if(err != nil){
    return loginData, err
  }
  if loginData.Code != 200 {
    return loginData, errors.New(loginData.Message)
  }
  return loginData, nil
}
