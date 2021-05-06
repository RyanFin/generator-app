package pkg

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os/user"
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

	i := 0

	for i < batchSize {
		// Create an event object
		var event Event

		// Create a unique uuid for this event
		viewId := uuid.NewV4()

		randEvent := GenerateRandomEventFloat()

		eventTypeStr := GenerateEventType(randEvent)

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

func GenerateEventType(randEvent float64) string {

	var ret string

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

func GenerateOutputFile(jsonString string, path string) {

	// Generate timestamp string
	timestmp := time.Now().Format("2006-01-02-15-04-05")

	// Generate file name
	fileName := string(timestmp + ".json")

	myUser, error := user.Current()
	if error != nil {
		panic(error)
	}

	homedir := myUser.HomeDir

	// Write to file
	err := ioutil.WriteFile(string(homedir+"/"+path+fileName), []byte(jsonString), 0644)
	if err != nil {
		fmt.Errorf("Error writing file: %s", err)
	}

	// Save json at target location as a json file
	fmt.Println("file output location: ", string(homedir+"/"+path+fileName))
}

func GenerateRandomEventFloat() float64 {
	var randEvent float64
	// generate a random event
	rand.Seed(time.Now().UnixNano()) // randomise with the use of a seed
	res := make([]float64, 5)
	min := 0.0
	max := 1.0

	for i := range res {
		res[i] = min + rand.Float64()*(max-min)
	}

	randEvent = rand.Float64()

	return randEvent
}
