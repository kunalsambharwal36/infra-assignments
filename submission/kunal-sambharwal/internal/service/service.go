package service

import (
	"github.com/kunalsambharwal36/infra-assignments/internal/domain"
	"github.com/kunalsambharwal36/infra-assignments/internal/repository"
)

type ConfigService struct {
	repo *repository.ConfigRepository
}

func NewConfigService() *ConfigService {

	repo := repository.NewConfigRepository()

	return &ConfigService{
		repo: repo,
	}
}

// Create Config
func (s *ConfigService) CreateConfig(config domain.Config) error {

	return s.repo.CreateConfig(config)

}

// Get Config
func (s *ConfigService) GetConfig(id string) (domain.Config, error) {

	return s.repo.GetConfig(id)

}
