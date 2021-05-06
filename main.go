package main

import (
	"RyanFin/generator-app/pkg"
	"flag"
)

func main() {

	// Get flags from the command line with default values
	numbOfGroupsPtr := flag.Int("number-of-groups", 0, "flag for number of groups")
	batchSizePtr := flag.Int("batch-size", 0, "flag specifying batch size")
	intervalPtr := flag.Int("interval", 0, "interval flag")
	outputDirPtr := flag.String("output-directory", ".", "specify output directory")
	flag.Parse()

	// Run generate Events JSON
	pkg.GenerateEvents(*numbOfGroupsPtr, *batchSizePtr, *intervalPtr, *outputDirPtr)

}
