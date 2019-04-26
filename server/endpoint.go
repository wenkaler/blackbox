package server

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-kit/kit/endpoint"
)

type requestFVJ struct {
	Various string          `json:"various"` // Various means that can contain different kind of character values. (name - project or token.)
	Setting json.RawMessage `json:"setting"`
}

type requestFTNJ struct {
	Token   string          `json:"token"`
	Name    string          `json:"name"`
	Setting json.RawMessage `json:"setting"`
}

type requestFTV struct {
	Token   string `json:"token"`
	Various string `json:"various"` // Various means that can contain different kind of character values. (token - project or time interval ('24 hour')
}

func makeCreateProjectEndpoint(s service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		var (
			data = request.([]byte)
			req  requestFVJ
		)
		err := json.Unmarshal(data, &req)
		if err != nil {
			return nil, fmt.Errorf("makeCreateProjectEndpoint: %v", err)
		}
		b, err := s.FTJ("box.create_project_v1", req.Various, req.Setting)
		if err != nil {
			return nil, fmt.Errorf("makeCreateProjectEndpoint: %v", err)
		}
		return b, nil
	}
}

func makeUpdateProjectEndpoint(s service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		var (
			data = request.([]byte)
			req  requestFTNJ
		)
		err := json.Unmarshal(data, &req)
		if err != nil {
			return nil, fmt.Errorf("makeUpdateProjectEndpoint: %v", err)
		}
		b, err := s.FTNJ("box.update_project_v1", req.Token, req.Name, req.Setting)
		if err != nil {
			return nil, fmt.Errorf("makeUpdateProjectEndpoint: %v", err)
		}
		return b, nil
	}
}

func makeGetProjectEndpoint(s service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		token := request.(string)
		b, err := s.FT("box.get_project_v1", token)
		if err != nil {
			return nil, fmt.Errorf("makeGetProjectEndpoint: %v", err)
		}
		return b, nil
	}
}

func makeListProjectEndpoint(s service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		token := request.(string)
		b, err := s.FT("box.list_project_v1", token)
		if err != nil {
			return nil, fmt.Errorf("makeListProjectEndpoint: %v", err)
		}
		return b, nil
	}
}

func makeRemoveProjectEndpoint(s service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		token := request.(string)
		b, err := s.FT("box.remove_project_v1", token)
		if err != nil {
			return nil, fmt.Errorf("makeRemoveProjectEndpoint: %v", err)
		}
		return b, nil
	}
}

func makeUpdateSettingEndpoint(s service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		var (
			data = request.([]byte)
			req  requestFVJ
		)
		err := json.Unmarshal(data, &req)
		if err != nil {
			return nil, fmt.Errorf("makeUpdateSettingEndpoint: %v", err)
		}
		b, err := s.FTJ("box.update_setting_v1", req.Various, req.Setting)
		if err != nil {
			return nil, fmt.Errorf("makeUpdateSettingEndpoint: %v", err)
		}
		return b, nil
	}
}

func makeGetSettingEndpoint(s service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		token := request.(string)
		b, err := s.FT("box.get_setting_v1", token)
		if err != nil {
			return nil, fmt.Errorf("makeGetSettingEndpoint: %v", err)
		}
		return b, nil
	}
}

func makeConfirmSettingEndpoint(s service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		token := request.(string)
		b, err := s.FT("box.confirm_setting_v1", token)
		if err != nil {
			return nil, fmt.Errorf("makeConfirmSettingEndpoint: %v", err)
		}
		return b, nil
	}
}

func makeInitialSettingEndpoint(s service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(requestFTV)
		b, err := s.FTV("box.initial_setting_v1", req.Token, req.Various)
		if err != nil {
			return nil, fmt.Errorf("makeInitialSettingEndpoint: %v", err)
		}
		return b, nil
	}
}

func makeCleanUnusedSettingsEndpoint(s service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(requestFTV)
		b, err := s.FTV("box.clean_unused_settings_v1", req.Token, req.Various)
		if err != nil {
			return nil, fmt.Errorf("makeCleanUnusedSettingsEndpoint: %v", err)
		}
		return b, nil
	}
}
