package pkg

import (
	"fmt"
	"testing"
)

func TestEventGeneration(t *testing.T) {
	generateEventType := generateEventType()
	fmt.Println(generateEventType)
}
