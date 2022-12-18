package compreface

import (
	"bytes"
	"context"
	"encoding/json"
	"image"
	"image/jpeg"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"time"
)

type ComprefaceClient struct {
	Recognition *RecognitionService
	baseURL     *url.URL
	apiKey      string
	userAgent   string
	httpClient  *http.Client
}

type ComprefaceRequest struct {
	req         *http.Request
	withTimeout bool
	client      *ComprefaceClient
}

var Client *ComprefaceClient

func NewClient(comprefaceURL string, comprefaceApiKey string, httpClient *http.Client) error {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	//Create Base URL
	baseURL, err := url.Parse(comprefaceURL)
	if err != nil {
		return err
	}

	Client = &ComprefaceClient{
		baseURL:    baseURL,
		apiKey:     comprefaceApiKey,
		userAgent:  "HESTIA",
		httpClient: httpClient,
	}

	Client.Recognition = InitRecognition(Client)

	return nil
}

func (c *ComprefaceClient) NewRequest(method, path string, body interface{}) (*ComprefaceRequest, error) {
	comprefaceRequest := &ComprefaceRequest{
		withTimeout: false,
		client:      c,
	}

	rel := &url.URL{Path: path}
	u := c.baseURL.ResolveReference(rel)
	var buf io.ReadWriter
	var contentType string
	if body != nil {
		switch body.(type) {
		case *image.RGBA:
			buf = &bytes.Buffer{}
			writer := multipart.NewWriter(buf)
			fw, err := writer.CreateFormFile("file", "image.jpg")
			if err != nil {
				return nil, err
			}

			err = jpeg.Encode(fw, body.(image.Image), nil)
			if err != nil {
				return nil, err
			}
			writer.Close()
			contentType = writer.FormDataContentType()
			comprefaceRequest.withTimeout = true
		default:
			buf = new(bytes.Buffer)
			err := json.NewEncoder(buf).Encode(body)
			if err != nil {
				return nil, err
			}
			contentType = "application/json"
		}

	}
	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		return nil, err
	}
	req.Header.Set("x-api-key", c.apiKey)
	req.Header.Set("User-Agent", c.userAgent)

	if body != nil {
		req.Header.Set("Content-Type", contentType)
	}

	comprefaceRequest.req = req

	return comprefaceRequest, nil
}

func (r *ComprefaceRequest) Do(v interface{}) (*http.Response, error) {
	var ctx context.Context
	if r.withTimeout {
		withTimeout, cancel := context.WithTimeout(context.Background(), time.Second*20)
		defer cancel()
		ctx = withTimeout
	} else {
		ctx = context.Background()
	}

	resp, err := r.client.httpClient.Do(r.req.WithContext(ctx))
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(v)
	return resp, err
}

func (r *ComprefaceRequest) DoImage() (image.Image, error) {
	resp, err := r.client.httpClient.Do(r.req)
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
