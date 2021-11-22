package t

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"

	"go/token"
	"go/types"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"golang.org/x/term"
)

type param struct {
	operator     string
	evalFlag     bool
	matchFlag    bool
	notMatchFlag bool
}

var rootCmd = &cobra.Command{
	Use:   "t [expected] [received]",
	Short: "t is a command line tool for testing on your terminal",
	Run: func(cmd *cobra.Command, args []string) {

		received := ""
		expected := ""

		if len(args) > 0 {
			expected = args[0]
		}
		if !term.IsTerminal(int(os.Stdin.Fd())) {
			stdin, _ := ioutil.ReadAll(os.Stdin)
			received = strings.TrimSpace(string(stdin))
		}
		if len(args) > 1 {
			received = args[1]
		}

		param := param{
			operator:     "",
			evalFlag:     false,
			matchFlag:    false,
			notMatchFlag: false,
		}

		param.operator, _ = cmd.Flags().GetString("operator")
		param.evalFlag, _ = cmd.Flags().GetBool("eval")
		param.matchFlag, _ = cmd.Flags().GetBool("match")
		param.notMatchFlag, _ = cmd.Flags().GetBool("not-match")

		if param.operator == "" {
			param.operator = "=="
		} else {
			param.evalFlag = true
		}
		t(expected, received, param)
	},
}

func init() {
	rootCmd.Flags().StringP("operator", "o", "", "Operator")
	rootCmd.Flags().BoolP("eval", "e", false, "Eval value")
	rootCmd.Flags().BoolP("match", "m", false, "Match regexp")
	rootCmd.Flags().BoolP("not-match", "", false, "Not match regexp")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func t(expected string, received string, param param) {
	passed := test(expected, received, param)

	operator := param.operator
	if param.matchFlag {
		operator = "match"
	} else if param.notMatchFlag {
		operator = "not match"
	}

	if passed {
		pass(expected, operator, received)
	}
	fail(expected, operator, received)
}

func test(expected string, received string, param param) bool {
	if param.matchFlag {
		return evalRegexp(expected, received)
	} else if param.notMatchFlag {
		return !evalRegexp(expected, received)
	}

	if param.evalFlag {
		expr := fmt.Sprintf("%s %s %s", expected, param.operator, received)
		return eval(expr)
	}
	return expected == received
}

func eval(expr string) bool {
	fset := token.NewFileSet()
	tv, err := types.Eval(fset, nil, token.NoPos, expr)
	if err != nil {
		fmt.Println(expr)
		fmt.Println(err)
		os.Exit(1)
	}
	return tv.Value.String() == "true"
}

func evalRegexp(regex string, input string) bool {
	r := regexp.MustCompile(regex)
	return r.MatchString(input)
}

func pass(expected string, operator string, received string) {
	check := "\u2713"
	bgGreen := color.New(color.Bold, color.BgGreen, color.FgHiWhite).SprintFunc()
	show(bgGreen(fmt.Sprintf("%s PASS ", check)), expected, operator, received)
	os.Exit(0)
}

func fail(expected string, operator string, received string) {
	ballot := "\u2717"
	bgRed := color.New(color.Bold, color.BgRed, color.FgHiWhite).SprintFunc()
	show(bgRed(fmt.Sprintf("%s FAIL ", ballot)), expected, operator, received)
	os.Exit(1)
}

func show(status string, expected string, operator string, received string) {
	white := color.New(color.Bold, color.FgHiWhite).SprintFunc()

	left := white(fmt.Sprintf("%s", expected))
	right := white(fmt.Sprintf("%s", received))

	fmt.Printf("%s ", status)
	fmt.Printf("[%s] ", operator)
	fmt.Printf("%s %s\n", left, right)
}
