package main

import (
	"flag"
	"os"

	service "rocket-server/server/api/cmd/service"
	generator "rocket-server/server/api/cmd/generator"
) 

var mainfs = flag.NewFlagSet("cmd", flag.ContinueOnError)
var generate = mainfs.Bool("generate", false, "Generate GraphQL server")

func main() {
	mainfs.Parse(os.Args[1:])
	
	if *generate {
		// run graphql generator
		generator.GraphQL()	
	} else {
		// run service
		service.Run()
	}
}
