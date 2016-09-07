package phpipam

import (
  "net/http"
  "io/ioutil"
  "encoding/json"
  "strings"
  "errors"
)

type Address struct {
  Code    int `json:"code"`
  Success bool `json:"success"`
  Data    AddressData `json:"data"`
  Message string `json:"message"`
}

type AddressData struct {
  Id        string `json:"id"`
  SubnetId  string `json:"subnetId"`
  Ip        string `json:"ip"`
  Hostname  string `json:"hostname"`
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

type AddressSearchIp struct {
  Code    int `json:"code"`
  Success bool `json:"success"`
  Data    []AddressSearchIpData `json:"data"`
  Message string `json:"message"`
}

type AddressSearchIpData struct {
  Id          string `json:"id"`
  SubnetId    string `json:"subnetId"`
  Ip          string `json:"ip"`
  Hostname    string `json:"hostname"`
}

type UpdateAddress struct {
  Code    int `json:"code"`
  Success bool `json:"success"`
  Message string `json:"message"`
}

func GetAddress(server_url string, application string, addressId string, token string) (*Address, error) {
  var addressData = new(Address)
  client := &http.Client{}
  req, err := http.NewRequest("GET", "https://" + server_url + "/api/" + application + "/addresses/" + addressId + "/", nil)
  req.Header.Add("token", token)
  resp, err := client.Do(req)
  if (err!=nil) {
    return addressData, err
  }
  body, err := ioutil.ReadAll(resp.Body)
  if (err!=nil) {
    return addressData, err
  }
  err = json.Unmarshal([]byte(body), &addressData)
  if(err != nil){
    return addressData, err
  }
  if addressData.Code == 200 || addressData.Code == 404 {
    // Accepted responses
  } else {
    return addressData, errors.New(addressData.Message)
  }
  return addressData, nil
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

func GetAddressSearchIp(server_url string, application string, address string, token string) (*AddressSearchIp, error) {
  var addressSearchIpData = new(AddressSearchIp)
  client := &http.Client{}
  req, err := http.NewRequest("GET", "https://" + server_url + "/api/" + application + "/addresses/search/" + address + "/", nil)
  req.Header.Add("token", token)
  resp, err := client.Do(req)
  if (err!=nil) {
    return addressSearchIpData, err
  }
  body, err := ioutil.ReadAll(resp.Body)
  if (err!=nil) {
    return addressSearchIpData, err
  }
  json_err := json.Unmarshal([]byte(body), &addressSearchIpData)
  if(json_err != nil){
    return addressSearchIpData, err
  }
  if addressSearchIpData.Code != 200 {
    return addressSearchIpData, errors.New(addressSearchIpData.Message)
  }
  return addressSearchIpData, nil
}

func PatchUpdateAddress(server_url string, application string, newHostname string, addressId string, token string) (*UpdateAddress, error) {
  var updateAddressData = new(UpdateAddress)
  client := &http.Client{}
  reqBody := "hostname=" + newHostname
  req, err := http.NewRequest("PATCH", "https://" + server_url + "/api/" + application + "/addresses/" + addressId + "/", strings.NewReader(reqBody))
  req.Header.Add("token", token)
  resp, err := client.Do(req)
  if (err!=nil) {
    return updateAddressData, err
  }
  body, err := ioutil.ReadAll(resp.Body)
  if (err!=nil) {
    return updateAddressData, err
  }
  json_err := json.Unmarshal([]byte(body), &updateAddressData)
  if(json_err != nil){
    return updateAddressData, err
  }
  if updateAddressData.Code != 200 {
    return updateAddressData, errors.New(updateAddressData.Message)
  }
  return updateAddressData, nil
}
