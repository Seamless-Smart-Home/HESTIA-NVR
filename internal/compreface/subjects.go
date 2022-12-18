package compreface

import (
	"errors"
)

type SubjectService struct {
	client *ComprefaceClient
}

type SubjectListResponse struct {
	Subjects []string `json:"subjects"`
}

type Subject struct {
	Subject string `json:"subject"`
}

type UpdateResponse struct {
	Updated bool `json:"updated"`
}

type DeletedResponse struct {
	Deleted int `json:"deleted"`
}

func (s *SubjectService) List() (*SubjectListResponse, error) {
	req, err := s.client.NewRequest("GET", "/api/v1/recognition/subjects", nil)
	if err != nil {
		return nil, err
	}

	var subjects SubjectListResponse
	_, err = req.Do(&subjects)
	if err != nil {
		return nil, err
	}

	return &subjects, nil
}

func (s *SubjectService) Add(name string) error {
	subject := &Subject{name}
	req, err := s.client.NewRequest("POST", "/api/v1/recognition/subjects", subject)
	if err != nil {
		return err
	}

	var subjects SubjectListResponse
	_, err = req.Do(&subjects)
	if err != nil {
		return err
	}

	return nil
}

func (s *SubjectService) Rename(currentName string, newName string) error {
	subject := &Subject{newName}
	req, err := s.client.NewRequest("PUT", "/api/v1/recognition/subjects/"+currentName, subject)
	if err != nil {
		return err
	}

	var response UpdateResponse
	_, err = req.Do(&response)
	if err != nil {
		return err
	}

	if !response.Updated {
		return errors.New("error updating subject")
	}

	return nil
}

func (s *SubjectService) Delete(name string) error {
	req, err := s.client.NewRequest("DEL", "/api/v1/recognition/subjects/"+name, nil)
	if err != nil {
		return err
	}

	var subject Subject
	_, err = req.Do(&subject)
	if err != nil {
		return err
	}

	return nil
}

func (s *SubjectService) DeleteAll() (*DeletedResponse, error) {
	req, err := s.client.NewRequest("DEL", "/api/v1/recognition/subjects/", nil)
	if err != nil {
		return nil, err
	}

	var deleted DeletedResponse
	_, err = req.Do(&deleted)
	if err != nil {
		return nil, err
	}

	return &deleted, nil
}
