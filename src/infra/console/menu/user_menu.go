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
	console   consolemenu.ConsoleMenu
	formatter func(*User) string
	callback  func(*User)
	exitItem  consolemenu.MenuItem
	users     []*User
}

func NewDefaultFormatter() func(*User) string {
	return func(u *User) string {
		return u.Username
	}
}

func NewUserMenuConsole(console consolemenu.ConsoleMenu, formatter func(*User) string) *UserMenuConsole {
	return &UserMenuConsole{
		console:   console,
		formatter: formatter,
		exitItem:  consolemenu.NewExitItem(),
		users:     []*User{},
	}
}

func (m *UserMenuConsole) SetCallback(callback func(*User)) {
	m.callback = callback
}

func (m *UserMenuConsole) AddUser(user *User) {
	m.console.RemoveItem(m.exitItem)
	defer m.console.AddItem(m.exitItem)

	m.users = append(m.users, user)
	m.console.AddItem(&UserItem{
		id:        len(m.users),
		user:      user,
		formatter: m.formatter,
		callback:  m.callback,
	})
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

	defer m.cleanUp()
	m.console.Display()
}

func (m *UserMenuConsole) cleanUp() {
	m.console.CleanUp()
	m.users = m.users[:0]
}

type UserMenuConsoleBuilder interface {
	Build(user *User) string
}

type UserItem struct {
	id        int
	user      *User
	formatter func(*User) string
	callback  func(*User)
}

func (u *UserItem) ID() int {
	return u.id
}
func (u *UserItem) Name() string {
	return u.formatter(u.user)
}
func (u *UserItem) ShouldExit() bool {
	return false
}

func (u *UserItem) Action() *any {
	if u.callback != nil {
		u.callback(u.user)
	}
	return nil
}
