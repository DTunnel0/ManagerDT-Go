package consolemenu

import (
	"fmt"
	"strings"
	"time"
)

const (
	lineLength = 50
)

func createBanner() string {
	distroName := GetDistroName()
	ipv4 := GetIPV4()
	stats := GetMemStats()
	clock := time.Now().Format("15:04:05")

	menu := createMenuBar("DT MANAGER") + "\n"
	menu += ApplyColor(fmt.Sprintf("OS: %-31s", ApplyColor(distroName, GREEN)), RED)
	menu += ApplyColor(fmt.Sprintf("BY: %-27s", ApplyColor("@DuTra01", GREEN)), RED)
	menu += ApplyColor(fmt.Sprintf("HORA: %s\n", ApplyColor(clock, GREEN)), RED)

	menu += ApplyColor(fmt.Sprintf("IP: %-31s", ApplyColor(ipv4, GREEN)), RED)
	menu += ApplyColor(fmt.Sprintf("RAM: %-26s", ApplyColor(fmt.Sprintf("%.1fG", stats.Total), GREEN)), RED)
	menu += ApplyColor(fmt.Sprintf("Em uso %s", ApplyColor(fmt.Sprintf("%.1fG", stats.Used), GREEN)), RED)

	return menu
}

type Formatter interface {
	Format(title string, items []MenuItem) string
}

type formatter struct{}

func NewFormatter() Formatter {
	return &formatter{}
}

func (f *formatter) Format(title string, items []MenuItem) string {
	var builder strings.Builder

	builder.WriteString(createBanner())
	builder.WriteString("\n")
	builder.WriteString(createMenuBar(title))
	builder.WriteString("\n")

	for _, item := range items {
		builder.WriteString(fmt.Sprintf("%s%s%s %s %s\n",
			ApplyColor("[", BLUE),
			ApplyColor(fmt.Sprintf("%02d", item.ID()), GREEN),
			ApplyColor("]", BLUE),
			ApplyColor("•", GREEN),
			ApplyColor(item.Name(), RED),
		))
	}

	builder.WriteString(ApplyColor(strings.Repeat("━", lineLength), BLUE))
	builder.WriteString(ApplyColor("\nEscolha uma opcao: ", RED))

	return builder.String()
}

func createMenuBar(title string) string {
	var builder strings.Builder

	titleLength := len(title)
	sideLength := (lineLength - titleLength - 2) / 2

	builder.WriteString(ApplyColor(strings.Repeat("━", lineLength), BLUE) + "\n")

	titleBar := fmt.Sprintf("%s%s%s", ApplyColor("[", BG_RED, WHITE), ApplyColor(title, GREEN, BG_RED), ApplyColor("]", BG_RED, WHITE))

	titleBar = ApplyColor(strings.Repeat("-", sideLength), BG_RED, WHITE) + titleBar + ApplyColor(strings.Repeat("-", sideLength), BG_RED)

	if (titleLength+2)%2 != 0 {
		titleBar += ApplyColor("-", BG_RED, WHITE)
	}

	builder.WriteString(titleBar + "\n")
	builder.WriteString(ApplyColor(strings.Repeat("━", lineLength), RESET, BLUE))

	return builder.String()
}
