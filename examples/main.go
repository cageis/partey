package examples

import partley2 "partey/src"

func partley() {
	command := partley2.PartialsCommand{Performable: partley2.NewPerformable()}
	command.Run()
}
