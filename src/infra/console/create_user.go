package console

import (
	"context"
	"fmt"

	consolemenu "github.com/DTunnel0/ManagerDT-Go/pkg/console_menu"
	usecase "github.com/DTunnel0/ManagerDT-Go/src/domain/use_case"
	"github.com/DTunnel0/ManagerDT-Go/src/infra/console/input"
)

const (
	MESSAGE_USER_CREATED = "\033[1;37m━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━\033[0m" +
		"\n\033[1;42m\033[1;37m          ✅USUARIO CRIADO COM SUCESSO✅          \033[0m" +
		"\n\033[1;37m━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━\033[0m" +
		"\n\033[1;37m👤Nome do usuario: \033[1;94m%s\033[0m" +
		"\n\033[1;37m🔑Senha: \033[1;94m%s\033[0m" +
		"\n\033[1;37m🔗Limite de conexoes: \033[1;94m%02d\033[0m" +
		"\n\033[1;37m📅Data de expiracao: \033[1;94m%s\033[0m" +
		"\n\033[1;37m━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━\033[0m\n"
)

type CreateUserConsole struct {
	createUserUseCase *usecase.CreateUserUseCase
	username          *input.Username
	password          *input.Password
	limit             *input.ConnectionLimit
	expiresAt         *input.ExpirationDate
}

func NewCreateUserConsole(createUserUseCase *usecase.CreateUserUseCase) *CreateUserConsole {
	return &CreateUserConsole{
		createUserUseCase: createUserUseCase,
		username:          input.NewUsername(),
		password:          input.NewPassword(),
		limit:             input.NewConnectionLimit(),
		expiresAt:         input.NewExpirationDate(),
	}
}

func (c *CreateUserConsole) Run() error {
	ctx := context.Background()

	username := c.username.Value()
	password := c.password.Value()
	limit := c.limit.Value()
	expiresAt := c.expiresAt.Value()

	err := c.createUserUseCase.Execute(ctx, &usecase.CreateUserInput{
		UUID:      "",
		Username:  username,
		Password:  password,
		Limit:     limit,
		ExpiresAt: expiresAt,
	})

	consolemenu.ClearScreen()
	defer consolemenu.PausePrompt()

	if err != nil {
		fmt.Println(err)
		return err
	}

	fmt.Printf(MESSAGE_USER_CREATED, username, password, limit, expiresAt.Format("02/01/2006"))
	return err
}
