package clearer

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

func ClearTerminal() {
	var cmd *exec.Cmd

	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("cmd", "/c", "cls")
	default:
		cmd = exec.Command("clear")
	}

	cmd.Stdout = os.Stdout

	if err := cmd.Run(); err != nil {
		clearWithEscapeSequence()
	}
}

func clearWithEscapeSequence() {
	const clearCode = "\033[H\033[2J"

	fmt.Print(clearCode)
	fmt.Fprint(os.Stdout, clearCode)
}
