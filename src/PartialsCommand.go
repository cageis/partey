package src

type PartialsCommand struct {
    Performable
}

func (p PartialsCommand) GetSubCommand() string {
    return p.GetArgN(0)
}

func (p PartialsCommand) Run() {
    subCommand := NewPartialsBuildCommand(p.GetArgN(0), p.GetArgN(1), p.GetArgN(2))
    subCommand.Run()
}
