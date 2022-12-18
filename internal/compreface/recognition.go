package compreface

import (
	"image"
)

type RecognitionService struct {
	Subjects *SubjectService
	client   *ComprefaceClient
}

type RecognitionResponse struct {
	Result []struct {
		Age       []int         `json:"age"`
		Gender    string        `json:"gender"`
		Embedding []interface{} `json:"embedding"`
		Box       struct {
			Probability float64 `json:"probability"`
			XMax        int     `json:"x_max"`
			YMax        int     `json:"y_max"`
			XMin        int     `json:"x_min"`
			YMin        int     `json:"y_min"`
		} `json:"box"`
		Landmarks [][]int `json:"landmarks"`
		Subjects  []struct {
			Similarity float64 `json:"similarity"`
			Subject    string  `json:"subject"`
		} `json:"subjects"`
		ExecutionTime struct {
			Age        float64 `json:"age"`
			Gender     float64 `json:"gender"`
			Detector   float64 `json:"detector"`
			Calculator float64 `json:"calculator"`
		} `json:"execution_time"`
	} `json:"result"`
	PluginsVersions struct {
		Age        string `json:"age"`
		Gender     string `json:"gender"`
		Detector   string `json:"detector"`
		Calculator string `json:"calculator"`
	} `json:"plugins_versions"`
}

func InitRecognition(client *ComprefaceClient) *RecognitionService {
	recognition := &RecognitionService{
		client: client,
	}

	recognition.Subjects = &SubjectService{client: client}

	return recognition
}

func (r *RecognitionService) RecognizeFaces(img image.Image) (*RecognitionResponse, error) {
	req, err := r.client.NewRequest("POST", "/api/v1/recognition/recognize", img)
	if err != nil {
		return nil, err
	}

	var response RecognitionResponse
	_, err = req.Do(&response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
