package phpipam

import (
  "net/http"
  "encoding/json"
  "strings"
  "errors"
)

type Address struct {
  Code    int `json:"code"`
  Success bool `json:"success"`
  Data struct {
    Id        string `json:"id"`
    SubnetId  string `json:"subnetId"`
    Ip        string `json:"ip"`
    Hostname  string `json:"hostname"`
  } `json:"data"`
  Message string `json:"message"`
}

type AddressSearch struct {
  Code    int `json:"code"`
  Success bool `json:"success"`
  Data []struct {
    Id          string `json:"id"`
    SubnetId    string `json:"subnetId"`
    Ip          string `json:"ip"`
    Description string `json:"description"`
  } `json:"data"`
  Message string `json:"message"`
}

type AddressPing struct {
  Code    int `json:"code"`
  Success bool `json:"success"`
  Data struct {
    ScanType  string `json:"scan_type"`
    ExitCode  int `json:"exit_code"`
  } `json:"data"`
  Message string `json:"message"`
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
  Data []struct {
    Id          string `json:"id"`
    SubnetId    string `json:"subnetId"`
    Ip          string `json:"ip"`
    Hostname    string `json:"hostname"`
  }`json:"data"`
  Message string `json:"message"`
}

type UpdateAddress struct {
  Code    int `json:"code"`
  Success bool `json:"success"`
  Message string `json:"message"`
}

func (c *Client) GetAddress(addressId string) (Address, error) {
  var addressData Address
  req, _ := http.NewRequest("GET", "https://" + c.ServerUrl + "/api/" + c.Application + "/addresses/" + addressId + "/", nil)
  body, err := c.Do(req)
  if err != nil {
    return addressData, err
  }
  err = json.Unmarshal([]byte(body), &addressData)
  if err != nil {
    return addressData, err
  }
  if addressData.Code == 200 || addressData.Code == 404 {
    // Accepted responses
  } else {
    return addressData, errors.New(addressData.Message)
  }
  return addressData, nil
}

func (c *Client) GetAddressSearch(searchHostname string) (AddressSearch, error) {
  var addressSearchData AddressSearch
  req, _ := http.NewRequest("GET", "https://" + c.ServerUrl + "/api/" + c.Application + "/addresses/search_hostname/" + searchHostname + "/", nil)
  body, err := c.Do(req)
  if err != nil {
    return addressSearchData, err
  }
  err = json.Unmarshal([]byte(body), &addressSearchData)
  if err != nil {
    return addressSearchData, err
  }
  if addressSearchData.Code == 200 || addressSearchData.Code == 404 {
    // Accepted responses
  } else {
    return addressSearchData, errors.New(addressSearchData.Message)
  }
  return addressSearchData, nil
}

func (c *Client) GetAddressPing(addressId string) (AddressPing, error) {
  var addressPingData AddressPing
  req, _ := http.NewRequest("GET", "https://" + c.ServerUrl + "/api/" + c.Application + "/addresses/" + addressId + "/ping/", nil)
  body, err := c.Do(req)
  if err != nil {
    return addressPingData, err
  }
  err = json.Unmarshal([]byte(body), &addressPingData)
  if err != nil{
    return addressPingData, err
  }
  if addressPingData.Code == 200 || addressPingData.Code == 404 {
    // Accepted responses
  } else {
    return addressPingData, errors.New(addressPingData.Message)
  }
  return addressPingData, nil
}

func (c *Client) DeleteAddress(addressId string) (AddressDelete, error) {
  var addressDeleteData AddressDelete
  req, _ := http.NewRequest("DELETE", "https://" + c.ServerUrl + "/api/" + c.Application + "/addresses/" + addressId + "/", nil)
  body, err := c.Do(req)
  if err != nil {
    return addressDeleteData, err
  }
  err = json.Unmarshal([]byte(body), &addressDeleteData)
  if err != nil {
    return addressDeleteData, err
  }
  if addressDeleteData.Code != 200 {
    return addressDeleteData, errors.New(addressDeleteData.Message)
  }
  return addressDeleteData, nil
}

func (c *Client) CreateAddressFirstFree(subnetId string, hostname string, owner string) (AddressFirstFree, error) {
  var addressFirstFreeData AddressFirstFree
  reqBody := "hostname=" + hostname + "&owner=" + owner
  req, _ := http.NewRequest("POST", "https://" + c.ServerUrl + "/api/" + c.Application + "/addresses/first_free/" + subnetId + "/", strings.NewReader(reqBody))
  body, err := c.Do(req)
  if err != nil {
    return addressFirstFreeData, err
  }
  err = json.Unmarshal([]byte(body), &addressFirstFreeData)
  if err != nil {
    return addressFirstFreeData, err
  }
  if addressFirstFreeData.Code != 201 {
    return addressFirstFreeData, errors.New(addressFirstFreeData.Message)
  }
  return addressFirstFreeData, nil
}

func (c *Client) GetAddressSearchIp(address string) (AddressSearchIp, error) {
  var addressSearchIpData AddressSearchIp
  req, _ := http.NewRequest("GET", "https://" + c.ServerUrl + "/api/" + c.Application + "/addresses/search/" + address + "/", nil)
  body, err := c.Do(req)
  if err != nil {
    return addressSearchIpData, err
  }
  err = json.Unmarshal([]byte(body), &addressSearchIpData)
  if err != nil {
    return addressSearchIpData, err
  }
  if addressSearchIpData.Code != 200 {
    return addressSearchIpData, errors.New(addressSearchIpData.Message)
  }
  return addressSearchIpData, nil
}

func (c *Client) PatchUpdateAddress(hostname string, addressId string) (UpdateAddress, error) {
  var updateAddressData UpdateAddress
  reqBody := "hostname=" + hostname
  req, _ := http.NewRequest("PATCH", "https://" + c.ServerUrl + "/api/" + c.Application + "/addresses/" + addressId + "/", strings.NewReader(reqBody))
  body, err := c.Do(req)
  if err != nil {
    return updateAddressData, err
  }
  err = json.Unmarshal([]byte(body), &updateAddressData)
  if err != nil {
    return updateAddressData, err
  }
  if updateAddressData.Code != 200 {
    return updateAddressData, errors.New(updateAddressData.Message)
  }
  return updateAddressData, nil
}
