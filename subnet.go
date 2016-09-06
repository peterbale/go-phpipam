package phpipam

import (
  "log"
  "net/http"
  "io/ioutil"
  "encoding/json"
)

type Calculation struct {
  Type          string `json:"Type"`
  IPAddress     string `json:"IP Address"`
  Network       string `json:"Network"`
  Broadcast     string `json:"Broadcast"`
  NumberOfHosts int `json:"Number of hosts"`
}

type Subnet struct {
  Code    int `json:"code"`
  Success bool `json:"success"`
  Data    SubnetData `json:"data"`
  Message string `json:"message"`
}

type SubnetData struct {
  Id            string `json:"id"`
  Subnet        string `json:"subnet"`
  Mask          string `json:"mask"`
  SectionId     string `json:"sectionId"`
  Description   string `json:"description"`
  IsFull        string `json:"isFull"`
  Calculation   Calculation `json:"calculation"`
}

func GetSubnet(server_url string, application string, subnetId string, token string) (*Subnet) {
  var subnetData = new(Subnet)
  client := &http.Client{}
  req, err := http.NewRequest("GET", "https://" + server_url + "/api/" + application + "/subnets/" + subnetId + "/", nil)
  req.Header.Add("token", token)
  resp, err := client.Do(req)
  if (err!=nil) {
    log.Fatal("Error Making Get Subnets Request: ", err)
  }
  body, err := ioutil.ReadAll(resp.Body)
  if (err!=nil) {
    log.Fatal("Error Reading Get Subnets Response: ", err)
  }
  json_err := json.Unmarshal([]byte(body), &subnetData)
  if(json_err != nil){
    log.Fatal("Error Parsing Get Subnets Response: ", json_err)
  }
  if subnetData.Code != 200 {
    log.Fatal("Get Subnets Failed: ", subnetData.Message)
  }
  return subnetData
}
