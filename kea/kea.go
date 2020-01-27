package kea

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"log"
)

type Client struct {
	baseURL string
}

func NewClient(baseURL string) *Client {
	return &Client{
		baseURL: baseURL,
	}
}

type ResourceReservation struct {
}

type GetConfigCommand struct {
	Command string   `json:"command"`
	Service []string `json:"service"`
}

func (s *Client) CreateReservation(saveHost *SaveHost) (*Host, error) {
	host := Host{}

	err := s.MakeRequest(saveHost, &host, nil, s.baseURL+"/host", "POST")
	if err != nil {
		return nil, err
	}

	log.Printf("[DEBUG] host response : %v", host)
	return &host, nil
}

func (s *Client) GetReservationById(id int64) (*Host, error) {
	host := Host{}
	log.Printf("[DEBUG] fetch host with id : %d", id)
	err := s.MakeRequest(nil, &host, nil, s.baseURL+"/host/"+strconv.FormatInt(id, 10), "GET")
	if err != nil {
		return nil, err
	}
	log.Printf("[DEBUG] host response : %v", host)
	return &host, nil
}

func (s *Client) GetReservationByHwAddress(hwAddress string) (*Host, error) {
	host := []Host{}
	err := s.MakeRequest(nil, &host, map[string]string{
		"identifier":     hwAddress,
		"identifierType": HwAddress,
	}, s.baseURL+"/host", "GET")
	if err != nil {
		return nil, err
	}

	if len(host) != 1 {
		return nil, nil
	}
	log.Printf("[DEBUG] host response : %v", host[0])
	return &host[0], nil
}

func (s *Client) MakeRequest(reqBody interface{}, respBody interface{}, query map[string]string, url string, method string) error {
	j, err := json.Marshal(reqBody)
	if err != nil {
		return err
	}
	log.Printf("[TRACE] body req: %s", bytes.NewBuffer(j))
	var req *http.Request

	if j == nil {
		req, err = http.NewRequest(method, url, nil)
	} else {
		req, err = http.NewRequest(method, url, bytes.NewBuffer(j))
	}
	if err != nil {
		log.Printf("[ERROR] could not build request")
		return err
	}
	req.Header.Add("Content-Type", "application/json")

	if query != nil {
		var urlQuery = req.URL.Query()
		for k, v := range query {
			urlQuery.Add(k, v)
		}
		log.Printf("[TRACE] setting url query : %s", urlQuery.Encode())
		req.URL.RawQuery = urlQuery.Encode()
	}

	log.Printf("[TRACE] calling %s with method %s", url, method)
	resp, err := s.doRequest(req)
	log.Printf("[DEBUG] response body: %s", resp)
	if err := json.Unmarshal(resp, &respBody); err != nil {
		return err
	}

	return nil
}

func (s *Client) GetConfig() (*Host, error) {
	reqBody := &GetConfigCommand{
		"config-get",
		[]string{"dhcp4"},
	}
	j, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}
	log.Printf("[TRACE] body req: %s", bytes.NewBuffer(j))
	req, err := http.NewRequest("POST", s.baseURL, bytes.NewBuffer(j))
	req.Header.Add("Content-Type", "application/json")
	if err != nil {
		return nil, err
	}
	resp, err := s.doRequest(req)
	var tmp *Host
	if err := json.Unmarshal(resp, &tmp); err != nil {
		return nil, err
	}

	return tmp, nil
}

func (s *Client) doRequest(req *http.Request) ([]byte, error) {
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if 200 != resp.StatusCode {
		return nil, fmt.Errorf("%s", body)
	}
	return body, nil
}
