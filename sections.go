package phpipam

import (
  "fmt"
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
  body, err := ioutil.ReadAll(resp.Body)
  if (err!=nil) {
    fmt.Print(err)
  }
  json_err := json.Unmarshal([]byte(body), &sectionsData)
  if(json_err != nil){
      fmt.Println("Failed to Unmarshal:", json_err)
  }
  return sectionsData
}

func GetSectionsSubnets(server_url string, application string, sectionId string, token string) (*SectionsSubnets) {
  var sectionsSubnetsData = new(SectionsSubnets)
  client := &http.Client{}
  req, err := http.NewRequest("GET", "https://" + server_url + "/api/" + application + "/sections/" + sectionId + "/subnets/", nil)
  req.Header.Add("token", token)
  resp, err := client.Do(req)
  body, err := ioutil.ReadAll(resp.Body)
  if (err!=nil) {
    fmt.Print(err)
  }
  json_err := json.Unmarshal([]byte(body), &sectionsSubnetsData)
  if(json_err != nil){
      fmt.Println("Failed to Unmarshal:", json_err)
  }
  return sectionsSubnetsData
}
