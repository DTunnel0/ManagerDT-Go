package usecase

import (
	"context"

	"github.com/DTunnel0/ManagerDT-Go/src/domain/contracts"
	"github.com/DTunnel0/ManagerDT-Go/src/domain/entity"
)

type ChangePasswordUserInput struct {
	ID       int
	Username string
	Password string
}

type ChangePasswordUserUseCase struct {
	userGateway    contracts.UserGateway
	userRepository contracts.UserRepository
}

func NewChangePasswordUserUseCase(
	userGateway contracts.UserGateway,
	userRepository contracts.UserRepository,
) *ChangePasswordUserUseCase {
	return &ChangePasswordUserUseCase{
		userGateway:    userGateway,
		userRepository: userRepository,
	}
}

func (c *ChangePasswordUserUseCase) Execute(ctx context.Context, data *ChangePasswordUserInput) error {
	user := &entity.User{
		ID:       data.ID,
		Username: data.Username,
		Password: data.Password,
	}

	err := c.userGateway.ChangePassword(ctx, user)
	if err != nil {
		return err
	}

	return c.userRepository.ChangePassword(ctx, user)
}
