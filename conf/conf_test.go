package conf

import (
	"fmt"
	"testing"
)

func TestLoadConf(t *testing.T) {
	fmt.Println("Test LoadConf...")

	err := LoadConf("xy-inc.json")
	if err != nil {
		t.Error("Error on load conf.")
	}
}
