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
	userRepository := repository.NewUserSQLiteRepository()
	userGateway := gateway.NewUserSystemGateway()
	getUsersUseCase := usecase.NewGetUsersUseCase(userRepository)
	deleteUserUseCase := usecase.NewDeleteUserUseCase(userRepository, userGateway)
	menu := menu.NewUserMenuConsole(consoleMenu)
	return NewDeleteUserConsole(menu, getUsersUseCase, deleteUserUseCase)
}
