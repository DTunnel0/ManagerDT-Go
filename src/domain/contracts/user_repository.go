package contracts

import (
	"context"

	"github.com/DTunnel0/ManagerDT-Go/src/domain/entity"
)

type UserRepository interface {
	Save(ctx context.Context, user *entity.User) error
	FindById(ctx context.Context, id int) (*entity.User, error)
	FindAll(ctx context.Context) ([]*entity.User, error)
	Delete(ctx context.Context, user ...*entity.User) error
}
