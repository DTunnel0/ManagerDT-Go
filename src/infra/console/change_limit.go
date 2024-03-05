package console

import (
	"context"
	"fmt"

	consolemenu "github.com/DTunnel0/ManagerDT-Go/pkg/console_menu"
	usecase "github.com/DTunnel0/ManagerDT-Go/src/domain/use_case"
	"github.com/DTunnel0/ManagerDT-Go/src/infra/console/input"
	"github.com/DTunnel0/ManagerDT-Go/src/infra/console/menu"
)

type ChangeLimitUserConsole struct {
	consoleMenuUser        *menu.UserMenuConsole
	getUsersUseCase        *usecase.GetUsersUseCase
	changeLimitUserUseCase *usecase.ChangeLimitUserUseCase
	limit                  *input.ConnectionLimit
}

func NewChangeLimitUserConsole(
	menu *menu.UserMenuConsole,
	getUsersUseCase *usecase.GetUsersUseCase,
	changeLimitUserUseCase *usecase.ChangeLimitUserUseCase,
) *ChangeLimitUserConsole {
	return &ChangeLimitUserConsole{
		consoleMenuUser:        menu,
		getUsersUseCase:        getUsersUseCase,
		changeLimitUserUseCase: changeLimitUserUseCase,
		limit:                  input.NewConnectionLimit(),
	}
}

func (c *ChangeLimitUserConsole) Run() error {
	ctx := context.Background()
	users, err := c.getUsersUseCase.Execute(ctx)
	if err != nil {
		return err
	}

	c.consoleMenuUser.SetCallback(func(u *menu.User) {
		limit := c.limit.Value()
		err := c.changeLimitUserUseCase.Execute(ctx, &usecase.ChangeLimitUserInput{
			ID:    u.ID,
			Limit: limit,
		})
		defer consolemenu.PausePrompt()
		if err == nil {
			u.Limit = limit
			fmt.Println(consolemenu.ApplyColor("Limite alterado com sucesso!", consolemenu.GREEN))
			return
		}
		fmt.Println(consolemenu.ApplyColor(fmt.Sprintf("%v", err), consolemenu.RED))
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

func NewLimitUserMenuConsoleFormatter() func(*menu.User) string {
	return func(u *menu.User) string {
		return fmt.Sprintf("%-15s Limite: %s", u.Username, consolemenu.ApplyColor(fmt.Sprintf("%02d", u.Limit), consolemenu.GREEN))
	}
}
