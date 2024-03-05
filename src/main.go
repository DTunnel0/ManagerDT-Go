package main

import (
	consoleMenu "github.com/DTunnel0/ManagerDT-Go/pkg/console_menu"
	"github.com/DTunnel0/ManagerDT-Go/src/infra/console"
)

func main() {
	formatter := consoleMenu.NewFormatter()

	createUserConsole := console.MakeCreateUserConsole()
	deleteUserConsole := console.MakeDeleteUserConsole()
	changePasswordUserConsole := console.MakeChangePasswordUserConsole()
	changeLimitUserConsole := console.MakeChangeLimitUserConsole()

	submenu := consoleMenu.NewConsoleMenu("GERENCIAR USUARIOS", formatter)
	submenu.AddItem(consoleMenu.NewFuncItem(1, "CRIAR USUARIO", false, func(a ...any) *any { createUserConsole.Run(); return nil }))
	submenu.AddItem(consoleMenu.NewFuncItem(2, "DELETAR USUARIO", false, func(a ...any) *any { deleteUserConsole.Run(); return nil }))
	submenu.AddItem(consoleMenu.NewFuncItem(3, "ALTERAR SENHA", false, func(a ...any) *any { changePasswordUserConsole.Run(); return nil }))
	submenu.AddItem(consoleMenu.NewFuncItem(4, "ALTERAR LIMITE", false, func(a ...any) *any { changeLimitUserConsole.Run(); return nil }))
	submenu.AddItem(consoleMenu.NewExitItem())

	menu := consoleMenu.NewConsoleMenu("GERENCIADOR", formatter)
	menu.AddItem(consoleMenu.NewFuncItem(1, "GERENCIAR USUARIOS", false, func(a ...any) *any { submenu.Display(); return nil }))
	menu.AddItem(consoleMenu.NewMenuItem(2, "GERENCIAR CONEXOES"))
	menu.AddItem(consoleMenu.NewMenuItem(3, "GERENCIAR FERRAMENTAS"))
	menu.AddItem(consoleMenu.NewExitItem())
	menu.Display()
}
