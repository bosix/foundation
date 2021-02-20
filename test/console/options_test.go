package console

import (
	"github.com/confetti-framework/contract/inter"
	"github.com/confetti-framework/foundation/console"
	"github.com/confetti-framework/foundation/console/service"
	"github.com/stretchr/testify/require"
	"io"
	"testing"
)

type mockCommandWithoutOptions struct{}

func Test_show_index_if_no_command(t *testing.T) {
	output, app := setUp()
	app.Bind("config.App.OsArgs", []interface{}{"/exe/main"})

	code := console.Kernel{
		App:      app,
		Output:   &output,
		Commands: []inter.Command{structWithDescription{}},
	}.Handle()

	require.Equal(t, inter.Success, code)
	require.Contains(t, TrimDoubleSpaces(
		output.String()),
		`

 COMMAND DESCRIPTION

 test test

`)
}

func Test_get_option_from_command_without_options(t *testing.T) {
	options := service.GetOptions(mockCommandWithoutOptions{})
	require.Len(t, options, 0)
}

type mockCommandOption struct {
	DryRun bool `flag:"dry-run"`
}

func Test_get_parsed_option(t *testing.T) {
	options := service.GetOptions(mockCommandOption{})

	require.Equal(t, "dry-run", options[0].Tag.Get("flag"))
	require.Equal(t, false, options[0].Value)
}

type mockCommandMultipleOptions struct {
	DryRun bool `flag:"dry-run"`
	Vvv    bool `flag:"vvv"`
}

func Test_get_parsed_option_multiple_fields(t *testing.T) {
	options := service.GetOptions(mockCommandMultipleOptions{})

	require.Equal(t, "dry-run", options[0].Tag.Get("flag"))
	require.Equal(t, false, options[0].Value)
	require.Equal(t, "vvv", options[1].Tag.Get("flag"))
	require.Equal(t, false, options[1].Value)
}

type mockCommandOptions struct {
	DryRun bool `short:"dr" flag:"dry-run"`
}

func Test_get_parsed_multiple_option(t *testing.T) {
	options := service.GetOptions(mockCommandOptions{})

	require.Equal(t, "dr", options[0].Tag.Get("short"))
	require.Equal(t, "dry-run", options[0].Tag.Get("flag"))
}

type mockCommandOptionBool struct {
	DryRun bool `flag:"dry-run"`
}

func Test_get_parsed_option_bool(t *testing.T) {
	options := service.GetOptions(mockCommandOptionBool{})

	require.Equal(t, "dry-run", options[0].Tag.Get("flag"))
	require.Equal(t, false, options[0].Value)
}

type mockCommandOptionString struct {
	Username string `flag:"username"`
}

func Test_get_parsed_option_string(t *testing.T) {
	options := service.GetOptions(mockCommandOptionString{})

	require.Equal(t, "username", options[0].Tag.Get("flag"))
	require.Equal(t, "", options[0].Value)
}

type mockCommandOptionInt struct {
	Amount int `flag:"amount"`
}

func Test_get_parsed_option_int(t *testing.T) {
	options := service.GetOptions(mockCommandOptionInt{})

	require.Equal(t, "amount", options[0].Tag.Get("flag"))
	require.Equal(t, 0, options[0].Value)
}

type mockCommandOptionFloat struct {
	Number float64 `flag:"number"`
}

func Test_get_parsed_option_float(t *testing.T) {
	options := service.GetOptions(mockCommandOptionFloat{})

	require.Equal(t, "number", options[0].Tag.Get("flag"))
	require.Equal(t, 0., options[0].Value)
}

type mockCommandOptionsWithDescription struct {
	DryRun bool `short:"dr" flag:"dry-run" description:"Execute the command as a dry run"`
}

func Test_get_parsed_option_with_description(t *testing.T) {
	options := service.GetOptions(mockCommandOptionsWithDescription{})

	require.Equal(t, "dr", options[0].Tag.Get("short"))
	require.Equal(t, "dry-run", options[0].Tag.Get("flag"))
	require.Equal(t, "Execute the command as a dry run", options[0].Tag.Get("description"))
}


func (s structWithDescription) Name() string        { return "test" }
func (s structWithDescription) Description() string { return "test" }
func (s structWithDescription) Handle(app inter.App, writer io.Writer) inter.ExitCode {
	return inter.Success
}

func Test_show_help_description_of_wrong_flag(t *testing.T) {
	output, app := setUp()
	app.Bind("config.App.OsArgs", []interface{}{"/exe/main", "test", "--fake_flag"})

	code := console.Kernel{
		App:      app,
		Output:   &output,
		Commands: []inter.Command{structWithDescription{}},
	}.Handle()

	require.Equal(t, inter.Failure, code)
	require.Contains(t, TrimDoubleSpaces(output.String()), `flag provided but not defined: -fake_flag`)
	require.Contains(t, TrimDoubleSpaces(
		output.String()),
		`
 -dry-run
 	The flag description
`)
}

func Test_show_help_description_of_wrong_short(t *testing.T) {
	output, app := setUp()
	app.Bind("config.App.OsArgs", []interface{}{"/exe/main", "test", "--fake_short"})

	code := console.Kernel{
		App:      app,
		Output:   &output,
		Commands: []inter.Command{structWithDescription{}},
	}.Handle()

	require.Equal(t, inter.Failure, code)
	require.Contains(t, TrimDoubleSpaces(output.String()), `flag provided but not defined: -fake_short`)
	require.Contains(t, TrimDoubleSpaces(
		output.String()),
		"-dr\n \t\n -dry-run\n \tThe flag description")
}


// todo test empty flag + empty short flag