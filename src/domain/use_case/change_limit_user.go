package usecase

import (
	"context"

	"github.com/DTunnel0/ManagerDT-Go/src/domain/contracts"
)

type ChangeLimitUserInput struct {
	ID    int
	Limit int
}

type ChangeLimitUserUseCase struct {
	userRepository contracts.UserRepository
}

func NewChangeLimitUserUseCase(userRepository contracts.UserRepository) *ChangeLimitUserUseCase {
	return &ChangeLimitUserUseCase{userRepository: userRepository}
}

func (c *ChangeLimitUserUseCase) Execute(ctx context.Context, data *ChangeLimitUserInput) error {
	user, err := c.userRepository.FindById(ctx, data.ID)
	if err != nil {
		return err
	}

	user.Limit = data.Limit

	return c.userRepository.Save(ctx, user)
}
