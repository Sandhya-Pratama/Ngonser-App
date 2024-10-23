package service

import (
	"context"

	"github.com/Sandhya-Pratama/Ngonser-App/entity"
	"golang.org/x/crypto/bcrypt"
)

// register
type RegistrationUseCase interface {
	Registration(ctx context.Context, user *entity.User) error
}

type RegistrationRepository interface {
	Registration(ctx context.Context, user *entity.User) error
	// GetByEmail(ctx context.Context, email string) (*entity.User, error)
}

type registrationService struct {
	repository RegistrationRepository
}

func NewRegistrationService(repository RegistrationRepository) *registrationService {
	return &registrationService{
		repository: repository,
	}
}

func (s *registrationService) Registration(ctx context.Context, user *entity.User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user.Password = string(hashedPassword)
	return s.repository.Registration(ctx, user)
}
