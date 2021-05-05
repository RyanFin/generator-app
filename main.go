package main

import (
	"flag"
	"fmt"
	"time"
)

type Event struct {
	Type string `json:"type"`
	Data struct {
		Viewid        string    `json:"viewId"`
		Eventdatetime time.Time `json:"eventDateTime"`
	} `json:"data"`
}

func generateEvents(numberOfGroups int, batchSize int, interval int, outpudDir string) {

}

func main() {

	// Get flags from the command line with default values
	numbOfGroupsPtr := flag.Int("number-of-groups", 0, "flag for number of groups")
	batchSizePtr := flag.Int("batch-size", 0, "flag specifying batch size")
	intervalPtr := flag.Int("interval", 0, "interval flag")
	outputDirPtr := flag.String("output-directory", ".", "specify output directory")
	flag.Parse()

	fmt.Println("numbOfGroupsPtr", *numbOfGroupsPtr)
	fmt.Println("batchSizePtr", *batchSizePtr)
	fmt.Println("intervalPtr", *intervalPtr)
	fmt.Println("outputDirPtr", *outputDirPtr)

	// Run generate Events JSON

}
