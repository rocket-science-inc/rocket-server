package generator

import (
	"os"

	cmd "github.com/99designs/gqlgen/cmd"
) 

// GraphQL generates GraphQL server
func GraphQL() {
	os.Args = []string{"", "-c", "cmd/generator/gqlgen.yml", "-v"}
	cmd.Execute()
}
