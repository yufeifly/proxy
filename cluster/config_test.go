package cluster

import "testing"

func TestLoadClusterConfig(t *testing.T) {
	err := LoadClusterConfig()
	if err != nil {
		t.Errorf("TestLoadClusterConfig err: %v", err)
	}
}
