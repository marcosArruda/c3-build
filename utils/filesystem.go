package utils

import (
	"fmt"
	"log"

	"github.com/kardianos/osext"
)

//Print the full path to the executed file.
func main() {
	folderPath, err := osext.ExecutableFolder()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(folderPath)
}
