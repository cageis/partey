package src

type PartialsCommand struct {
    Performable
}

func (p PartialsCommand) GetSubCommand() string {
    return p.GetArgN(0)
}

func (p PartialsCommand) Run() {
    println(p.GetSubCommand())
    if p.GetSubCommand() == "build" {
        subCommand := NewPartialsBuildCommand(p.GetArgN(1), p.GetArgN(2), p.GetArgN(3))
        subCommand.Run()
    }
}
