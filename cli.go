package t

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/fatih/color"
	"golang.org/x/crypto/ssh/terminal"
)

type CLI struct{}

func (cli *CLI) Run() {
	actual := ""
	expected := ""

	if len(os.Args) <= 1 {
		cli.showUsage()
		os.Exit(0)
	}

	if len(os.Args) > 1 {
		expected = os.Args[1]
	}

	if len(os.Args) >= 2 {
		if !terminal.IsTerminal(int(os.Stdin.Fd())) {
			stdin, _ := ioutil.ReadAll(os.Stdin)
			actual = strings.TrimSpace(string(stdin))
		} else {
			actual = os.Args[2]
		}
	}

	cli.test(expected, actual)
}

func (cli *CLI) test(expected string, actual string) {
	if expected == actual {
		check := "\u2713"
		color.Green("%s PASS: got %v, want %v\n", check, actual, expected)
	} else {
		ballot := "\u2717"
		color.Red("%s FAIL: got %v, want %v\n", ballot, actual, expected)
	}
}

func (cli *CLI) showUsage() {
	fmt.Println("Usage: t hello hell")
}
