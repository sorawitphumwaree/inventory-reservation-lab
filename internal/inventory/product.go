package inventory

import (
	"errors"
	"strings"
)

type CreateProductInput struct {
	Name         string
	InitialStock int64
}

func (input CreateProductInput) Validate() error {
	if strings.TrimSpace(input.Name) == "" {
		return errors.New("product name is required")
	}

	if input.InitialStock < 0 {
		return errors.New("initial stock must not be negative")
	}

	return nil
}
