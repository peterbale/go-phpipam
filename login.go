package phpipam

import (
  "log"
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

func NewLogin(server_url string, application string, username string, password string) (*Login) {
  var loginData = new(Login)
  client := &http.Client{}
  req, _ := http.NewRequest("POST", "https://" + server_url + "/api/" + application + "/user/", nil)
  req.SetBasicAuth(username, password)
  resp, err := client.Do(req)
  if (err!=nil) {
    log.Fatal("Error Making Login Request: ", err)
  }
  body, err := ioutil.ReadAll(resp.Body)
  if (err!=nil) {
    log.Fatal("Error Reading Login Response: ", err)
  }
  json_err := json.Unmarshal([]byte(body), &loginData)
  if(json_err != nil){
    log.Fatal("Error Parsing Login Response: ", json_err)
  }
  if loginData.Code != 200 {
    log.Fatal("Login Failed: ", loginData.Message)
  }
  return loginData
}
