package provider

//go:generate mockery --name=Repository -r --case underscore --with-expecter --structname ProviderRepository --filename=provider_repository.go --output=./mocks

import (
	"context"
	"strings"
	"time"
)

type Repository interface {
	Migrate(ctx context.Context) error

	GetByURN(ctx context.Context, urn string) (*Provider, error)
	List(ctx context.Context, filter map[string]string) ([]*Provider, error)
	Create(ctx context.Context, r Provider) error
}

type Provider struct {
	URN       string                 `bson:"urn"`
	Name      string                 `bson:"name"`
	Kind      string                 `bson:"kind"`
	Parent    string                 `bson:"parent"`
	Labels    map[string]string      `bson:"labels"`
	Configs   map[string]interface{} `bson:"configs"`
	CreatedAt time.Time              `bson:"created_at"`
	UpdatedAt time.Time              `bson:"updated_at"`
}

func (p *Provider) Validate() error {
	p.URN = GenerateURN(*p)

	// TODO: add basic sanitization and validations here.
	return nil
}

func GenerateURN(pro Provider) string {
	return strings.Join([]string{
		sanitizeString(pro.Parent),
		sanitizeString(pro.Name),
	}, "-")
}

func sanitizeString(s string) string {
	return strings.ReplaceAll(s, " ", "_")
}
