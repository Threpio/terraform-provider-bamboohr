package clients

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// HostURL - Default BambooHR Host URL
const HostURL string = "https://api.bamboohr.com/api/gateway.php/"

type Client struct {
	HostURL    string
	HTTPClient *http.Client
	Auth       string
}

func NewClient(subdomain, apiKey string) (*Client, error) {

	c := Client{
		HTTPClient: &http.Client{Timeout: 10 * time.Second},
		HostURL:    fmt.Sprintf("%s%s/v1", HostURL, subdomain),
		Auth:       fmt.Sprintf("Basic %s", base64.StdEncoding.EncodeToString([]byte(apiKey+":x"))),
	}

	return &c, nil
}

func (c *Client) DoRequest(req *http.Request) ([]byte, error) {
	// Set standard headers
	req.Header.Set("Authorization", c.Auth)
	req.Header.Set("Accept", "application/json")

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if res.StatusCode < http.StatusOK || res.StatusCode >= http.StatusBadRequest {
		return body, fmt.Errorf("error from api, status code: %d", res.StatusCode)
	}

	return body, err
}
