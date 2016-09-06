package phpipam

import (
  "net/http"
  "io/ioutil"
  "encoding/json"
  "strings"
  "errors"
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

func GetAddressSearch(server_url string, application string, searchHostname string, token string) (*AddressSearch, error) {
  var addressSearchData = new(AddressSearch)
  client := &http.Client{}
  req, err := http.NewRequest("GET", "https://" + server_url + "/api/" + application + "/addresses/search_hostname/" + searchHostname + "/", nil)
  req.Header.Add("token", token)
  resp, err := client.Do(req)
  if (err!=nil) {
    return addressSearchData, err
  }
  body, err := ioutil.ReadAll(resp.Body)
  if (err!=nil) {
    return addressSearchData, err
  }
  err = json.Unmarshal([]byte(body), &addressSearchData)
  if(err != nil){
    return addressSearchData, err
  }
  if addressSearchData.Code == 200 || addressSearchData.Code == 404 {
    // Accepted responses
  } else {
    return addressSearchData, errors.New(addressSearchData.Message)
  }
  return addressSearchData, nil
}

func GetAddressPing(server_url string, application string, addressId string, token string) (*AddressPing, error) {
  var addressPingData = new(AddressPing)
  client := &http.Client{}
  req, err := http.NewRequest("GET", "https://" + server_url + "/api/" + application + "/addresses/" + addressId + "/ping/", nil)
  req.Header.Add("token", token)
  resp, err := client.Do(req)
  if (err!=nil) {
    return addressPingData, err
  }
  body, err := ioutil.ReadAll(resp.Body)
  if (err!=nil) {
    return addressPingData, err
  }
  err = json.Unmarshal([]byte(body), &addressPingData)
  if(err != nil){
    return addressPingData, err
  }
  if addressPingData.Code == 200 || addressPingData.Code == 404 {
    // Accepted responses
  } else {
    return addressPingData, errors.New(addressPingData.Message)
  }
  return addressPingData, nil
}

func DeleteAddress(server_url string, application string, addressId string, token string) (*AddressDelete, error) {
  var addressDeleteData = new(AddressDelete)
  client := &http.Client{}
  req, err := http.NewRequest("DELETE", "https://" + server_url + "/api/" + application + "/addresses/" + addressId + "/", nil)
  req.Header.Add("token", token)
  resp, err := client.Do(req)
  if (err!=nil) {
    return addressDeleteData, err
  }
  body, err := ioutil.ReadAll(resp.Body)
  if (err!=nil) {
    return addressDeleteData, err
  }
  err = json.Unmarshal([]byte(body), &addressDeleteData)
  if(err != nil){
    return addressDeleteData, err
  }
  if addressDeleteData.Code != 200 {
    return addressDeleteData, errors.New(addressDeleteData.Message)
  }
  return addressDeleteData, nil
}

func CreateAddressFirstFree(server_url string, application string, subnetId string, newHostname string, newOwner string, token string) (*AddressFirstFree, error) {
  var addressFirstFreeData = new(AddressFirstFree)
  client := &http.Client{}
  reqBody := "hostname=" + newHostname + "&owner=" + newOwner
  req, err := http.NewRequest("POST", "https://" + server_url + "/api/" + application + "/addresses/first_free/" + subnetId + "/", strings.NewReader(reqBody))
  req.Header.Add("token", token)
  resp, err := client.Do(req)
  if (err!=nil) {
    return addressFirstFreeData, err
  }
  body, err := ioutil.ReadAll(resp.Body)
  if (err!=nil) {
    return addressFirstFreeData, err
  }
  json_err := json.Unmarshal([]byte(body), &addressFirstFreeData)
  if(json_err != nil){
    return addressFirstFreeData, err
  }
  if addressFirstFreeData.Code != 201 {
    return addressFirstFreeData, errors.New(addressFirstFreeData.Message)
  }
  return addressFirstFreeData, nil
}
