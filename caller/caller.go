package caller

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

const baseServiceUrl = "http://159.65.52.236/api"

var (
	ErrInvalidEndpoint error = errors.New("Add prefix / to endpoint")
)

type service struct {
	token      string
	baseURL    string
	apiVersion int
}

type Service interface {
	Get(endpoint string, response interface{}) error
	Post(endpoint string, request, response interface{}) error
	Delete(endpoint string, response interface{}) error
	Update(endpoint string, request, response interface{}) error
}

func New(token string) Service {
	return &service{
		token:      token,
		baseURL:    baseServiceUrl,
		apiVersion: 1,
	}
}

func (s *service) Get(endpoint string, response interface{}) error {
	if endpoint[0] != '/' {
		return ErrInvalidEndpoint
	}

	client := &http.Client{}
	url := fmt.Sprintf("%s/v%d%s", s.baseURL, s.apiVersion, endpoint)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", s.token))
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, response)
	if err != nil {
		return err
	}

	return nil
}

func (s *service) Post(endpoint string, request, response interface{}) error {
	if endpoint[0] != '/' {
		return ErrInvalidEndpoint
	}

	client := &http.Client{}
	url := fmt.Sprintf("%s/v%d%s", s.baseURL, s.apiVersion, endpoint)

	b, err := json.Marshal(request)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(b))
	if err != nil {
		return err
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", s.token))
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, response)
	if err != nil {
		return err
	}

	return nil
}

func (s *service) Delete(endpoint string, response interface{}) error {
	if endpoint[0] != '/' {
		return ErrInvalidEndpoint
	}

	client := &http.Client{}
	url := fmt.Sprintf("%s/v%d%s", s.baseURL, s.apiVersion, endpoint)

	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return err
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", s.token))
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, response)
	if err != nil {
		return err
	}

	return nil
}

func (s *service) Update(endpoint string, request, response interface{}) error {
	if endpoint[0] != '/' {
		return ErrInvalidEndpoint
	}

	client := &http.Client{}
	url := fmt.Sprintf("%s/v%d%s", s.baseURL, s.apiVersion, endpoint)

	b, err := json.Marshal(request)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("PATCH", url, bytes.NewBuffer(b))
	if err != nil {
		return err
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", s.token))
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, response)
	if err != nil {
		return err
	}

	return nil
}
