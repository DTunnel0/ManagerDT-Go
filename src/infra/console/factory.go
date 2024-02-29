package console

import (
	consolemenu "github.com/DTunnel0/ManagerDT-Go/pkg/console_menu"
	"github.com/DTunnel0/ManagerDT-Go/src/data/gateway"
	"github.com/DTunnel0/ManagerDT-Go/src/data/repository"
	usecase "github.com/DTunnel0/ManagerDT-Go/src/domain/use_case"
	"github.com/DTunnel0/ManagerDT-Go/src/infra/console/menu"
)

func MakeCreateUserConsole() *CreateUserConsole {
	userGateway := gateway.NewUserSystemGateway()
	userRepository := repository.NewUserSQLiteRepository()
	createUserUseCase := usecase.NewCreateUserUseCase(userGateway, userRepository)
	return NewCreateUserConsole(createUserUseCase)
}

func MakeDeleteUserConsole() *DeleteUserConsole {
	consoleMenu := consolemenu.NewConsoleMenu("DELETAR USUARIO", consolemenu.NewFormatter())
	formatter := menu.NewDefaultFormatter()
	userRepository := repository.NewUserSQLiteRepository()
	userGateway := gateway.NewUserSystemGateway()
	getUsersUseCase := usecase.NewGetUsersUseCase(userRepository)
	deleteUserUseCase := usecase.NewDeleteUserUseCase(userRepository, userGateway)
	menu := menu.NewUserMenuConsole(consoleMenu, formatter)
	return NewDeleteUserConsole(menu, getUsersUseCase, deleteUserUseCase)
}

func MakeChangePasswordUserConsole() *ChangePasswordUserConsole {
	userRepository := repository.NewUserSQLiteRepository()
	userGateway := gateway.NewUserSystemGateway()
	getUsersUseCase := usecase.NewGetUsersUseCase(userRepository)
	changePasswordUserUseCase := usecase.NewChangePasswordUserUseCase(userGateway, userRepository)
	consoleMenu := consolemenu.NewConsoleMenu("ALTERAR SENHA", consolemenu.NewFormatter())
	formatter := NewPasswordUserMenuConsoleFormatter()
	menu := menu.NewUserMenuConsole(consoleMenu, formatter)
	return NewChangePasswordUserConsole(menu, getUsersUseCase, changePasswordUserUseCase)
}
