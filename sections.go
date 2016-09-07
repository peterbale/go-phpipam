package phpipam

import (
  "errors"
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

type Section struct {
  Code    int `json:"code"`
  Success bool `json:"success"`
  Data    SectionsData `json:"data"`
  Message string `json:"message"`
}

type SectionData struct {
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

func GetSections(server_url string, application string, token string) (*Sections, error) {
  var sectionsData = new(Sections)
  client := &http.Client{}
  req, err := http.NewRequest("GET", "https://" + server_url + "/api/" + application + "/sections/", nil)
  req.Header.Add("token", token)
  resp, err := client.Do(req)
  if (err!=nil) {
    return sectionsData, err
  }
  body, err := ioutil.ReadAll(resp.Body)
  if (err!=nil) {
    return sectionsData, err
  }
  err = json.Unmarshal([]byte(body), &sectionsData)
  if(err != nil){
    return sectionsData, err
  }
  if sectionsData.Code != 200 {
    return sectionsData, errors.New(sectionsData.Message)
  }
  return sectionsData, nil
}

func GetSection(server_url string, application string, sectionId string, token string) (*Sections, error) {
  var sectionData = new(Section)
  client := &http.Client{}
  req, err := http.NewRequest("GET", "https://" + server_url + "/api/" + application + "/sections/" + sectionId + "/", nil)
  req.Header.Add("token", token)
  resp, err := client.Do(req)
  if (err!=nil) {
    return sectionData, err
  }
  body, err := ioutil.ReadAll(resp.Body)
  if (err!=nil) {
    return sectionData, err
  }
  err = json.Unmarshal([]byte(body), &sectionData)
  if(err != nil){
    return sectionData, err
  }
  if sectionData.Code != 200 {
    return sectionData, errors.New(sectionData.Message)
  }
  return sectionData, nil
}

func GetSectionsSubnets(server_url string, application string, sectionId string, token string) (*SectionsSubnets, error) {
  var sectionsSubnetsData = new(SectionsSubnets)
  client := &http.Client{}
  req, err := http.NewRequest("GET", "https://" + server_url + "/api/" + application + "/sections/" + sectionId + "/subnets/", nil)
  req.Header.Add("token", token)
  resp, err := client.Do(req)
  if (err!=nil) {
    return sectionsSubnetsData, err
  }
  body, err := ioutil.ReadAll(resp.Body)
  if (err!=nil) {
    return sectionsSubnetsData, err
  }
  err = json.Unmarshal([]byte(body), &sectionsSubnetsData)
  if(err != nil){
    return sectionsSubnetsData, err
  }
  if sectionsSubnetsData.Code != 200 {
    return sectionsSubnetsData, errors.New(sectionsSubnetsData.Message)
  }
  return sectionsSubnetsData, nil
}
