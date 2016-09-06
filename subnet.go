package phpipam

import (
  "errors"
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

func GetSubnet(server_url string, application string, subnetId string, token string) (*Subnet, error) {
  var subnetData = new(Subnet)
  client := &http.Client{}
  req, err := http.NewRequest("GET", "https://" + server_url + "/api/" + application + "/subnets/" + subnetId + "/", nil)
  req.Header.Add("token", token)
  resp, err := client.Do(req)
  if (err!=nil) {
    return subnetData, err
  }
  body, err := ioutil.ReadAll(resp.Body)
  if (err!=nil) {
    return subnetData, err
  }
  err = json.Unmarshal([]byte(body), &subnetData)
  if(err != nil){
    return subnetData, err
  }
  if subnetData.Code != 200 {
    return subnetData, errors.New(subnetData.Message)
  }
  return subnetData, nil
}
