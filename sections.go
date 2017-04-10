package phpipam

import (
	"encoding/json"
	"errors"
	"net/http"
)

type Sections struct {
	Code    int  `json:"code"`
	Success bool `json:"success"`
	Data    []struct {
		Id          string `json:"id"`
		Name        string `json:"name"`
		Description string `json:"description"`
	} `json:"data"`
	Message string `json:"message"`
}

type Section struct {
	Code    int  `json:"code"`
	Success bool `json:"success"`
	Data    struct {
		Id          string `json:"id"`
		Name        string `json:"name"`
		Description string `json:"description"`
	} `json:"data"`
	Message string `json:"message"`
}

type SectionsSubnets struct {
	Code    int  `json:"code"`
	Success bool `json:"success"`
	Data    []struct {
		Id          string `json:"id"`
		Subnet      string `json:"subnet"`
		Description string `json:"description"`
	} `json:"data"`
	Message string `json:"message"`
}

func (c *Client) GetSections() (Sections, error) {
	var sectionsData Sections
	req, _ := http.NewRequest("GET", "https://"+c.ServerUrl+"/api/"+c.Application+"/sections/", nil)
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

func (c *Client) GetSection(sectionId string) (Section, error) {
	var sectionData Section
	req, _ := http.NewRequest("GET", "https://"+c.ServerUrl+"/api/"+c.Application+"/sections/"+sectionId+"/", nil)
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

func (c *Client) GetSectionsSubnets(sectionId string) (SectionsSubnets, error) {
	var sectionsSubnetsData SectionsSubnets
	req, _ := http.NewRequest("GET", "https://"+c.ServerUrl+"/api/"+c.Application+"/sections/"+sectionId+"/subnets/", nil)
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
