package input

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	consolemenu "github.com/DTunnel0/ManagerDT-Go/pkg/console_menu"
)

type ConnectionLimit struct{}

func NewConnectionLimit() *ConnectionLimit {
	return &ConnectionLimit{}
}

func (c *ConnectionLimit) Value() int {
	reader := bufio.NewReader(os.Stdin)
	value := 0

	for value == 0 {
		fmt.Print(consolemenu.ApplyColor("Limite de conexões: ", consolemenu.YELLOW))
		text, _ := reader.ReadString('\n')
		text = strings.TrimSpace(text)
		limit, err := strconv.Atoi(text)
		if err != nil {
			fmt.Println(consolemenu.ApplyColor("Erro: Limite de conexões deve ser um número.", consolemenu.RED))
			continue
		}
		err = c.validate(limit)
		if err != nil {
			fmt.Println(consolemenu.ApplyColor(fmt.Sprintf("%v", err), consolemenu.RED))
			continue
		}
		value = limit
	}
	return value
}

func (c *ConnectionLimit) validate(value int) error {
	if value < 1 {
		return fmt.Errorf("%s", "Limite de conexões deve ser maior que 0.")
	}
	return nil
}
