package manager

import "SplitAll/repository"

type RepositoryManager interface {
	ImageUploadRepo() repository.ImageUploadRepository
	RecepientRepo() repository.RecepientRepository
}

type repositoryManager struct {
	infra Infra
}

func (r *repositoryManager) RecepientRepo() repository.RecepientRepository {
	return repository.NewRecepientRepository(r.infra.SqlDb())
}

func (r *repositoryManager) ImageUploadRepo() repository.ImageUploadRepository {
	return repository.NewImageUploadRepository(r.infra.SqlDb())
}

func NewRepositoryManager(infra Infra) RepositoryManager {
	return &repositoryManager{
		infra: infra,
	}
}
