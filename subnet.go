package phpipam

import (
  "fmt"
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

func GetSubnet(hostname string, application string, subnetId string, token string) (*Subnet) {
  var subnetData = new(Subnet)
  client := &http.Client{}
  req, err := http.NewRequest("GET", "https://" + hostname + "/api/" + application + "/subnets/" + subnetId + "/", nil)
  req.Header.Add("token", token)
  resp, err := client.Do(req)
  body, err := ioutil.ReadAll(resp.Body)
  if (err!=nil) {
    fmt.Print(err)
  }
  json_err := json.Unmarshal([]byte(body), &subnetData)
  if(json_err != nil){
      fmt.Println("Failed to Unmarshal:", json_err)
  }
  return subnetData
}
