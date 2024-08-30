package manager

import "SplitAll/usecase"

type UsecaseManager interface {
	UserUsecase() usecase.UserUsecase
}

type usecaseManager struct {
	repoManager RepositoryManager
}

func (u *usecaseManager) UserUsecase() usecase.UserUsecase {
	return usecase.NewUserUsecase(u.repoManager.ImageUploadRepo(), u.repoManager.RecepientRepo(), u.repoManager.OcrRepo())
}

func NewUsecaseManager(repoManager RepositoryManager) UsecaseManager {
	return &usecaseManager{
		repoManager: repoManager,
	}
}
