package t

import (
	"fmt"
	"go/constant"
	"io/ioutil"
	"os"
	"strings"

	"go/token"
	"go/types"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"golang.org/x/crypto/ssh/terminal"
)

type param struct {
	operator string
}

var rootCmd = &cobra.Command{
	Use:   "t [expected] [actual]",
	Short: "t is a command line tool for testing on your terminal",
	Run: func(cmd *cobra.Command, args []string) {

		actual := ""
		expected := ""

		if len(args) > 0 {
			expected = args[0]
		}

		if len(args) >= 1 {
			if !terminal.IsTerminal(int(os.Stdin.Fd())) {
				stdin, _ := ioutil.ReadAll(os.Stdin)
				actual = strings.TrimSpace(string(stdin))
			} else {
				actual = args[1]
			}
		}
		param := param{
			operator: "",
		}
		param.operator, _ = cmd.Flags().GetString("operator")
		test(expected, actual, param)
	},
}

func init() {
	rootCmd.Flags().StringP("operator", "o", "", "Operator")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func test(expected string, actual string, param param) {
	if param.operator == "" {
		if expected == actual {
			pass(fmt.Sprintf("%s == %s", expected, actual))
		} else {
			fail(fmt.Sprintf("%s != %s", expected, actual))
		}
	}

	expr := fmt.Sprintf("%s %s %s", expected, param.operator, actual)
	v := eval(expr)

	if v.String() == "true" {
		pass(expr)
	} else {
		fail(expr)
	}
}

func pass(message string) {
	check := "\u2713"
	color.Green("%s PASS: %s\n", check, message)
	os.Exit(0)
}

func fail(message string) {
	ballot := "\u2717"
	color.Red("%s FAIL: %v\n", ballot, message)
	os.Exit(0)
}

func eval(expr string) constant.Value {
	fset := token.NewFileSet()
	tv, err := types.Eval(fset, nil, token.NoPos, expr)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return tv.Value
}
