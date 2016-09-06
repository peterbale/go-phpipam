package phpipam

import (
  "log"
  "net/http"
  "io/ioutil"
  "encoding/json"
)

type Sections struct {
  Code    int `json:"code"`
  Success bool `json:"success"`
  Data    []SectionsData `json:"data"`
  Message string `json:"message"`
}

type SectionsData struct {
  Id          string `json:"id"`
  Name        string `json:"name"`
  Description string `json:"description"`
}

type SectionsSubnets struct {
  Code    int `json:"code"`
  Success bool `json:"success"`
  Data    []SectionsSubnetsData `json:"data"`
  Message string `json:"message"`
}

type SectionsSubnetsData struct {
  Id            string `json:"id"`
  Subnet        string `json:"subnet"`
  Description   string `json:"description"`
}

func GetSections(server_url string, application string, token string) (*Sections) {
  var sectionsData = new(Sections)
  client := &http.Client{}
  req, err := http.NewRequest("GET", "https://" + server_url + "/api/" + application + "/sections/", nil)
  req.Header.Add("token", token)
  resp, err := client.Do(req)
  if (err!=nil) {
    log.Fatal("Error Making Get Sections Request: ", err)
  }
  body, err := ioutil.ReadAll(resp.Body)
  if (err!=nil) {
    log.Fatal("Error Reading Get Sections Response: ", err)
  }
  json_err := json.Unmarshal([]byte(body), &sectionsData)
  if(json_err != nil){
    log.Fatal("Error Parsing Get Sections Response: ", json_err)
  }
  if sectionsData.Code != 200 {
    log.Fatal("Get Sections Failed: ", sectionsData.Message)
  }
  return sectionsData
}

func GetSectionsSubnets(server_url string, application string, sectionId string, token string) (*SectionsSubnets) {
  var sectionsSubnetsData = new(SectionsSubnets)
  client := &http.Client{}
  req, err := http.NewRequest("GET", "https://" + server_url + "/api/" + application + "/sections/" + sectionId + "/subnets/", nil)
  req.Header.Add("token", token)
  resp, err := client.Do(req)
  if (err!=nil) {
    log.Fatal("Error Making Get Sections Subnets Request: ", err)
  }
  body, err := ioutil.ReadAll(resp.Body)
  if (err!=nil) {
    log.Fatal("Error Reading Get Sections Subnets Response: ", err)
  }
  json_err := json.Unmarshal([]byte(body), &sectionsSubnetsData)
  if(json_err != nil){
    log.Fatal("Error Parsing Get Sections Subnets Response: ", json_err)
  }
  if sectionsSubnetsData.Code != 200 {
    log.Fatal("Get Sections Failed: ", sectionsSubnetsData.Message)
  }
  return sectionsSubnetsData
}
