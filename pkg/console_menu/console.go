package consolemenu

import (
	"fmt"
	"os"
	"os/exec"
)

type ConsoleMenu interface {
	AddItem(item MenuItem)
	Display()
}

type consoleMenu struct {
	title        string
	formatter    Formatter
	items        []MenuItem
	itemSelected MenuItem
	itemReturned *any
}

func NewConsoleMenu(title string, formatter Formatter) ConsoleMenu {
	return &consoleMenu{
		title:     title,
		formatter: formatter,
		items:     []MenuItem{},
	}
}

func (c *consoleMenu) AddItem(item MenuItem) {
	c.items = append(c.items, item)
}

func (c *consoleMenu) Display() {
	for {
		c.clearScreen()
		fmt.Print(c.formatter.Format(c.title, c.items))
		c.selectInput()

		if c.itemSelected.ShouldExit() {
			break
		}
	}
}

func (c *consoleMenu) selectInput() {
	var index int
	fmt.Scan(&index)

	for _, item := range c.items {
		if item.ID() == index {
			c.itemReturned = item.Action()
			c.itemSelected = item
			break
		}
	}
}

func (c *consoleMenu) clearScreen() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}
