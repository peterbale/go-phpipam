package phpipam

import (
	"encoding/json"
	"errors"
	"net/http"
)

// Sections struct from phpipam
type Sections struct {
	Code    int  `json:"code"`
	Success bool `json:"success"`
	Data    []struct {
		ID          string `json:"id"`
		Name        string `json:"name"`
		Description string `json:"description"`
	} `json:"data"`
	Message string `json:"message"`
}

// Section struct from phpipam
type Section struct {
	Code    int  `json:"code"`
	Success bool `json:"success"`
	Data    struct {
		ID          string `json:"id"`
		Name        string `json:"name"`
		Description string `json:"description"`
	} `json:"data"`
	Message string `json:"message"`
}

// SectionsSubnets struct from phpipam
type SectionsSubnets struct {
	Code    int  `json:"code"`
	Success bool `json:"success"`
	Data    []struct {
		ID          string `json:"id"`
		Subnet      string `json:"subnet"`
		Description string `json:"description"`
	} `json:"data"`
	Message string `json:"message"`
}

// GetSections Client pointer method to get all sections from phpipam, returns
// Sections struct and error
func (c *Client) GetSections() (Sections, error) {
	var sectionsData Sections
	req, _ := http.NewRequest("GET", "https://"+c.ServerURL+"/api/"+c.Application+"/sections/", nil)
	body, err := c.Do(req)
	if err != nil {
		return sectionsData, err
	}
	err = json.Unmarshal([]byte(body), &sectionsData)
	if err != nil {
		return sectionsData, err
	}
	if sectionsData.Code != 200 {
		return sectionsData, errors.New(sectionsData.Message)
	}
	return sectionsData, nil
}

// GetSection Client pointer method to get information about a specific section
// using sectionID string, returns Section struct and error
func (c *Client) GetSection(sectionID string) (Section, error) {
	var sectionData Section
	req, _ := http.NewRequest("GET", "https://"+c.ServerURL+"/api/"+c.Application+"/sections/"+sectionID+"/", nil)
	body, err := c.Do(req)
	if err != nil {
		return sectionData, err
	}
	err = json.Unmarshal([]byte(body), &sectionData)
	if err != nil {
		return sectionData, err
	}
	if sectionData.Code != 200 {
		return sectionData, errors.New(sectionData.Message)
	}
	return sectionData, nil
}

// GetSectionsSubnets Client pointer method to get all subnets within a specific
// phpipam section using sectionID string, returns SectionsSubnets struct and
// error
func (c *Client) GetSectionsSubnets(sectionID string) (SectionsSubnets, error) {
	var sectionsSubnetsData SectionsSubnets
	req, _ := http.NewRequest("GET", "https://"+c.ServerURL+"/api/"+c.Application+"/sections/"+sectionID+"/subnets/", nil)
	body, err := c.Do(req)
	if err != nil {
		return sectionsSubnetsData, err
	}
	err = json.Unmarshal([]byte(body), &sectionsSubnetsData)
	if err != nil {
		return sectionsSubnetsData, err
	}
	if sectionsSubnetsData.Code != 200 {
		return sectionsSubnetsData, errors.New(sectionsSubnetsData.Message)
	}
	return sectionsSubnetsData, nil
}
