package input

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	consolemenu "github.com/DTunnel0/ManagerDT-Go/pkg/console_menu"
)

type Username struct {
	size    int
	minSize int
}

func NewUsername() *Username {
	return &Username{
		size:    32,
		minSize: 3,
	}
}

func (u *Username) Value() string {
	reader := bufio.NewReader(os.Stdin)
	value := ""

	for value == "" {
		fmt.Print(consolemenu.ApplyColor("Nome de usuário: ", consolemenu.YELLOW))
		text, _ := reader.ReadString('\n')
		text = strings.TrimSpace(text)
		err := u.validate(text)
		if err != nil {
			fmt.Println(consolemenu.ApplyColor(fmt.Sprintf("%v", err), consolemenu.RED))
			continue
		}
		value = text
	}
	return value
}

func (u *Username) validate(value string) error {
	if value == "" {
		return fmt.Errorf("%s", "Nome de usuário não pode ser vazio.")
	}
	if len(value) < u.minSize {
		return fmt.Errorf("%s", fmt.Sprintf("Nome de usuário deve ter no mínimo %d caracteres.", u.minSize))
	}
	if len(value) > u.size {
		return fmt.Errorf("%s", fmt.Sprintf("Nome de usuário deve ter no máximo %d caracteres.", u.size))
	}
	if strings.Contains(value, " ") {
		return fmt.Errorf("%s", "Nome de usuário não pode conter espaços.")
	}
	return nil
}
