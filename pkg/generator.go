package pkg

import (
	"fmt"
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

func GenerateEvents(numberOfGroups int, batchSize int, interval int, outpudDir string) {
	fmt.Println("numbOfGroupsPtr", numberOfGroups)
	fmt.Println("batchSizePtr", batchSize)
	fmt.Println("intervalPtr", interval)
	fmt.Println("outputDirPtr", outpudDir)

	// Create an event object
	var event Event

	// Create a unique uuid for this event
	viewId := uuid.NewV4()

	eventTypeStr := generateEventType()

	event.loadEventData(eventTypeStr, viewId.String(), time.Now().Format(time.RFC3339))

	fmt.Println(event)

}

func generateEventType() string {

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
	fmt.Println(randEvent)

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
