package pass

import (
	"errors"
	"os"

	"github.com/MihaiBlebea/go-pass-client/caller"
	"github.com/MihaiBlebea/go-pass-client/resource/environment"
)

func EnvFrom(token string) error {
	passToken := os.Getenv("PASS_MANAGER_TOKEN")
	if passToken == "" {
		return errors.New("Could not find pass token. Add a PASS_MANAGER_TOKEN nv variable")
	}

	if token == "" {
		return errors.New("Environment token cannot be an empty string")
	}

	envi := environment.New(caller.New(passToken))
	if err := envi.PrepareEnvFrom(token); err != nil {
		return err
	}

	return nil
}
