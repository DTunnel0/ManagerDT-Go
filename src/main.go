package main

import (
	consoleMenu "github.com/DTunnel0/ManagerDT-Go/pkg/console_menu"
)

func main() {
	formatter := consoleMenu.NewFormatter()

	submenu := consoleMenu.NewConsoleMenu("GERENCIAR USUARIOS", formatter)
	submenu.AddItem(consoleMenu.NewMenuItem(1, "CRIAR USUARIO"))
	submenu.AddItem(consoleMenu.NewMenuItem(2, "DELETAR USUARIO"))
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
