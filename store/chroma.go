package store

import (
	"context"

	chroma "github.com/amikos-tech/chroma-go"
	"github.com/amikos-tech/chroma-go/types"
	"github.com/pkg/errors"
)

type Chroma struct {
	Url string

	Client     *chroma.Client
	Collection *chroma.Collection
	RecordSet  *types.RecordSet
}

func (c *Chroma) Init(ctx context.Context, name string) error {
	var err error

	if c.Client, err = chroma.NewClient(chroma.WithBasePath(c.Url)); err != nil {
		return errors.Wrap(err, "failed to create client\n")
	}

	if c.Collection, err = c.Client.NewCollection(ctx, name); err != nil {
		return errors.Wrap(err, "failed to create collection\n")
	}

	if c.RecordSet, err = types.NewRecordSet(types.WithEmbeddingFunction(c.Collection.EmbeddingFunction),
		types.WithIDGenerator(types.NewULIDGenerator())); err != nil {
		return errors.Wrap(err, "failed to create recordset\n")
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

	c.Client = nil
	c.Collection = nil

	return nil
}

func (c *Chroma) Reset(ctx context.Context) error {
	// Set ALLOW_RESET to true in the environment variables of the chroma server
	if _, err := c.Client.Reset(ctx); err != nil {
		return errors.Wrap(err, "failed to reset client\n")
	}

	c.Collection = nil

	return nil
}

func (c *Chroma) Save(ctx context.Context, text string, meta map[string]interface{}, _ string) error {
	for key, val := range meta {
		c.RecordSet.WithRecord(types.WithDocument(text), types.WithMetadata(key, val))
	}

	if _, err := c.RecordSet.BuildAndValidate(ctx); err != nil {
		return errors.Wrap(err, "failed to build and validate\n")
	}

	if _, err := c.Collection.AddRecords(ctx, c.RecordSet); err != nil {
		return errors.Wrap(err, "failed to add records\n")
	}

	return nil
}

func (c *Chroma) Search(ctx context.Context, query string, limit int32, threshold float32) ([]interface{}, error) {
	var ret []interface{}

	buf, err := c.Collection.Query(ctx, []string{query}, limit, nil, nil, nil)
	if err != nil {
		return nil, errors.Wrap(err, "failed to query collection\n")
	}

	for i := range len(buf.Ids[0]) {
		b := Collection{
			Id:      buf.Ids[0][i],
			Meta:    buf.Metadatas[0][i],
			Context: buf.Documents[0][i],
			Score:   buf.Distances[0][i],
		}
		if b.Score >= threshold {
			ret = append(ret, b)
		}
	}

	return ret, nil
}
