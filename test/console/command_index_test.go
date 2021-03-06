package console

import (
	"github.com/confetti-framework/contract/inter"
	"github.com/confetti-framework/foundation/console"
	"github.com/stretchr/testify/require"
	"regexp"
	"strings"
	"testing"
)

func Test_index_show_title(t *testing.T) {
	output, app := setUp()
	code := console.Kernel{App: app, Writer: &output}.Handle()

	require.Equal(t, inter.Success, code)
	require.Contains(t, output.String(), "Confetti (testing)")
}

func Test_index_with_one_command(t *testing.T) {
	output, app := setUp()
	code := console.Kernel{
		App:    app,
		Writer: &output,
	}.Handle()

	require.Equal(t, inter.Success, code)
	require.Contains(
		t,
		TrimDoubleSpaces(output.String()),
		"Confetti (testing)\x1b[39m" +
			"\n\n" +
			" -h --help Can be used with any command to show\n" +
			" the command's available arguments and options.\n\n" +
			" baker Interact with your application.\n" +
			" log:clear Clear the log files as indicated in the configuration.",
	)
}

type aCommand struct {
	DryRun bool `flag:"-dry-run"`
}

func (s aCommand) Name() string        { return "a_command" }
func (s aCommand) Description() string { return "" }
func (s aCommand) Handle(_ inter.Cli) inter.ExitCode {
	return inter.Success
}

func Test_index_in_correct_order(t *testing.T) {
	output, app := setUp()
	code := console.Kernel{
		App:      app,
		Writer:   &output,
		Commands: []inter.Command{console.LogClear{}, aCommand{}},
	}.Handle()

	require.Equal(t, inter.Success, code)
	require.Regexp(t, "(?s)a_command.*log", output.String())
}

func TrimDoubleSpaces(value string) string {
	// Replace double spaces
	regex := regexp.MustCompile(` {2,}`)
	value = regex.ReplaceAllString(strings.Trim(value, " "), " ")

	// replace newline with only one space
	regex = regexp.MustCompile(` \n`)
	value = regex.ReplaceAllString(value, "\n")

	return value
}
