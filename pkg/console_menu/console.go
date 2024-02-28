package consolemenu

import (
	"fmt"
)

type ConsoleMenu interface {
	AddItem(item MenuItem)
	RemoveItem(item MenuItem)
	FindItem(ID int) MenuItem
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

func (c *consoleMenu) RemoveItem(item MenuItem) {
	for index, i := range c.items {
		if i == item {
			c.items = append(c.items[:index], c.items[index+1:]...)
			break
		}
	}
}

func (c *consoleMenu) FindItem(ID int) MenuItem {
	for _, item := range c.items {
		if item.ID() == ID {
			return item
		}
	}
	return nil
}

func (c *consoleMenu) Display() {
	defer c.cleanUp()

	for c.itemSelected == nil || !c.itemSelected.ShouldExit() {
		c.clearScreen()
		fmt.Print(c.formatter.Format(c.title, c.items))
		c.selectInput()
	}
}

func (c *consoleMenu) cleanUp() {
	c.itemReturned = nil
	c.itemSelected = nil
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
	ClearScreen()
}
