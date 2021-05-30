package cluster

import "testing"

func TestLoadConfig(t *testing.T) {
	err := LoadConfig()
	if err != nil {
		t.Errorf("TestLoadConfig err: %v", err)
	}
}
