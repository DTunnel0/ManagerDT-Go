package input

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	consolemenu "github.com/DTunnel0/ManagerDT-Go/pkg/console_menu"
)

type ExpirationDate struct{}

func NewExpirationDate() *ExpirationDate {
	return &ExpirationDate{}
}

func (e *ExpirationDate) Value() time.Time {
	reader := bufio.NewReader(os.Stdin)
	var value time.Time

	for value.IsZero() {
		fmt.Print(consolemenu.ApplyColor("Data de expiração: ", consolemenu.YELLOW))
		text, _ := reader.ReadString('\n')
		text = strings.TrimSpace(text)
		days, err := strconv.Atoi(text)
		if err != nil {
			fmt.Println(consolemenu.ApplyColor("Erro: Data de expiração deve ser um número inteiro.", consolemenu.RED))
			continue
		}
		err = e.validate(days)
		if err != nil {
			fmt.Println(consolemenu.ApplyColor(fmt.Sprintf("%v", err), consolemenu.RED))
			continue
		}
		value = time.Now().Add(time.Hour * 24 * time.Duration(days))
	}
	return value
}

func (e *ExpirationDate) validate(days int) error {
	if days <= 0 {
		return fmt.Errorf("%s", "Data de expiração deve ser maior que zero.")
	}
	return nil
}
