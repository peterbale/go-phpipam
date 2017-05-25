package phpipam

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"
)

// Address struct from phpipam
type Address struct {
	Code    int  `json:"code"`
	Success bool `json:"success"`
	Data    struct {
		ID       string `json:"id"`
		SubnetID string `json:"subnetId"`
		IP       string `json:"ip"`
		Hostname string `json:"hostname"`
	} `json:"data"`
	Message string `json:"message"`
}

// AddressSearch struct from phpipam
type AddressSearch struct {
	Code    int  `json:"code"`
	Success bool `json:"success"`
	Data    []struct {
		ID          string `json:"id"`
		SubnetID    string `json:"subnetId"`
		IP          string `json:"ip"`
		Description string `json:"description"`
	} `json:"data"`
	Message string `json:"message"`
}

// AddressPing struct from phpipam
type AddressPing struct {
	Code    int  `json:"code"`
	Success bool `json:"success"`
	Data    struct {
		ScanType string `json:"scan_type"`
		ExitCode int    `json:"exit_code"`
	} `json:"data"`
	Message string `json:"message"`
}

// AddressDelete struct from phpipam
type AddressDelete struct {
	Code    int    `json:"code"`
	Success bool   `json:"success"`
	Message string `json:"message"`
}

// AddressFirstFree struct from phpipam
type AddressFirstFree struct {
	Code    int    `json:"code"`
	Success bool   `json:"success"`
	Message string `json:"message"`
	IP      string `json:"ip"`
}

// AddressSearchIP struct from phpipam
type AddressSearchIP struct {
	Code    int  `json:"code"`
	Success bool `json:"success"`
	Data    []struct {
		ID       string `json:"id"`
		SubnetID string `json:"subnetId"`
		IP       string `json:"ip"`
		Hostname string `json:"hostname"`
	} `json:"data"`
	Message string `json:"message"`
}

// UpdateAddress struct from phpipam
type UpdateAddress struct {
	Code    int    `json:"code"`
	Success bool   `json:"success"`
	Message string `json:"message"`
}

// GetAddress Client pointer method to get phpipam address using addressID
// string, returns Address struct and error
func (c *Client) GetAddress(addressID string) (Address, error) {
	var addressData Address
	req, _ := http.NewRequest("GET", "https://"+c.ServerURL+"/api/"+c.Application+"/addresses/"+addressID+"/", nil)
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

// GetAddressSearch Client pointer method to search for phpiapm address using
// hostname string, returns AddressSearch struct and error
func (c *Client) GetAddressSearch(searchHostname string) (AddressSearch, error) {
	var addressSearchData AddressSearch
	req, _ := http.NewRequest("GET", "https://"+c.ServerURL+"/api/"+c.Application+"/addresses/search_hostname/"+searchHostname+"/", nil)
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

// GetAddressPing Client pointer method to get ping status about phpipam address
// using addressID string, returns AddressPring struct and error
func (c *Client) GetAddressPing(addressID string) (AddressPing, error) {
	var addressPingData AddressPing
	req, _ := http.NewRequest("GET", "https://"+c.ServerURL+"/api/"+c.Application+"/addresses/"+addressID+"/ping/", nil)
	body, err := c.Do(req)
	if err != nil {
		return addressPingData, err
	}
	err = json.Unmarshal([]byte(body), &addressPingData)
	if err != nil {
		return addressPingData, err
	}
	if addressPingData.Code == 200 || addressPingData.Code == 404 {
		// Accepted responses
	} else {
		return addressPingData, errors.New(addressPingData.Message)
	}
	return addressPingData, nil
}

// DeleteAddress Client pointer method to delete a phpipam address using
// addressID string, returns AddressDelete struct and error
func (c *Client) DeleteAddress(addressID string) (AddressDelete, error) {
	var addressDeleteData AddressDelete
	req, _ := http.NewRequest("DELETE", "https://"+c.ServerURL+"/api/"+c.Application+"/addresses/"+addressID+"/", nil)
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

// CreateAddressFirstFree Client pointer method to create the first avalible
// phpipam address (starting from the top of the subnet) using subnetID string,
// hostname string and owner string, returns AddressFirstFree struct and error
func (c *Client) CreateAddressFirstFree(subnetID string, hostname string, owner string) (AddressFirstFree, error) {
	var addressFirstFreeData AddressFirstFree
	reqBody := "hostname=" + hostname + "&owner=" + owner
	req, _ := http.NewRequest("POST", "https://"+c.ServerURL+"/api/"+c.Application+"/addresses/first_free/"+subnetID+"/", strings.NewReader(reqBody))
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

// GetAddressSearchIP Client pointer method to search for phpipam address using
// address string, returns AddressSearchIP struct and error
func (c *Client) GetAddressSearchIP(address string) (AddressSearchIP, error) {
	var addressSearchIPData AddressSearchIP
	req, _ := http.NewRequest("GET", "https://"+c.ServerURL+"/api/"+c.Application+"/addresses/search/"+address+"/", nil)
	body, err := c.Do(req)
	if err != nil {
		return addressSearchIPData, err
	}
	err = json.Unmarshal([]byte(body), &addressSearchIPData)
	if err != nil {
		return addressSearchIPData, err
	}
	if addressSearchIPData.Code != 200 {
		return addressSearchIPData, errors.New(addressSearchIPData.Message)
	}
	return addressSearchIPData, nil
}

// PatchUpdateAddress Client pointer method to be able to update a phpipam
// address to a new hostname using hostname string and addressID string, returns
// UpdateAddress struct and error
func (c *Client) PatchUpdateAddress(hostname string, addressID string) (UpdateAddress, error) {
	var updateAddressData UpdateAddress
	reqBody := "hostname=" + hostname
	req, _ := http.NewRequest("PATCH", "https://"+c.ServerURL+"/api/"+c.Application+"/addresses/"+addressID+"/", strings.NewReader(reqBody))
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
