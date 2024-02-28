package main

import (
	consoleMenu "github.com/DTunnel0/ManagerDT-Go/pkg/console_menu"
	"github.com/DTunnel0/ManagerDT-Go/src/infra/console"
)

func main() {
	formatter := consoleMenu.NewFormatter()

	submenu := consoleMenu.NewConsoleMenu("GERENCIAR USUARIOS", formatter)
	submenu.AddItem(consoleMenu.NewFuncItem(1, "CRIAR USUARIO", false, func(a ...any) *any { console.MakeCreateUserConsole().Run(); return nil }))
	submenu.AddItem(consoleMenu.NewFuncItem(2, "DELETAR USUARIO", false, func(a ...any) *any { console.MakeDeleteUserConsole().Run(); return nil }))
	submenu.AddItem(consoleMenu.NewMenuItem(3, "ALTERAR SENHA"))
	submenu.AddItem(consoleMenu.NewMenuItem(4, "ALTERAR LIMITE"))
	submenu.AddItem(consoleMenu.NewExitItem())

	menu := consoleMenu.NewConsoleMenu("GERENCIADOR", formatter)
	menu.AddItem(consoleMenu.NewFuncItem(1, "GERENCIAR USUARIOS", false, func(a ...any) *any { submenu.Display(); return nil }))
	menu.AddItem(consoleMenu.NewMenuItem(2, "GERENCIAR CONEXOES"))
	menu.AddItem(consoleMenu.NewMenuItem(3, "GERENCIAR FERRAMENTAS"))
	menu.AddItem(consoleMenu.NewExitItem())
	menu.Display()
}
