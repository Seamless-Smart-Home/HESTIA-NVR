package frigateHTTP

import (
	"bytes"
	"encoding/json"
	"image"
	"io"
	"net/http"
	"net/url"
)

type Client struct {
	Config *ConfigService
	Camera *CameraService

	baseURL    *url.URL
	userAgent  string
	httpClient *http.Client
}

func NewClient(frigateURL string, httpClient *http.Client) (*Client, error) {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	//Create Base URL
	baseURL, err := url.Parse(frigateURL)
	if err != nil {
		return nil, err
	}

	client := &Client{
		baseURL:    baseURL,
		userAgent:  "HESTIA",
		httpClient: httpClient,
	}

	client.Config = &ConfigService{client: client}
	client.Camera = &CameraService{client: client}

	return client, nil
}

func (c *Client) NewRequest(method, path string, body interface{}) (*http.Request, error) {
	rel := &url.URL{Path: path}
	u := c.baseURL.ResolveReference(rel)
	var buf io.ReadWriter
	if body != nil {
		buf = new(bytes.Buffer)
		err := json.NewEncoder(buf).Encode(body)
		if err != nil {
			return nil, err
		}
	}
	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		return nil, err
	}
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	req.Header.Set("User-Agent", c.userAgent)
	return req, nil
}

func (c *Client) Do(req *http.Request, v interface{}) (*http.Response, error) {
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(v)
	return resp, err
}

func (c *Client) DoImage(req *http.Request) (image.Image, error) {
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	//Convert http response to Image
	snapshot, _, err := image.Decode(resp.Body)
	if err != nil {
		return nil, err
	}

	resp.Body.Close()

	return snapshot, nil
}
