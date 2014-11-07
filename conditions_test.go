package openweathermap

import (
	"fmt"
	"os"
	"testing"
)

func TestRetrieveIcon(t *testing.T) {
	tmpDir := "/tmp"
	iconFile := "01d.png"
	s, err := RetrieveIcon(tmpDir, iconFile)
	if err != nil {
		t.Error("Unable to retrieve icon file")
	}
	f, err := os.Stat(fmt.Sprintf("%s/%s", tmpDir, iconFile))
	if err != nil {
		t.Error("Test file does not exist")
	}
	if f.Size() != s {
		t.Error("Size of downloaded file does not match actual size of file")
	}
	err = os.Remove(fmt.Sprintf("%s/%s", tmpDir, iconFile))
	if err != nil {
		t.Error("Unable to delete test file")
	}
}
