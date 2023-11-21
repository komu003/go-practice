package services

import "app/pkg/repository"

type GoPracticeService struct {
	*MicropostService
	*UserService
}

func NewGoPracticeService(micropostRepo repository.MicropostRepository, userRepo repository.UserRepository) *GoPracticeService {
	return &GoPracticeService{
		MicropostService: NewMicropostService(micropostRepo),
		UserService:      NewUserService(userRepo),
	}
}
