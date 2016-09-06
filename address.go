package phpipam

import (
  "log"
  "net/http"
  "io/ioutil"
  "encoding/json"
  "strings"
)

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

func GetAddressSearch(server_url string, application string, searchHostname string, token string) (*AddressSearch) {
  var addressSearchData = new(AddressSearch)
  client := &http.Client{}
  req, err := http.NewRequest("GET", "https://" + server_url + "/api/" + application + "/addresses/search_hostname/" + searchHostname + "/", nil)
  req.Header.Add("token", token)
  resp, err := client.Do(req)
  if (err!=nil) {
    log.Fatal("Error Making Address Search Request: ", err)
  }
  body, err := ioutil.ReadAll(resp.Body)
  if (err!=nil) {
    log.Fatal("Error Reading Address Search Response: ", err)
  }
  json_err := json.Unmarshal([]byte(body), &addressSearchData)
  if(json_err != nil){
    log.Fatal("Error Parsing Address Search Response: ", json_err)
  }
  if addressSearchData.Code == 200 {
    // All is good
  } else if addressSearchData.Code == 404 {
    // Not found but all is still good
  } else {
    log.Fatal("Address Search Failed: ", addressSearchData.Message)
  }
  return addressSearchData
}

func GetAddressPing(server_url string, application string, addressId string, token string) (*AddressPing) {
  var addressPingData = new(AddressPing)
  client := &http.Client{}
  req, err := http.NewRequest("GET", "https://" + server_url + "/api/" + application + "/addresses/" + addressId + "/ping/", nil)
  req.Header.Add("token", token)
  resp, err := client.Do(req)
  if (err!=nil) {
    log.Fatal("Error Making Address Ping Request: ", err)
  }
  body, err := ioutil.ReadAll(resp.Body)
  if (err!=nil) {
    log.Fatal("Error Reading Address Ping Response: ", err)
  }
  json_err := json.Unmarshal([]byte(body), &addressPingData)
  if(json_err != nil){
    log.Fatal("Error Parsing Address Ping Response: ", json_err)
  }
  if addressPingData.Code == 200 {
    // All is good
  } else if addressPingData.Code == 404 {
    // Not found but all is still good
  } else {
    log.Fatal("Address Ping Failed: ", addressPingData.Message)
  }
  return addressPingData
}

func DeleteAddress(server_url string, application string, addressId string, token string) (*AddressDelete) {
  var addressDeleteData = new(AddressDelete)
  client := &http.Client{}
  req, err := http.NewRequest("DELETE", "https://" + server_url + "/api/" + application + "/addresses/" + addressId + "/", nil)
  req.Header.Add("token", token)
  resp, err := client.Do(req)
  if (err!=nil) {
    log.Fatal("Error Making Delete Address Request: ", err)
  }
  body, err := ioutil.ReadAll(resp.Body)
  if (err!=nil) {
    log.Fatal("Error Reading Delete Address Response: ", err)
  }
  json_err := json.Unmarshal([]byte(body), &addressDeleteData)
  if(json_err != nil){
    log.Fatal("Error Parsing Delete Address Response: ", json_err)
  }
  if addressDeleteData.Code == 200 {
    log.Fatal("Delete Address Failed: ", addressDeleteData.Message)
  }
  return addressDeleteData
}

func CreateAddressFirstFree(server_url string, application string, subnetId string, newHostname string, newOwner string, token string) (*AddressFirstFree) {
  var addressFirstFreeData = new(AddressFirstFree)
  client := &http.Client{}
  reqBody := "hostname=" + newHostname + "&owner=" + newOwner
  req, err := http.NewRequest("POST", "https://" + server_url + "/api/" + application + "/addresses/first_free/" + subnetId + "/", strings.NewReader(reqBody))
  req.Header.Add("token", token)
  resp, err := client.Do(req)
  if (err!=nil) {
    log.Fatal("Error Making Create Address First Free Request: ", err)
  }
  body, err := ioutil.ReadAll(resp.Body)
  if (err!=nil) {
    log.Fatal("Error Reading Create Address First Free Response: ", err)
  }
  json_err := json.Unmarshal([]byte(body), &addressFirstFreeData)
  if(json_err != nil){
    log.Fatal("Error Parsing Create Address First Free Response: ", json_err)
  }
  if addressFirstFreeData.Code == 201 {
    log.Fatal("Create Address First Free Failed: ", addressFirstFreeData.Message)
  }
  return addressFirstFreeData
}
