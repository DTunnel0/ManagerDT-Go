package usecase

import (
	"context"

	"github.com/DTunnel0/ManagerDT-Go/src/domain/contracts"
)

type ChangePasswordUserInput struct {
	ID       int
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
	user, err := c.userRepository.FindById(ctx, data.ID)
	if err != nil {
		return err
	}

	user.Password = data.Password
	err = c.userGateway.ChangePassword(ctx, user)
	if err != nil {
		return err
	}

	return c.userRepository.Save(ctx, user)
}
