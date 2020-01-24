package main

import (
	"vsp"
	"flag"
	"fmt"
)

func main() {
	fmt.Println("VSP")
	nbwordCli := flag.Int("nb", 3, "Number of word to include in the password")
	apiCli := flag.Bool("api", false, "a bool")
	portCli := flag.Int("port", 8080, "API port")
	pathWord := flag.String("file", "/tmp/dico_fr.txt", "Path to the file with all words")
	flag.Parse()
	v := vsp.New(*nbwordCli, *pathWord)
	vsp.ReadFile(v)
	if *apiCli {

		vsp.Server(*portCli, *v)
	} else {

		vsp.CreatePassword(v)
	}
}
