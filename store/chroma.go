package store

import (
	"context"

	chroma "github.com/amikos-tech/chroma-go"
	"github.com/pkg/errors"
)

type Chroma struct {
	Client     *chroma.Client
	Collection []Collection
}

func (c *Chroma) Init(_ context.Context) error {
	var err error

	c.Client, err = chroma.NewClient(chroma.WithBasePath("http://localhost:8000"))
	if err != nil {
		return errors.Wrap(err, "failed to create client\n")
	}

	return nil
}

func (c *Chroma) Deinit(_ context.Context) error {
	if c.Client == nil {
		return nil
	}

	if err := c.Client.Close(); err != nil {
		return errors.Wrap(err, "failed to close client\n")
	}

	return nil
}

func (c *Chroma) Reset(_ context.Context) error {
	return nil
}

func (c *Chroma) Save(_ context.Context, value interface{}, meta map[string]interface{}, agent string) error {
	return nil
}

func (c *Chroma) Search(_ context.Context, query string, limit int, threshold float64) ([]interface{}, error) {
	return nil, nil
}

func (c *Chroma) embedding(_ context.Context, text string, meta map[string]interface{}) ([]Collection, error) {
	return nil, nil
}
