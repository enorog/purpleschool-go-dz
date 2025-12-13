package api

import (
	"cli/bins"
	"cli/config"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

const ROOT = "https://api.jsonbin.io/v3/b"

type JsonBin struct {
	key string
}

func NewJsonBin(config *config.Config) *JsonBin {
	return &JsonBin{
		key: config.Key,
	}
}

type CreateResponse struct {
	Bin *bins.Bin `json:"metadata"`
}

type ErrorResponse struct {
	Message string `json:"message"`
}

func (jsonBin *JsonBin) Create(name string, content io.Reader) (*bins.Bin, error) {
	req, err := http.NewRequest(http.MethodPost, ROOT, content)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Master-Key", jsonBin.key)
	req.Header.Set("X-Bin-Name", name)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == 400 || resp.StatusCode == 401 || resp.StatusCode == 403 || resp.StatusCode == 404 {
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		var errResponse *ErrorResponse
		err = json.Unmarshal(body, &errResponse)
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("ошибка: status=%d, message=%s", resp.StatusCode, errResponse.Message)
	} else if resp.StatusCode != 200 {
		return nil, fmt.Errorf("ошибка: status=%d", resp.StatusCode)
	}
	body, err := io.ReadAll(resp.Body)
	var createResponse *CreateResponse
	err = json.Unmarshal(body, &createResponse)
	if err != nil {
		return nil, err
	}
	return createResponse.Bin, nil
}

func (jsonBin *JsonBin) Update(id string, content io.Reader) error {
	req, err := http.NewRequest(http.MethodPut, ROOT+"/"+url.PathEscape(id), content)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Master-Key", jsonBin.key)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode == 400 || resp.StatusCode == 401 || resp.StatusCode == 403 || resp.StatusCode == 404 {
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		var errResponse *ErrorResponse
		err = json.Unmarshal(body, &errResponse)
		if err != nil {
			return err
		}
		return fmt.Errorf("ошибка: status=%d, message=%s", resp.StatusCode, errResponse.Message)
	} else if resp.StatusCode != 200 {
		return fmt.Errorf("ошибка: status=%d", resp.StatusCode)
	}

	return nil
}

func (jsonBin *JsonBin) Get(id string) ([]byte, error) {
	req, err := http.NewRequest(http.MethodGet, ROOT+"/"+url.PathEscape(id), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("X-Master-Key", jsonBin.key)
	queryParams := req.URL.Query()
	queryParams.Set("meta", "false")
	req.URL.RawQuery = queryParams.Encode()
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == 400 || resp.StatusCode == 401 || resp.StatusCode == 403 || resp.StatusCode == 404 {
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		var errResponse *ErrorResponse
		err = json.Unmarshal(body, &errResponse)
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("ошибка: status=%d, message=%s", resp.StatusCode, errResponse.Message)
	} else if resp.StatusCode != 200 {
		return nil, fmt.Errorf("ошибка: status=%d", resp.StatusCode)
	}
	body, err := io.ReadAll(resp.Body)

	return body, nil
}

func (jsonBin *JsonBin) Delete(id string) error {
	req, err := http.NewRequest(http.MethodDelete, ROOT+"/"+url.PathEscape(id), nil)
	if err != nil {
		return err
	}
	req.Header.Set("X-Master-Key", jsonBin.key)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode == 400 || resp.StatusCode == 401 || resp.StatusCode == 403 || resp.StatusCode == 404 {
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		var errResponse *ErrorResponse
		err = json.Unmarshal(body, &errResponse)
		if err != nil {
			return err
		}
		return fmt.Errorf("ошибка: status=%d, message=%s", resp.StatusCode, errResponse.Message)
	} else if resp.StatusCode != 200 {
		return fmt.Errorf("ошибка: status=%d", resp.StatusCode)
	}

	return nil
}
