package input

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	consolemenu "github.com/DTunnel0/ManagerDT-Go/pkg/console_menu"
)

type Password struct {
	size    int
	minSize int
}

func NewPassword() *Password {
	return &Password{
		size:    32,
		minSize: 3,
	}
}

func (p *Password) Value() string {
	reader := bufio.NewReader(os.Stdin)
	value := ""

	for value == "" {
		fmt.Print(consolemenu.ApplyColor("Senha: ", consolemenu.YELLOW))
		text, _ := reader.ReadString('\n')
		text = strings.TrimSpace(text)
		err := p.validate(text)
		if err != nil {
			fmt.Println(consolemenu.ApplyColor(fmt.Sprintf("%v", err), consolemenu.RED))
			continue
		}
		value = text
	}
	return value
}

func (p *Password) validate(value string) error {
	if value == "" {
		return fmt.Errorf("%s", "Senha não pode ser vazia.")
	}
	if len(value) < p.minSize {
		return fmt.Errorf("%s", fmt.Sprintf("Senha deve ter no mínimo %d caracteres.", p.minSize))
	}
	if len(value) > p.size {
		return fmt.Errorf("%s", fmt.Sprintf("Senha deve ter no máximo %d caracteres.", p.size))
	}
	if strings.Contains(value, " ") {
		return fmt.Errorf("%s", "Senha não pode conter espaços.")
	}
	return nil
}
