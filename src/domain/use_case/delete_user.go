package usecase

import (
	"context"

	"github.com/DTunnel0/ManagerDT-Go/src/domain/contracts"
	"github.com/DTunnel0/ManagerDT-Go/src/domain/entity"
)

type DeleteUserInput struct {
	ID       int
	Username string
}

type DeleteUserUseCase struct {
	userGateway    contracts.UserGateway
	userRepository contracts.UserRepository
}

func NewDeleteUserUseCase(
	userRepository contracts.UserRepository,
	userGateway contracts.UserGateway,
) *DeleteUserUseCase {
	return &DeleteUserUseCase{
		userRepository: userRepository,
		userGateway:    userGateway,
	}
}

func (c *DeleteUserUseCase) Execute(ctx context.Context, input ...*DeleteUserInput) error {
	users := []*entity.User{}

	for _, data := range input {
		user := &entity.User{ID: data.ID, Username: data.Username}
		if err := c.userGateway.Delete(ctx, user); err == nil {
			users = append(users, user)
		}
	}

	return c.userRepository.Delete(ctx, users...)
}
