package pkg

import (
	"fmt"
	"testing"
	"time"

	uuid "github.com/satori/go.uuid"
)

func TestEventObjectPopulation(t *testing.T) {
	var event Event
	event.loadEventData("Viewed", uuid.NewV4().String(), time.Now().Format(time.RFC3339))
	// Test for type
	if event.Type != "Viewed" {
		t.Errorf("Incorrect Type value -  expected %s, got: %s", "Viewed", event.Type)
	}
}

func TestViewedEventType(t *testing.T) {

	eventTypeStr := GenerateEventType(0.5)
	if eventTypeStr != "Viewed" {
		t.Errorf("Incorrect String Value want: %s, got: %s", "Viewed", eventTypeStr)
	}
}

func TestInteractedEventType(t *testing.T) {

	eventTypeStr := GenerateEventType(0.88)
	if eventTypeStr != "Interacted" {
		t.Errorf("Incorrect String Value want: %s, got: %s", "Interacted", eventTypeStr)
	}
}

func TestClickedThroughEventType(t *testing.T) {

	eventTypeStr := GenerateEventType(0.97)
	if eventTypeStr != "Click-Through" {
		t.Errorf("Incorrect String Value want: %s, got: %s", "Click-Through", eventTypeStr)
	}
}

func TestGenerateEvents(t *testing.T) {
	jsonStrOutput := GenerateEvents(100, 400, 1)
	fmt.Println("json response: ", jsonStrOutput)
}

// Random float must be between 0 and 1
func TestGenerateRandomEventFloat(t *testing.T) {
	myFloat := GenerateRandomEventFloat()
	if myFloat < 0.0 || myFloat > 1.0 {
		t.Error("float not within  specified range 0.0 - 1.0")
	}
}
