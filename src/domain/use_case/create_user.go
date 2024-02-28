package usecase

import (
	"context"
	"time"

	"github.com/DTunnel0/ManagerDT-Go/src/domain/contracts"
	"github.com/DTunnel0/ManagerDT-Go/src/domain/entity"
)

type CreateUserInput struct {
	UUID      string
	Username  string
	Password  string
	Limit     int
	ExpiresAt time.Time
}

type CreateUserUseCase struct {
	userRepository contracts.UserRepository
	userGateway    contracts.UserGateway
}

func NewCreateUserUseCase(userGateway contracts.UserGateway, userRepository contracts.UserRepository) *CreateUserUseCase {
	return &CreateUserUseCase{userGateway: userGateway, userRepository: userRepository}
}

func (u *CreateUserUseCase) Execute(ctx context.Context, data *CreateUserInput) error {
	user := &entity.User{
		UUID:      data.UUID,
		Username:  data.Username,
		Password:  data.Password,
		Limit:     data.Limit,
		CreatedAt: time.Now(),
		ExpiresAt: data.ExpiresAt,
	}

	user, err := u.userGateway.Create(ctx, user)
	if err != nil {
		return err
	}

	err = u.userRepository.Save(ctx, user)
	if err != nil {
		u.userGateway.Delete(ctx, user)
		return err
	}

	return nil
}
