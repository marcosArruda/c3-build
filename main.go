package main

import (
	"os"

	"github.com/codegangsta/cli"
)

func main() {

	app := cli.NewApp()
	app.Name = Name
	app.Version = Version
	app.Author = "marcosArruda"
	app.Email = "marcos.arruda@summa.com.br"
	app.Usage = ""

	app.Flags = GlobalFlags
	app.Commands = Commands
	app.CommandNotFound = CommandNotFound

	app.Run(os.Args)
}
