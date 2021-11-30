package src

import (
	"os"
)

type Performable struct {
	command string
	args    []string
}

func (p *Performable) GetCommandName() string {
	return p.command
}

func (p *Performable) GetArgN(n int) string {
	return p.args[n]
}

func (p *Performable) GetArgs() []string {
	return p.args
}

func NewPerformable() Performable {
	return Performable{"partley", os.Args[1:]}
}
