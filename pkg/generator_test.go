package pkg

import (
	"fmt"
	"testing"
)

func TestEventGeneration(t *testing.T) {
	generateEventType := GenerateEventType()
	fmt.Println(generateEventType)
}

func TestGenerateEvents(t *testing.T) {
	jsonStrOutput := GenerateEvents(100, 400, 1)
	fmt.Println("json response: ", jsonStrOutput)
}
