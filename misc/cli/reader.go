package cli

import (
	"bufio"
	"os"
)

var reader = bufio.NewReader(os.Stdin)

func ReadLine() string {
	line, _ := reader.ReadString('\n')
	return line
}
