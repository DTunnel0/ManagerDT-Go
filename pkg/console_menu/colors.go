package consolemenu

import (
	"fmt"
	"strings"
)

const (
	RED    = "\033[1;91m"
	GREEN  = "\033[1;92m"
	YELLOW = "\033[1;93m"
	BLUE   = "\033[1;94m"
	PINK   = "\033[1;95m"
	CYAN   = "\033[1;96m"
	WHITE  = "\033[1;97m"

	BG_RED    = "\033[1;41m"
	BG_GREEN  = "\033[1;42m"
	BG_YELLOW = "\033[1;43m"
	BG_BLUE   = "\033[1;44m"
	BG_PINK   = "\033[1;45m"
	BG_CYAN   = "\033[1;46m"
	BG_WHITE  = "\033[1;47m"

	RESET = "\033[0m"
)

func ApplyColor(text string, colors ...string) string {
	return RESET + fmt.Sprintf("%s%s", strings.Join(colors, ""), text) + RESET
}
