package usecase

import (
	"context"
	"time"

	"github.com/DTunnel0/ManagerDT-Go/src/domain/contracts"
)

type GetUsersOutput struct {
	ID        int
	UUID      string
	Username  string
	Password  string
	Limit     int
	ExpiresAt time.Time
}

type GetUsersUseCase struct {
	userRepository contracts.UserRepository
}

func NewGetUsersUseCase(userRepository contracts.UserRepository) *GetUsersUseCase {
	return &GetUsersUseCase{
		userRepository: userRepository,
	}
}

func (c *GetUsersUseCase) Execute(ctx context.Context) ([]*GetUsersOutput, error) {
	users, err := c.userRepository.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	output := make([]*GetUsersOutput, len(users))
	for i, user := range users {
		output[i] = &GetUsersOutput{
			ID:        user.ID,
			UUID:      user.UUID,
			Username:  user.Username,
			Password:  user.Password,
			Limit:     user.Limit,
			ExpiresAt: user.ExpiresAt,
		}
	}

	return output, nil
}
