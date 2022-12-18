package compreface

import (
	"image"
)

type ExampleService struct {
	client *ComprefaceClient
}

type ExampleListResponse struct {
	Faces         []ExampleFacesResponse `json:"faces"`
	PageNumber    int                    `json:"page_number"`
	PageSize      int                    `json:"page_size"`
	TotalPages    int                    `json:"total_pages"`
	TotalElements int                    `json:"total_elements"`
}

type ExampleFacesResponse struct {
	ImageId string `json:"image_id"`
	Subject string `json:"subject"`
}

func (s *ExampleService) List() (*ExampleListResponse, error) {
	req, err := s.client.NewRequest("GET", "/api/v1/recognition/faces", nil)
	if err != nil {
		return nil, err
	}

	var examples ExampleListResponse
	_, err = req.Do(&examples)
	if err != nil {
		return nil, err
	}

	return &examples, nil
}

func (s *ExampleService) GetImage(imageID string) (image.Image, error) {
	req, err := s.client.NewRequest("GET", "/api/v1/recognition/faces/"+imageID+"/img", nil)
	if err != nil {
		return nil, err
	}

	img, err := req.DoImage()
	if err != nil {
		return nil, err
	}

	return img, nil
}
