package contracts

import (
	"context"

	"github.com/DTunnel0/ManagerDT-Go/src/domain/entity"
)

type UserGateway interface {
	Create(ctx context.Context, user *entity.User) (*entity.User, error)
	Delete(ctx context.Context, user *entity.User) error
}
