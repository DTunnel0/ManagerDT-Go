package menu

import (
	"fmt"
	"time"

	consolemenu "github.com/DTunnel0/ManagerDT-Go/pkg/console_menu"
)

type User struct {
	ID        int
	Username  string
	Password  string
	Limit     int
	ExpiresAt time.Time
}

type UserMenuConsole struct {
	console  consolemenu.ConsoleMenu
	users    []*User
	callback func(*User)
	exitItem consolemenu.MenuItem
}

func NewUserMenuConsole(console consolemenu.ConsoleMenu) *UserMenuConsole {
	return &UserMenuConsole{
		console:  console,
		users:    []*User{},
		exitItem: consolemenu.NewExitItem(),
	}
}

func (m *UserMenuConsole) SetCallback(callback func(*User)) {
	m.callback = callback
}

func (m *UserMenuConsole) AddUser(user *User) {
	m.console.RemoveItem(m.exitItem)
	defer m.console.AddItem(m.exitItem)

	m.users = append(m.users, user)
	m.console.AddItem(consolemenu.NewFuncItem(
		len(m.users),
		user.Username,
		false,
		func(a ...any) *any {
			m.callback(user)
			return nil
		},
	))
}

func (m *UserMenuConsole) RemoveUser(user *User) {
	for i, u := range m.users {
		if u.ID == user.ID {
			item := m.console.FindItem(i + 1)
			m.console.RemoveItem(item)
			m.users = append(m.users[:i], m.users[i+1:]...)
			break
		}
	}
}

func (m *UserMenuConsole) Display() {
	if len(m.users) == 0 {
		fmt.Println(consolemenu.ApplyColor("Nenhum usuario encontrado...", consolemenu.RED))
		consolemenu.PausePrompt()
		return
	}

	m.console.Display()
}
