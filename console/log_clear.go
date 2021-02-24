package console

import (
	"fmt"
	"github.com/confetti-framework/contract/inter"
	"io"
)

type LogClear struct {}

func (l LogClear) Name() string {
	return "log:clear"
}

func (l LogClear) Description() string {
	return "Clear the log files as indicated in the configuration."
}

func (l LogClear) Handle(app inter.App, writer io.Writer) inter.ExitCode {
	channelsRaw, err := app.MakeE("config.Logging.Channels")
	loggers := channelsRaw.(map[string]interface{})
	if err != nil || len(loggers) == 0 {
		_, _ = fmt.Fprintln(writer, "No files to clear. No loggers found")
	}

	for channel, logger := range loggers {
		if logger.(inter.Logger).Clear() {
			_, _ = fmt.Fprintf(writer, "Files cleaned for channel: %s\n", channel)
		}
	}

	_, _ = fmt.Fprintf(writer, "Done")

	return inter.Success
}
