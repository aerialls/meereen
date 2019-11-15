package processor

import (
	"os/exec"
	"strings"

	d "github.com/aerialls/meereen/pkg/data"
	p "github.com/aerialls/meereen/pkg/processor"
)

// Exec processor
type Exec struct {
	command string
}

// NewExec validates and returns a new exec processor
func NewExec(data map[string]string) (p.Processor, error) {
	command, err := d.GetRequiredParameter(data, "command")
	if err != nil {
		return nil, err
	}

	return &Exec{
		command: command,
	}, nil
}

// Process the exec processor
func (e *Exec) Process() (p.State, string) {
	args := strings.Fields(e.command)
	cmd := exec.Command(args[0], args[1:]...)

	if err := cmd.Run(); err != nil {
		return p.Error, err.Error()
	}
	return p.Ok, ""
}

func init() {
	p.RegisterProcessor("exec", NewExec)
}
