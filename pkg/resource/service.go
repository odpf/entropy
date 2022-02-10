package resource

import (
	"context"
	"github.com/odpf/entropy/domain"
	"github.com/odpf/entropy/store"
)

type ServiceInterface interface {
	CreateResource(ctx context.Context, res *domain.Resource) (*domain.Resource, error)
	UpdateResource(ctx context.Context, res *domain.Resource) (*domain.Resource, error)
	GetResource(ctx context.Context, urn string) (*domain.Resource, error)
	ListResources(ctx context.Context, parent string, kind string) ([]*domain.Resource, error)
}

type Service struct {
	resourceRepository store.ResourceRepository
}

func NewService(repository store.ResourceRepository) *Service {
	return &Service{
		resourceRepository: repository,
	}
}

func (s *Service) CreateResource(ctx context.Context, res *domain.Resource) (*domain.Resource, error) {
	res.Status = domain.ResourceStatusPending
	err := s.resourceRepository.Create(res)
	if err != nil {
		return nil, err
	}
	createdResource, err := s.resourceRepository.GetByURN(res.Urn)
	if err != nil {
		return nil, err
	}
	return createdResource, nil
}

func (s *Service) UpdateResource(ctx context.Context, res *domain.Resource) (*domain.Resource, error) {
	err := s.resourceRepository.Update(res)
	if err != nil {
		return nil, err
	}
	updatedRes, err := s.resourceRepository.GetByURN(res.Urn)
	if err != nil {
		return nil, err
	}
	return updatedRes, nil
}

func (s *Service) GetResource(ctx context.Context, urn string) (*domain.Resource, error) {
	return s.resourceRepository.GetByURN(urn)
}

func (s *Service) ListResources(ctx context.Context, parent string, kind string) ([]*domain.Resource, error) {
	return s.resourceRepository.List(parent, kind)
}
