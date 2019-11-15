package processor

import (
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"

	p "github.com/aerialls/meereen/pkg/processor"
)

func TestNewExec(t *testing.T) {
	ExecWrapper := func(params map[string]string) error {
		_, err := NewExec(params)
		return err
	}

	errors := map[error]bool{
		// empty
		ExecWrapper(map[string]string{}): false,
		// missing "command"
		ExecWrapper(map[string]string{
			"foo": "bar",
		}): false,
		// good
		ExecWrapper(map[string]string{
			"command": "ls -al",
		}): true,
	}

	for err, ok := range errors {
		if ok {
			assert.Nil(t, err)
		} else {
			assert.NotNil(t, err)
		}
	}
}

func TestExecUnixCommands(t *testing.T) {
	if runtime.GOOS == "windows" {
		return
	}

	commands := map[string]bool{
		"ls":     true,
		"ls -al": true,
		"pwd":    true,
		"foobar": false,
	}

	for command, ok := range commands {
		processor, err := NewExec(map[string]string{
			"command": command,
		})

		assert.Nil(t, err)

		state, _ := processor.Process()

		if ok {
			assert.Equal(t, state, p.Ok)
		} else {
			assert.Equal(t, state, p.Error)
		}
	}

}
