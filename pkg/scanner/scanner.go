package scanner

import (
	"bufio"
	"os"
)

func InitScanner() *bufio.Scanner {
	return bufio.NewScanner(os.Stdin)
}
