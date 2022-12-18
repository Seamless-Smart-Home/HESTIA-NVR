package frigateHTTP

import "image"

type CameraService struct {
	client *Client
}

type CameraResponse struct {
	Birdseye BirdseyeConfig          `json:"birdseye"`
	Cameras  map[string]CameraConfig `json:"cameras"`
}

func (c *CameraService) GetLatestImage(name string) (image.Image, error) {
	req, err := c.client.NewRequest("GET", "/api/"+name+"/latest.jpg", nil)
	if err != nil {
		return nil, err
	}
	q := req.URL.Query()
	q.Add("h", "720")
	req.URL.RawQuery = q.Encode()

	img, err := c.client.DoImage(req)
	if err != nil {
		return nil, err
	}

	return img, nil
}
