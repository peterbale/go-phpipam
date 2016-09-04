package phpipam

import (
  "fmt"
  "net/http"
  "io/ioutil"
  "encoding/json"
  "strings"
)

type Calculation struct {
  Type          string `json:"Type"`
  IPAddress     string `json:"IP Address"`
  Network       string `json:"Network"`
  Broadcast     string `json:"Broadcast"`
  NumberOfHosts int `json:"Number of hosts"`
}

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

type AddressSearch struct {
  Code    int `json:"code"`
  Success bool `json:"success"`
  Data    []AddressSearchData `json:"data"`
  Message string `json:"message"`
}

type AddressSearchData struct {
  Id          string `json:"id"`
  SubnetId    string `json:"subnetId"`
  Ip          string `json:"ip"`
  Description string `json:"description"`
}

type AddressPing struct {
  Code    int `json:"code"`
  Success bool `json:"success"`
  Data    AddressPingData `json:"data"`
  Message string `json:"message"`
}

type AddressPingData struct {
  ScanType  string `json:"scan_type"`
  ExitCode  int `json:"exit_code"`
}

type AddressDelete struct {
  Code    int `json:"code"`
  Success bool `json:"success"`
  Message string `json:"message"`
}

type AddressFirstFree struct {
  Code    int `json:"code"`
  Success bool `json:"success"`
  Message string `json:"message"`
  Ip      string `json:"ip"`
}

func login(hostname string, application string, username string, password string) (*Login) {
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

func sections(hostname string, application string, token string) (*Sections) {
  var sectionsData = new(Sections)
  client := &http.Client{}
  req, err := http.NewRequest("GET", "https://" + hostname + "/api/" + application + "/sections/", nil)
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

func sectionsSubnets(hostname string, application string, sectionId string, token string) (*SectionsSubnets) {
  var sectionsSubnetsData = new(SectionsSubnets)
  client := &http.Client{}
  req, err := http.NewRequest("GET", "https://" + hostname + "/api/" + application + "/sections/" + sectionId + "/subnets/", nil)
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

func subnet(hostname string, application string, subnetId string, token string) (*Subnet) {
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

func addressSearch(hostname string, application string, searchHostname string, token string) (*AddressSearch) {
  var addressSearchData = new(AddressSearch)
  client := &http.Client{}
  req, err := http.NewRequest("GET", "https://" + hostname + "/api/" + application + "/addresses/search_hostname/" + searchHostname + "/", nil)
  req.Header.Add("token", token)
  resp, err := client.Do(req)
  body, err := ioutil.ReadAll(resp.Body)
  if (err!=nil) {
    fmt.Print(err)
  }
  json_err := json.Unmarshal([]byte(body), &addressSearchData)
  if(json_err != nil){
      fmt.Println("Failed to Unmarshal:", json_err)
  }
  return addressSearchData
}

func addressPing(hostname string, application string, addressId string, token string) (*AddressPing) {
  var addressPingData = new(AddressPing)
  client := &http.Client{}
  req, err := http.NewRequest("GET", "https://" + hostname + "/api/" + application + "/addresses/" + addressId + "/ping/", nil)
  req.Header.Add("token", token)
  resp, err := client.Do(req)
  body, err := ioutil.ReadAll(resp.Body)
  if (err!=nil) {
    fmt.Print(err)
  }
  json_err := json.Unmarshal([]byte(body), &addressPingData)
  if(json_err != nil){
      fmt.Println("Failed to Unmarshal:", json_err)
  }
  return addressPingData
}

func addressDelete(hostname string, application string, addressId string, token string) (*AddressDelete) {
  var addressDeleteData = new(AddressDelete)
  client := &http.Client{}
  req, err := http.NewRequest("DELETE", "https://" + hostname + "/api/" + application + "/addresses/" + addressId + "/", nil)
  req.Header.Add("token", token)
  resp, err := client.Do(req)
  body, err := ioutil.ReadAll(resp.Body)
  if (err!=nil) {
    fmt.Print(err)
  }
  json_err := json.Unmarshal([]byte(body), &addressDeleteData)
  if(json_err != nil){
      fmt.Println("Failed to Unmarshal:", json_err)
  }
  return addressDeleteData
}

func addressFirstFree(hostname string, application string, subnetId string, newHostname string, newOwner string, token string) (*AddressFirstFree) {
  var addressFirstFreeData = new(AddressFirstFree)
  client := &http.Client{}
  reqBody := "hostname=" + newHostname + "&owner=" + newOwner
  req, err := http.NewRequest("POST", "https://" + hostname + "/api/" + application + "/addresses/first_free/" + subnetId + "/", strings.NewReader(reqBody))
  req.Header.Add("token", token)
  resp, err := client.Do(req)
  body, err := ioutil.ReadAll(resp.Body)
  if (err!=nil) {
    fmt.Print(err)
  }
  json_err := json.Unmarshal([]byte(body), &addressFirstFreeData)
  if(json_err != nil){
      fmt.Println("Failed to Unmarshal:", json_err)
  }
  return addressFirstFreeData
}
