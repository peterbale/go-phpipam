package phpipam

import (
	"encoding/json"
	"errors"
	"net/http"
)

type Subnet struct {
	Code    int  `json:"code"`
	Success bool `json:"success"`
	Data    struct {
		Id          string `json:"id"`
		Subnet      string `json:"subnet"`
		Mask        string `json:"mask"`
		SectionId   string `json:"sectionId"`
		Description string `json:"description"`
		IsFull      string `json:"isFull"`
		Gateway     struct {
			IPAddress string `json:"ip_addr"`
		} `json:"gateway"`
		Calculation struct {
			Type          string `json:"Type"`
			IPAddress     string `json:"IP Address"`
			Network       string `json:"Network"`
			Broadcast     string `json:"Broadcast"`
			BitMask       string `json:"Subnet bitmask"`
			NumberOfHosts int    `json:"Number of hosts"`
		} `json:"calculation"`
	} `json:"data"`
	Message string `json:"message"`
}

func (c *Client) GetSubnet(subnetId string) (Subnet, error) {
	var subnetData Subnet
	req, _ := http.NewRequest("GET", "https://"+c.ServerUrl+"/api/"+c.Application+"/subnets/"+subnetId+"/", nil)
	body, err := c.Do(req)
	if err != nil {
		return subnetData, err
	}
	err = json.Unmarshal([]byte(body), &subnetData)
	if err != nil {
		return subnetData, err
	}
	if subnetData.Code != 200 {
		return subnetData, errors.New(subnetData.Message)
	}
	return subnetData, nil
}
