package console

import (
	"context"
	"fmt"

	consolemenu "github.com/DTunnel0/ManagerDT-Go/pkg/console_menu"
	usecase "github.com/DTunnel0/ManagerDT-Go/src/domain/use_case"
	"github.com/DTunnel0/ManagerDT-Go/src/infra/console/input"
	"github.com/DTunnel0/ManagerDT-Go/src/infra/console/menu"
)

type ChangePasswordUserConsole struct {
	consoleMenuUser           *menu.UserMenuConsole
	getUsersUseCase           *usecase.GetUsersUseCase
	changePasswordUserUseCase *usecase.ChangePasswordUserUseCase
	password                  *input.Password
}

func NewChangePasswordUserConsole(
	menu *menu.UserMenuConsole,
	getUsersUseCase *usecase.GetUsersUseCase,
	changePasswordUserUseCase *usecase.ChangePasswordUserUseCase,
) *ChangePasswordUserConsole {
	return &ChangePasswordUserConsole{
		consoleMenuUser:           menu,
		getUsersUseCase:           getUsersUseCase,
		changePasswordUserUseCase: changePasswordUserUseCase,
		password:                  input.NewPassword(),
	}
}

func (c *ChangePasswordUserConsole) Run() error {
	ctx := context.Background()
	users, err := c.getUsersUseCase.Execute(ctx)
	if err != nil {
		return err
	}

	c.consoleMenuUser.SetCallback(func(u *menu.User) {
		password := c.password.Value()
		err := c.changePasswordUserUseCase.Execute(ctx, &usecase.ChangePasswordUserInput{
			ID:       u.ID,
			Username: u.Username,
			Password: password,
		})
		defer consolemenu.PausePrompt()
		if err == nil {
			u.Password = password
			fmt.Println(consolemenu.ApplyColor("Senha alterada com sucesso!", consolemenu.GREEN))
		}
	})

	for _, user := range users {
		c.consoleMenuUser.AddUser(&menu.User{
			ID:        user.ID,
			Username:  user.Username,
			Password:  user.Password,
			Limit:     user.Limit,
			ExpiresAt: user.ExpiresAt,
		})
	}

	c.consoleMenuUser.Display()
	return nil
}

type PasswordUserMenuConsoleBuilder struct{}

func NewPasswordUserMenuConsoleFormatter() func(*menu.User) string {
	return func(u *menu.User) string {
		return fmt.Sprintf("%-15s Senha: %s", u.Username, consolemenu.ApplyColor(u.Password, consolemenu.GREEN))
	}
}
