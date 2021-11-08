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
	evalFlag bool
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
		if !terminal.IsTerminal(int(os.Stdin.Fd())) {
			stdin, _ := ioutil.ReadAll(os.Stdin)
			actual = strings.TrimSpace(string(stdin))
		}
		if len(args) > 1 {
			actual = args[1]
		}

		param := param{
			operator: "",
			evalFlag: false,
		}
		param.operator, _ = cmd.Flags().GetString("operator")
		param.evalFlag, _ = cmd.Flags().GetBool("eval")

		if param.operator == "" {
			param.operator = "=="
		} else {
			param.evalFlag = true
		}
		t(expected, actual, param)
	},
}

func init() {
	rootCmd.Flags().StringP("operator", "o", "", "Operator")
	rootCmd.Flags().BoolP("eval", "e", false, "Eval value")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func t(expected string, actual string, param param) {
	passed := test(expected, actual, param)
	if passed {
		pass(expected, param.operator, actual)
	}
	fail(expected, param.operator, actual)
}

func test(expected string, actual string, param param) bool {
	if param.evalFlag {
		expr := fmt.Sprintf("%s %s %s", expected, param.operator, actual)
		v := eval(expr)
		return v.String() == "true"
	}
	return expected == actual
}

func eval(expr string) constant.Value {
	fset := token.NewFileSet()
	tv, err := types.Eval(fset, nil, token.NoPos, expr)
	if err != nil {
		fmt.Println(expr)
		fmt.Println(err)
		os.Exit(1)
	}
	return tv.Value
}

func pass(expected string, operator string, actual string) {
	check := "\u2713"
	bgGreen := color.New(color.Bold, color.BgGreen, color.FgHiWhite).SprintFunc()
	show(bgGreen(fmt.Sprintf("%s PASS ", check)), expected, operator, actual)
	os.Exit(0)
}

func fail(expected string, operator string, actual string) {
	ballot := "\u2717"
	bgRed := color.New(color.Bold, color.BgRed, color.FgHiWhite).SprintFunc()
	show(bgRed(fmt.Sprintf("%s FAIL ", ballot)), expected, operator, actual)
	os.Exit(1)
}

func show(status string, expected string, operator string, actual string) {
	reset := color.New(color.Reset).SprintfFunc()
	white := color.New(color.Bold, color.FgHiWhite).SprintFunc()

	fmt.Printf("%s ", status)
	fmt.Printf("%s %s ", reset(""), white(expected))
	fmt.Printf("%s ", white(operator))
	fmt.Printf("%s %s\n", white(actual), reset(""))
}
