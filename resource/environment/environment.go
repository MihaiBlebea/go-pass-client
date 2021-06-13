package environment

import (
	"fmt"
	"os"
	"strings"

	"github.com/MihaiBlebea/go-pass-client/caller"
)

type service struct {
	caller caller.Service
}

type Service interface {
	PrepareEnvFrom(token string) error
}

type EnvResponse struct {
	Envs    []Env  `json:"envs,omitempty"`
	Success bool   `json:"success"`
	Message string `json:"message,omitempty"`
}

type Env struct {
	Label string `json:"label"`
	Value string `json:"value"`
}

func New(caller caller.Service) Service {
	return &service{caller}
}

func (s *service) PrepareEnvFrom(token string) error {
	resp := EnvResponse{}
	endpoint := fmt.Sprintf("/env/%s", token)

	err := s.caller.Get(endpoint, &resp)
	if err != nil {
		return err
	}

	if len(resp.Envs) == 0 {
		return nil
	}

	for _, env := range resp.Envs {
		os.Setenv(strings.ToUpper(env.Label), env.Value)
	}

	return nil
}
