package main

import "partey/src"

func main() {
	command := src.PartialsCommand{Performable: src.NewPerformable()}
	command.Run()
}
