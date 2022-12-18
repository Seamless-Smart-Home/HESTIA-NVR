package frigateHTTP

type ConfigService struct {
	client *Client
}

type ConfigResponse struct {
	Birdseye BirdseyeConfig          `json:"birdseye"`
	Cameras  map[string]CameraConfig `json:"cameras"`
}

type BirdseyeConfig struct {
	Enabled bool   `json:"enabled"`
	Height  int    `json:"height"`
	Mode    string `json:"mode"`
	Quality int    `json:"quality"`
	Width   int    `json:"width"`
}

type CameraConfig struct {
	BestImageTimeout int            `json:"best_image_timeout"`
	Birdseye         BirdseyeConfig `json:"birdseye"`
	Name             string         `json:"name"`
}

func (c ConfigService) Get() (*ConfigResponse, error) {
	req, err := c.client.NewRequest("GET", "/api/config", nil)
	if err != nil {
		return nil, err
	}

	var config ConfigResponse
	_, err = c.client.Do(req, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
