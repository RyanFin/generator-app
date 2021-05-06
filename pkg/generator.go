package pkg

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"time"

	uuid "github.com/satori/go.uuid"
)

type Event struct {
	Type string `json:"type"`
	Data struct {
		Viewid        string `json:"viewId"`
		Eventdatetime string `json:"eventDateTime"`
	} `json:"data"`
}

func (e *Event) loadEventData(typeString string, viewId string, eventDateTime string) {
	e.Type = typeString
	e.Data.Viewid = viewId
	e.Data.Eventdatetime = eventDateTime
}

func GenerateEvents(numberOfGroups int, batchSize int, interval int) string {

	var jsonStr string

	fmt.Println("numbOfGroupsPtr", numberOfGroups)
	fmt.Println("batchSizePtr", batchSize)
	fmt.Println("intervalPtr", interval)

	i := 0

	for i < batchSize {
		// Create an event object
		var event Event

		// Create a unique uuid for this event
		viewId := uuid.NewV4()

		eventTypeStr := GenerateEventType()
		event.loadEventData(eventTypeStr, viewId.String(), time.Now().Format(time.RFC3339))
		i++

		// marshall event into json format
		j, err := json.Marshal(event)
		if err != nil {
			fmt.Errorf("Error Marshalling JSON: %s", err)
		}
		// append json string to return variable
		jsonStr = jsonStr + string(j) + ("\n")
	}

	return jsonStr

}

func GenerateEventType() string {

	var ret string

	// generate a random event
	rand.Seed(time.Now().UnixNano()) // randomise with the use of a seed
	res := make([]float64, 5)
	min := 0.0
	max := 1.0

	for i := range res {
		res[i] = min + rand.Float64()*(max-min)
	}
	randEvent := rand.Float64()

	if randEvent < 0.85 {
		ret = "Viewed"
	}

	if randEvent > 0.85 {
		ret = "Interacted"
	}

	if randEvent > 0.95 {
		ret = "Click-Through"
	}

	return ret
}

func GenerateOutputFile(jsonString string, outputDirPtr *string) {

	// Generate timestamp string
	timestmp := time.Now().Format("2006-01-02-15-04-05")
	// Generate file name
	fileName := string(timestmp + ".json")
	err := ioutil.WriteFile(fileName, []byte(jsonString), 0644)
	if err != nil {
		fmt.Errorf("Error writing file: %s", err)
	}

	// Save json at target location as a json file
	fmt.Println(outputDirPtr)
}
