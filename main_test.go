package main

import (
	"regexp"
	"testing"
)

func TestValidUUID(t *testing.T) {

	uuid1 := "daf04ba6-b9dc-48d0-810a-58458f23a44f"

	uuid2 := "aa072e9e-cb04-45c8-9ad6-cfeae76dcbe7"

	re := regexp.MustCompile("^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$")

	ub1 := re.MatchString(uuid1)
	ub2 := re.MatchString(uuid2)

	// Valid UUIDs have 36 chars
	if ub1 != true {
		t.Errorf("UUID1 not a valid uuid, got: %t, want: %t", ub1, true)
	}

	if ub2 != true {
		t.Errorf("UUID2 not a valid uuid, got: %v, want: %v", ub2, true)
	}
}
