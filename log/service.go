package log

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/netomathias/bts-client/errors"
	"net/http"
)

type Service struct {
	Client      *http.Client
	UrlBase     string
	SourceToken string
}

func NewService(client *http.Client, urlBase string, sourceToken string) Service {
	return Service{
		Client:      client,
		UrlBase:     urlBase,
		SourceToken: sourceToken,
	}
}

func (s *Service) Create(ctx context.Context, logData LogDataRequest) (*LogDataResponse, error) {
	j, err := json.Marshal(logData)
	if err != nil {
		return nil, err
	}

	endpoint := fmt.Sprintf("%s/")

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, endpoint, bytes.NewBuffer(j))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+s.SourceToken)
	req.Header.Set("Content-Type", "application/json")
	resp, err := s.Client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusBadRequest {
		return nil, errors.FromBadRequest(resp)
	}

	if resp.StatusCode != http.StatusAccepted {
		return nil, errors.FromHTTPResponse(resp)
	}

	response := &LogDataResponse{
		Message: "Success",
	}

	return response, nil
}

func (s *Service) CreateLogDataRequest(message string, nested Nested) LogDataRequest {
	return LogDataRequest{
		Message: message,
		Nested:  nested,
	}
}

func (s *Service) CreateNested(values any) Nested {
	return Nested{
		Values: values,
	}
}

func (s *Service) CreateLogDataResponse(message string) LogDataResponse {
	return LogDataResponse{
		Message: message,
	}
}
