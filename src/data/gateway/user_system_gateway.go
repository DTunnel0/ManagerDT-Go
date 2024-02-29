package gateway

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"

	"github.com/DTunnel0/ManagerDT-Go/src/domain/contracts"
	"github.com/DTunnel0/ManagerDT-Go/src/domain/entity"
)

type userSystemGateway struct{}

func NewUserSystemGateway() contracts.UserGateway {
	return &userSystemGateway{}
}

func (u *userSystemGateway) Create(ctx context.Context, user *entity.User) (*entity.User, error) {
	passwd, err := hashPassword(user.Password)
	if err != nil {
		return nil, err
	}

	cmd := exec.CommandContext(ctx, "useradd",
		"-p", passwd,
		"-s", "/bin/false",
		"-M",
		"-e", user.ExpiresAt.Format("2006-01-02"),
		user.Username,
	)

	err = cmd.Run()
	if err != nil {
		return nil, fmt.Errorf("failed to create user: %w", err)
	}

	cmd = exec.CommandContext(ctx, "id", "-u", user.Username)
	output, err := cmd.CombinedOutput()
	if err != nil {
		u.Delete(ctx, user)
		return nil, fmt.Errorf("failed to get user ID: %w", err)
	}

	ID, err := strconv.Atoi(strings.TrimSpace(string(output)))
	if err != nil {
		u.Delete(ctx, user)
		return nil, fmt.Errorf("failed to parse user ID: %w", err)
	}

	user.ID = ID
	user.CreatedAt = time.Now()

	return user, nil
}

func (u *userSystemGateway) Delete(ctx context.Context, user *entity.User) error {
	cmd := exec.CommandContext(ctx, "userdel", "--force", user.Username)
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("failed to delete user: %w", err)
	}
	return nil
}

func (u *userSystemGateway) ChangePassword(ctx context.Context, user *entity.User) error {
	cmd := exec.CommandContext(ctx,
		"bash",
		"-c",
		fmt.Sprintf("echo %s:%s | chpasswd", user.Username, user.Password),
	)

	cmd.Stdout = os.Stdout

	return cmd.Run()
}

func hashPassword(password string) (string, error) {
	cmd := exec.Command("openssl", "passwd", "-1", password)
	output, err := cmd.Output()
	if err != nil {
		return "", nil
	}
	return strings.TrimSpace(string(output)), nil
}
