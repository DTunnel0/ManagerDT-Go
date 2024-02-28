package console

import (
	"context"
	"fmt"

	consolemenu "github.com/DTunnel0/ManagerDT-Go/pkg/console_menu"
	usecase "github.com/DTunnel0/ManagerDT-Go/src/domain/use_case"
	"github.com/DTunnel0/ManagerDT-Go/src/infra/console/menu"
)

type DeleteUserConsole struct {
	consoleMenuUser   *menu.UserMenuConsole
	getUsersUseCase   *usecase.GetUsersUseCase
	deleteUserUseCase *usecase.DeleteUserUseCase
}

func NewDeleteUserConsole(
	consoleMenuUser *menu.UserMenuConsole,
	getUsersUseCase *usecase.GetUsersUseCase,
	deleteUserUseCase *usecase.DeleteUserUseCase,
) *DeleteUserConsole {
	return &DeleteUserConsole{
		consoleMenuUser:   consoleMenuUser,
		getUsersUseCase:   getUsersUseCase,
		deleteUserUseCase: deleteUserUseCase,
	}
}

func (c *DeleteUserConsole) Run() error {
	ctx := context.Background()
	users, err := c.getUsersUseCase.Execute(ctx)
	if err != nil {
		return err
	}

	for _, user := range users {
		c.consoleMenuUser.AddUser(&menu.User{
			ID:        user.ID,
			Username:  user.Username,
			Password:  user.Password,
			Limit:     user.Limit,
			ExpiresAt: user.ExpiresAt,
		})
	}

	c.consoleMenuUser.SetCallback(func(u *menu.User) {
		err := c.deleteUserUseCase.Execute(ctx, &usecase.DeleteUserInput{ID: u.ID, Username: u.Username})
		defer consolemenu.PausePrompt()
		if err == nil {
			fmt.Println(consolemenu.ApplyColor(fmt.Sprintf("Usuario %s deletado com suceso!", u.Username), consolemenu.GREEN))
			c.consoleMenuUser.RemoveUser(u)
		} else {
			fmt.Println(consolemenu.ApplyColor(fmt.Sprintf("%v", err), consolemenu.RED))
		}
	})

	c.consoleMenuUser.Display()
	return nil
}
