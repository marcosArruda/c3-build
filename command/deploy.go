package command

import (
	"github.com/codegangsta/cli"
	"github.com/fatih/color"
)

func CmdDeploy(c *cli.Context) {

	//USANDO CORES
	// Print with default helper functions
	color.Cyan("Prints text in cyan.")

	// A newline will be appended automatically
	color.Blue("Prints %s in blue.", "text")

	// These are using the default foreground colors
	color.Red("We have red")
	color.Magenta("And many others ..")

	println(c.Args()[0])
	//USAR ESSA ESTRUTURA ABAIXO PARA RODAR SHELL SCRIPTS.
	// docker build current directory
	/*
		cmdName := "docker"
		cmdArgs := []string{"build", "."}

		cmd := exec.Command(cmdName, cmdArgs...)
		cmdReader, err := cmd.StdoutPipe()
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error creating StdoutPipe for Cmd", err)
			os.Exit(1)
		}

		scanner := bufio.NewScanner(cmdReader)
		go func() {
			for scanner.Scan() {
				fmt.Printf("docker build out | %s\n", scanner.Text())
			}
		}()

		err = cmd.Start()
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error starting Cmd", err)
			os.Exit(1)
		}

		err = cmd.Wait()
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error waiting for Cmd", err)
			os.Exit(1)
		}
	*/
}
